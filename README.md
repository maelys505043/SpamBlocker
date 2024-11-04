# This is not the real SpamBlocker
Edited for a single change to the non-contact blocking function.
Android Call/SMS blocker. (Android 10+)

<img src="https://github.com/aj3423/SpamBlocker/assets/4710875/9d44afe7-2524-4b34-8bf3-ba285200bb5c" height="90">  [<img src="https://github.com/user-attachments/assets/8757c78c-b0d5-4b8a-9adb-934d8a758e9e"
     alt="Get it on F-Droid"
     height="70">](https://f-droid.org/packages/spam.blocker/)  


Table of Contents
=================
   * [Screenshot](#screenshot)
   * [Target Audience](#target-audience)
   * [How it works](#how-it-works)
   * [Features](#features)
   * [Permissions](#permissions)
   * [Privacy](#privacy)
   * [Support](#support)
   * [FAQ](#faq)
   * [Language Support](#language-support)
   * [Donate](#donate)

# Screenshot
| Call  | SMS  | Setting  | Notification  |
|-------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------|
| <img src="https://github.com/aj3423/SpamBlocker/assets/4710875/7f03d0a0-d12e-4e1b-a064-2412fc1cee8e" width="200"> | <img src="https://github.com/aj3423/SpamBlocker/assets/4710875/ff1dd6c3-56dc-4f64-96a5-e7ca379af035" width="200"> | <img src="https://github.com/aj3423/SpamBlocker/assets/4710875/a86fff09-d30b-428e-866c-0f07b874d479" width="200"> | <img src="https://github.com/aj3423/SpamBlocker/assets/4710875/633e0e24-5ba0-44d7-90ec-09324081d37b" width="200"> |

# Target Audience
- :white_check_mark: For people who tends to avoid unknown calls.
- For those who need to answer unknown calls, for example, salespeople or lawyers, etc:
  - :white_check_mark: If your carrier supports STIR/SHAKEN attestation, or there is any public spam database available for your country.
  - :x: Otherwise, consider using a blocker with its own spam database.

# How it works
It works without replacing your default call/SMS app.
 - For call: <br>
 &ensp; It's the Caller ID app 
 - For SMS: <br>
 &ensp; It takes over the SMS notifications, it only filters the notifications, the spam messages will still be present in the SMS app.
   > 💡 Please turn off the notification permission of the default SMS app in system settings, otherwise there will be double SMS notifications.


# Features:

| Filter        | It checks  |
|---------------|------|
| Contacts      | Whether from a contact |
| Contact Group | Whether it's a member of some contact group |
| STIR/SHAKEN   | STIR/SHAKEN attestation  |
| Repeated      | Whether it's been calling repeatedly |
| Dialed        | Whether you have dialed the number |
| Recent Apps   | If some specific apps have been used recently, all calls are allowed.<br>Use case:<br>&emsp; You ordered Pizza online and soon they call you to refund.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     |
| In Meeting    | Decline calls during online video meetings |
| Off Time      | A time period that always permits calls, usually no spams at night. |
| Spam Database | If it matches any spam number in the database. Any public downloadable spam databases can be integrated, such as the [DNC](https://www.ftc.gov/policy-notices/open-government/data-sets/do-not-call-data).  |
| Regex Pattern | Some typical patterns:<br> - Any number: `.*` (the regex `.*` is equivalent to the wildcard `*` in other apps) <br> - Exact number: `12345` <br> - Starts with 400: `400.*` <br> - Ends with 123: `.*123` <br> - Shorter than 5: `.{0,4}` <br> - Longer than 10: `.{11,}` <br> - Unknown number (it's empty string): `.{0}` or `^$`<br>  - Contains "verification": `.*verification.*` <br> - Contains any of the words: `.*(police\|hospital\|verification).*` <br> - Starts with 400, with leading country code 11 or not: `(?:11)?400.*` <br>- Extract verification code from SMS message: `code.*?(\d+)`<br><br> Ask AI to generate or explain a regex: <br>&emsp; "Show me regex for checking if a string starts with 400 or 200"<br> &emsp; Results in `(400\|200).*` |


# Permissions 

| Permission (all optional)                                                         | Why                                                                             |
|-----------------------------------------------------------------------------------|---------------------------------------------------------------------------------|
| INTERNET                                                                          | For downloading spam numbers from public databases                              | 
| MANAGE_EXTERNAL_STORAGE (Android 11+)<br>READ/WRITE_EXTERNAL_STORAGE (Android 10) | For file access from automated workflow                                         | 
| ANSWER_PHONE_CALLS                                                                | Reject, Answer and Hang-up calls                                                |
| POST_NOTIFICATIONS                                                                | Show notifications                                                              |
| READ_CONTACTS                                                                     | For matching contacts                                                           |
| RECEIVE_SMS                                                                       | For receiving new messages                                                      |
| READ_CALL_LOG<br>READ_SMS                                                         | For checking if a call is repeated                                              |
| PACKAGE_USAGE_STATS                                                               | For feature: Recent Apps <br>For checking whether an app has been used recently |
| READ_PHONE_STATE                                                                  | For block mode: Answer + Hang-up (monitor ringing state)                        |

# Privacy
 No data collection
 - Internet access:
   - For downloading spam numbers only, if you don't use this feature, you can [disable the internet access](https://github.com/aj3423/SpamBlocker/issues/147), or download the offline version from the release page.
 - No communication with other app
 - [Reproducible](https://f-droid.org/docs/Reproducible_Builds/) apk

# Support
 - Most problems have already been discussed in the issue list, please search first.
 - There's also a [matrix channel](https://matrix.to/#/#spam-blocker:matrix.org)

# FAQ
 - [Google shows security warning when installing this app](https://github.com/aj3423/SpamBlocker/issues/108)
 - [How does the "Priority" work, how to always block a particular number regardless of any other rules](https://github.com/aj3423/SpamBlocker/issues/166)
 - [It stopped working after being killed](https://github.com/aj3423/SpamBlocker/issues/100)
 - [Android 9- support](https://github.com/aj3423/SpamBlocker/issues/38)
 - [Dual SIM support](https://github.com/aj3423/SpamBlocker/issues/169)
 - [Auto delete spam SMS messages](https://github.com/aj3423/SpamBlocker/discussions/164#discussioncomment-11037830)

# Language support

Languages are translated using Gemini AI([golang script](https://github.com/aj3423/SpamBlocker/blob/master/auto_translate/translate.go)), fire an issue for requesting a new language support.

# Donate

:heart:  https://aj3423.github.io/donate
