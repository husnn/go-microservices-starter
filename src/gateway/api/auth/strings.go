package auth

import "golang.org/x/text/message"

var msgLoginTitle = message.Key("msgLoginTitle", "Login")
var msgLoginHeading = message.Key("msgLoginHeading", "Log into your account")
var msgLoginSubheadingEmail = message.Key("msgLoginSubheadingEmail", "Enter your email and password to continue.")
var msgLoginSubheadingPhone = message.Key("msgLoginSubheadingPhone", "Enter your phone number and password to continue.")

var msgSignupHeading = message.Key("msgSignupHeading", "Create an account")
var msgSignupSubheadingEmail = message.Key("msgSignupSubheadingEmail", "Enter your email and password to continue.")
var msgSignupSubheadingPhone = message.Key("msgSignupSubheadingPhone", "Enter your phone number and password to continue.")

var msgLoginUseEmail = message.Key("msgLoginUseEmail", "Use email address")
var msgLoginUsePhone = message.Key("msgLoginUsePhone", "Use phone number")

var msgLoginStepOTPHeading = message.Key("msgLoginStepOTPHeading",
	"Enter the six-digit code we sent you")

var msgLoginStepOTPSubheadingEmail = message.Key("msgLoginStepOTPSubheadingEmail",
	"You should have received a one-time code to your email.")
var msgLoginStepOTPSubheadingPhone = message.Key("msgLoginStepOTPSubheadingPhone",
	"You should have received a one-time code to your phone number.")

var msgLoginStepOTPInputLabel = message.Key("msgLoginStepOTPInputLabel",
	"One-time password")

var msgLoginStepOTPSubmit = message.Key("msgLoginStepOTPSubmit",
	"Submit code")
var msgLoginStepOTPBack = message.Key("msgLoginStepOTPBack",
	"Sign into a different account")

var msgLoginStepSubmit = message.Key("msgLoginStepSubmit", "Login")
var msgLoginStepResetPassword = message.Key("msgLoginStepResetPassword", "Forgot password?")
var msgLoginStepRegister = message.Key("msgLoginStepRegister", "Don't have an account? Sign up.")

var msgLoginErrUserNotFound = message.Key("msgLoginErrUserNotFound", "We couldn't find a user with those details.")
var msgLoginErrInvalidEmail = message.Key("msgLoginErrInvalidEmail", "Invalid email address provided.")
var msgLoginErrInvalidPhone = message.Key("msgLoginErrInvalidPhone", "Invalid phone number provided.")
var msgLoginErrInvalidPassword = message.Key("msgLoginErrInvalidPassword", "Invalid password provided.")
var msgLoginErrIncorrectPassword = message.Key("msgLoginErrIncorrectPassword", "Incorrect password.")
var msgLoginErrIncorrectOtp = message.Key("msgLoginErrIncorrectOtp", "Incorrect code entered.")
var msgLoginErrGeneral = message.Key("msgLoginErrGeneral", "Could not login. Please try again.")

var msgSignupFooter = message.Key("msgSignupFooter", "By signing up, you agree to boilerplate's terms & conditions and privacy policy.")
var msgSignupCreateAccount = message.Key("msgSignupCreateAccount", "Create account")
var msgSignupErrGeneral = message.Key("msgSignupErrGeneral", "Could not sign up.")
var msgSignupErrUserExists = message.Key("msgSignupErrUserExists", "User with given email or phone already exists.")

var msgResetPasswordHeading = message.Key("msgResetPasswordHeading", "Forgot your password?")
var msgResetPasswordHeadingOTP = message.Key("msgResetPasswordHeadingOTP", "We sent you a code")
var msgResetPasswordHeadingOTPResent = message.Key("msgResetPasswordHeadingOTPResent", "We sent you another code")
var msgResetPasswordHeadingNewPassword = message.Key("msgResetPasswordHeadingNewPassword", "Set your new password")

var msgResetPasswordSubheadingEmail = message.Key("msgResetPasswordSubheadingEmail", "Enter the email address linked to your account to reset your password.")
var msgResetPasswordSubheadingPhone = message.Key("msgResetPasswordSubheadingPhone", "Enter the phone number linked to your account to reset your password.")

var msgResetPasswordSubheadingOTPEmail = message.Key("msgResetPasswordSubheadingOTPEmail", "A code was sent to the email address provided. Enter that here to set a new password.")
var msgResetPasswordSubheadingOTPPhone = message.Key("msgResetPasswordSubheadingOTPPhone", "A code was sent to the phone number provided. Enter that here to set a new password.")

var msgResetPasswordSubheadingNewPassword = message.Key("msgResetPasswordSubheadingNewPassword", "Enter new password to use in the future.")

var msgResetPasswordRequestSubmit = message.Key("msgResetPasswordRequestSubmit", "Request password reset")
var msgResetPasswordSubmitOTP = message.Key("msgResetPasswordSubmitOTP", "Submit code")
var msgResetPasswordOTPRequest = message.Key("msgResetPasswordOTPRequest", "Request another code")
var msgResetPasswordOTPRequestCountdown = message.Key("msgResetPasswordOTPRequestCountdown", "<!tick=1!>You can request another code in <!seconds_until_unix=%s!> seconds")
var msgResetPasswordOTPExpired = message.Key("msgResetPasswordOTPExpired", "Code expired. Try again.")
var msgResetPasswordOTPResendError = message.Key("msgResetPasswordOTPResendError", "Could not resend code. Please try again shortly.")
var msgResetPasswordChangePassword = message.Key("msgResetPasswordChangePassword", "Change password")

var msgResetPasswordNewLabel = message.Key("msgResetPasswordNewLabel", "Password")
var msgResetPasswordNewPlaceholder = message.Key("msgResetPasswordNewPlaceholder", "New password")

var msgResetPasswordOtpLabel = message.Key("msgResetPasswordOtpLabel", "One-time code")
var msgResetPasswordOtpMissing = message.Key("msgResetPasswordOtpMissing", "Please enter a valid code")

var msgResetPasswordErrGeneral = message.Key("msgResetPasswordErrGeneral", "Could not reset password. Please try again.")
var msgResetPasswordErrIncorrectOtp = message.Key("msgResetPasswordErrIncorrectOtp", "Incorrect code entered.")
var msgResetPasswordTooManyAttemptsError = message.Key("msgResetPasswordErrTooManyAttempts", "Attempts exceeded.")
var msgResetPasswordErrGrantExpired = message.Key("msgResetPasswordErrGrantExpired", "Code expired.")
var msgResetPasswordErrGrantExhausted = message.Key("msgResetPasswordErrGrantExhausted", "Code already used.")
