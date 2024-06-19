package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"github.com/panjf2000/ants/v2"
	"google.golang.org/api/option"
)

var lang_str string

var filter_str string

var RES_DIR string

var wg sync.WaitGroup

var pool *ants.Pool

var langs = []string{
	`fr`, `ru`, `zh`, `de`, `es`, `uk`,
	//  `ja`, `ko`, `vi`, `zh-rTW`,
}

func init() {
	cwd, _ := os.Getwd()
	RES_DIR = cwd + "/../app/src/main/res"

	flag.StringVar(&lang_str, "lang", "", fmt.Sprintf("Required, available languages: %v", langs))
	flag.StringVar(&filter_str, "filter", "", "")

	wg = sync.WaitGroup{}
	pool, _ = ants.NewPool(3)
}

func check_param() []string {
	if len(lang_str) == 0 {
		panic("must specify language")
	}
	var languages []string
	if lang_str == "all" {
		languages = langs
	} else {
		languages = strings.Split(lang_str, ",")
	}
	for _, lang := range languages {
		if !slices.Contains(langs, lang) {
			panic("language " + lang + " not supported yet")
		}
	}
	return languages
}

func usage() {

	flag.Usage()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read_file(fn string) string {
	s, err := os.ReadFile(fn)
	if err != nil {
		return ""
	}
	return string(s)
}

func write_file(fn string, data string) error {
	return os.WriteFile(fn, []byte(data), 0666)
}

func translate_1_file(lang string, fn string) error {

	src_file := fmt.Sprintf("%s/values/%s", RES_DIR, fn)
	dest_file := fmt.Sprintf("%s/values-%s/%s", RES_DIR, lang, fn)

	fmt.Printf("processing: %s -> %s\n", filepath.Base(src_file), lang)

	to_translate := read_file(src_file)

	GeminiToken := os.Getenv("GeminiToken")

	text := fmt.Sprintf(
		"Translate the following xml content to %s, it's about a call blocking app "+
			"which blocks incoming spam calls and SMS messages. "+
			"Better use short words, make sure leave the XML tags unmodified, show me the result only:\n"+
			"%s",
		lang, to_translate)

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(GeminiToken))
	check(err)
	defer client.Close()

	model := client.GenerativeModel("gemini-pro")

	// max is 8192 for gemini-pro v1.0 (8192 by default)
	// but actually it's only 2048...
	// model.SetMaxOutputTokens(8192)

	resp, err := model.GenerateContent(ctx, genai.Text(text))
	check(err)

	fmt.Println("  - ",
		"TotalTokenCount", resp.UsageMetadata.TotalTokenCount,
		"PromptTokenCount", resp.UsageMetadata.PromptTokenCount,
		"CandidatesTokenCount", resp.UsageMetadata.CandidatesTokenCount,
	)
	fmt.Println()

	if resp.UsageMetadata.CandidatesTokenCount >= 2048 {
		panic("CandidatesTokenCount reached 2048, preferably < 1800, split the xml")
	}

	sb := &strings.Builder{}

	for _, c := range resp.Candidates {
		if c.Content == nil {
			fmt.Println(c)
			return errors.New("no c.Content returned")
		}
		for _, p := range c.Content.Parts {
			fmt.Fprintf(sb, "%v", p)
		}
	}

	// French has lots of '
	cleared := strings.ReplaceAll(sb.String(), "'", "\\'")
	if !strings.HasPrefix(cleared, "<resources>") {
		return Retryable(errors.New("malformed result"))
	}
	write_file(dest_file, cleared)
	return nil
}

func translate_lang(lang string) {
	path := RES_DIR + "/values-" + lang

	if filter_str == "" {
		fmt.Printf("clearing: %s\n\n", path)
		os.RemoveAll(path)
		os.Mkdir(path, os.ModePerm)
	}

	filepath.Walk(
		RES_DIR+"/values",

		func(path string, fi os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if fi.IsDir() || !strings.HasPrefix(fi.Name(), "strings_") {
				return nil
			}
			if strings.Contains(fi.Name(), "no_translate") {
				return nil
			}

			if filter_str != "" {
				if !strings.Contains(fi.Name(), filter_str) {
					return nil
				}
			}

			wg.Add(1)
			pool.Submit(func() {
				e := Retry(5, func(attempt int) error {
					err := translate_1_file(lang, fi.Name())
					if IsRetryable(err) {
						color.HiWhite("retry %s", color.HiYellowString(fi.Name()))
					}
					if err == nil {
						color.HiWhite("done %s %s", lang, color.HiGreenString(fi.Name()))
					}
					return err
				})
				if e != nil {
					color.HiWhite("translate %s failed", color.HiRedString(fi.Name()))
				}

				wg.Done()
			})
			return nil
		})

}

func main() {
	flag.Parse()

	languages := check_param()

	for _, lang := range languages {
		translate_lang(lang)
	}
	wg.Wait()
}
