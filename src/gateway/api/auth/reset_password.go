package auth

import (
	"fmt"
	"github.com/luno/jettison/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/message"
	"net/http"
	"boilerplate/gateway/httpx"
	"boilerplate/gateway/state"
	"boilerplate/guard"
	authpb "boilerplate/proto/auth"
	"boilerplate/proto/components"
	"boilerplate/users"
)

func resetPasswordResponse(mpr *message.Printer, req *authpb.ResetPasswordRequest,
	grantId string, err message.Reference) *authpb.ResetPasswordResponse {
	var inputs []*components.Input
	var actions []*components.Action
	var formGroupActions []*components.Action

	step := authpb.ResetPasswordStep_RESET_PASSWORD_STEP_INITIAL
	heading := msgResetPasswordHeading
	subheading := msgResetPasswordSubheadingEmail
	method := authpb.AuthMethod_AUTH_METHOD_EMAIL
	toggleMessage := msgLoginUsePhone

	usePhone := req.AuthMethod == authpb.AuthMethod_AUTH_METHOD_PHONE

	var codeResent bool
	if req.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_RESEND {
		req.Step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_INITIAL
		if err == "" {
			codeResent = true
		}
	}

	if req.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_INITIAL {
		heading = msgResetPasswordHeadingOTP
		subheading = msgResetPasswordSubheadingOTPEmail
		if usePhone {
			subheading = msgResetPasswordSubheadingOTPPhone
		}

		if codeResent {
			heading = msgResetPasswordHeadingOTPResent
		}

		step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_SUBMIT_OTP
		inputs = []*components.Input{
			{
				Type:        components.Input_OTP,
				Name:        "otp",
				Label:       mpr.Sprintf(msgResetPasswordOtpLabel),
				Placeholder: "••••••",
				Value:       "",
				Validation: &components.Input_Validation{
					MinLength: guard.OTPLength,
					MaxLength: guard.OTPLength,
				}},
		}

		resendLink := &components.Action{
			Type:     components.Action_TYPE_ENDPOINT,
			Style:    components.Style_STYLE_PRIMARY,
			Element:  components.Element_ELEMENT_LINK,
			Text:     mpr.Sprintf(msgResetPasswordOTPRequest),
			Endpoint: components.Action_ENDPOINT_AUTH_REQUEST_OTP,
		}

		var resendAction *components.Action

		if req.NextOtpSend > 0 {
			resendAction = &components.Action{
				Style:   components.Style_STYLE_TERTIARY,
				Element: components.Element_ELEMENT_TEXT,
				Text: mpr.Sprintf(msgResetPasswordOTPRequestCountdown,
					fmt.Sprint(req.NextOtpSend)),
				NextAfter: req.NextOtpSend,
				Next:      resendLink,
			}
		} else {
			resendAction = resendLink
		}

		formGroupActions = []*components.Action{resendAction}

		actions = []*components.Action{
			{
				Type:     components.Action_TYPE_ENDPOINT,
				Style:    components.Style_STYLE_PRIMARY,
				Element:  components.Element_ELEMENT_BUTTON,
				Text:     mpr.Sprintf(msgResetPasswordSubmitOTP),
				Endpoint: components.Action_ENDPOINT_AUTH_RESET_PASSWORD,
			},
		}
	} else if req.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_SUBMIT_OTP {
		heading = msgResetPasswordHeadingNewPassword
		subheading = msgResetPasswordSubheadingNewPassword

		step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_NEW_PASSWORD
		inputs = []*components.Input{
			{
				Type:        components.Input_PASSWORD,
				Name:        "password",
				Label:       mpr.Sprintf(msgResetPasswordNewLabel),
				Placeholder: mpr.Sprintf(msgResetPasswordNewPlaceholder),
				Value:       "",
			},
		}

		actions = []*components.Action{
			{
				Type:     components.Action_TYPE_ENDPOINT,
				Style:    components.Style_STYLE_PRIMARY,
				Element:  components.Element_ELEMENT_BUTTON,
				Text:     mpr.Sprintf(msgResetPasswordChangePassword),
				Endpoint: components.Action_ENDPOINT_AUTH_RESET_PASSWORD,
			},
		}
	} else {
		if usePhone {
			subheading = msgResetPasswordSubheadingPhone
			method = authpb.AuthMethod_AUTH_METHOD_PHONE
			toggleMessage = msgLoginUseEmail
			inputs = []*components.Input{
				{
					Type:        components.Input_PHONE,
					Name:        "phone",
					Placeholder: "+92 3xx xxxx xxx",
					Value:       req.GetPhone(),
				},
			}
		} else {
			inputs = []*components.Input{
				{
					Type:        components.Input_EMAIL,
					Name:        "email",
					Placeholder: "wayne@boilerplate.pk",
					Value:       req.GetEmail(),
				},
			}
		}

		actions = []*components.Action{
			{
				Type:     components.Action_TYPE_ENDPOINT,
				Style:    components.Style_STYLE_PRIMARY,
				Element:  components.Element_ELEMENT_BUTTON,
				Text:     mpr.Sprintf(msgResetPasswordRequestSubmit),
				Endpoint: components.Action_ENDPOINT_AUTH_RESET_PASSWORD,
			},
		}
	}

	res := &authpb.ResetPasswordResponse{
		Step: step,
		Screen: &components.Screen{
			Type: components.ScreenType_SCREEN_RESET_PASSWORD,
			Content: &components.Screen_ResetPasswordScreen{
				ResetPasswordScreen: &components.ResetPasswordScreen{
					Form: &components.Form{
						Heading:    mpr.Sprintf(heading),
						Subheading: mpr.Sprintf(subheading),
						Groups: []*components.Form_Group{
							{
								Inputs:  inputs,
								Actions: formGroupActions,
								Position: &components.GroupPosition{
									PosX: components.Position_POSITION_FILL,
									PosY: components.Position_POSITION_START,
								},
							},
						},
						Error:   mpr.Sprintf(err),
						Actions: actions,
					},
				},
			},
		},
		GrantId:     grantId,
		NextOtpSend: req.NextOtpSend,
		AuthMethod:  method,
	}

	if req.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_UNSPECIFIED {
		res.Screen.ActionsTop = &components.ActionGroup{
			Actions: []*components.Action{
				{
					Type:     components.Action_TYPE_TOGGLE,
					Style:    components.Style_STYLE_PRIMARY,
					Element:  components.Element_ELEMENT_LINK,
					Text:     mpr.Sprintf(toggleMessage),
					Endpoint: components.Action_ENDPOINT_AUTH_RESET_PASSWORD,
					Toggled:  usePhone,
				},
			},
			Position: &components.GroupPosition{
				PosX:    components.Position_POSITION_FILL,
				PosY:    components.Position_POSITION_START,
				IsRow:   true,
				Reverse: true,
			},
		}
	}

	return res
}

func ResetPassword(d state.Deps, userType users.UserType) httpx.Handler {
	return func(w http.ResponseWriter, r *httpx.Request) {
		var reqpb authpb.ResetPasswordRequest
		if err := httpx.ParseJSON(r, &reqpb); err != nil {
			httpx.Fail(w, err)
			return
		}

		mpr := message.NewPrinter(r.Lang)

		var res *authpb.ResetPasswordResponse
		if reqpb.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_INITIAL {
			gid, nextOtpSend, err := d.UsersClient().RequestPasswordReset(r.Context(), userType,
				reqpb.GetEmail(), reqpb.GetPhone(), r.IP, r.UserAgent())
			if err != nil {
				if errors.Is(err, users.ErrUserNotFound) {
					reqpb.Step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_UNSPECIFIED
					res = resetPasswordResponse(mpr, &reqpb, reqpb.GrantId, msgLoginErrUserNotFound)
				} else {
					httpx.Fail(w, err)
					return
				}
			} else {
				reqpb.NextOtpSend = nextOtpSend.Unix()
				res = resetPasswordResponse(mpr, &reqpb, gid, "")
			}
		} else if reqpb.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_SUBMIT_OTP {
			if reqpb.Otp == "" {
				reqpb.Step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_INITIAL
				res = resetPasswordResponse(mpr, &reqpb, reqpb.GrantId, msgResetPasswordOtpMissing)
			} else {
				res = resetPasswordResponse(mpr, &reqpb, reqpb.GrantId, "")
			}
		} else if reqpb.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_RESEND {
			nextResend, err := d.GuardClient().ResendOTP(r.Context(), reqpb.GrantId, r.IP)
			if err != nil {
				if errors.Is(err, guard.ErrGrantExpired) {
					reqpb.Step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_UNSPECIFIED
					res = resetPasswordResponse(mpr, &reqpb, reqpb.GrantId, msgResetPasswordOTPExpired)
				} else {
					res = resetPasswordResponse(mpr, &reqpb, reqpb.GrantId, msgResetPasswordOTPResendError)
				}
				log.Error().Err(err).Msg("")
			} else {
				reqpb.NextOtpSend = nextResend.Unix()
				res = resetPasswordResponse(mpr, &reqpb, reqpb.GrantId, "")
			}
		} else if reqpb.Step == authpb.ResetPasswordStep_RESET_PASSWORD_STEP_NEW_PASSWORD {
			reqpb.Step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_SUBMIT_OTP
			err := d.UsersClient().ResetPassword(r.Context(),
				reqpb.Password, reqpb.GrantId, r.IP, reqpb.Otp)
			if err != nil {
				errMsg := msgResetPasswordErrGeneral

				if errors.Is(err, users.ErrInvalidPassword) {
					errMsg = msgLoginErrInvalidPassword
				} else if errors.Is(err, guard.ErrIncorrectOtp) {
					reqpb.Step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_INITIAL
					errMsg = msgResetPasswordErrIncorrectOtp
				} else {
					reqpb.Step = authpb.ResetPasswordStep_RESET_PASSWORD_STEP_UNSPECIFIED
					if errors.Is(err, guard.ErrTooManyAttempts) {
						errMsg = msgResetPasswordTooManyAttemptsError
					} else if errors.Is(err, guard.ErrGrantExpired) {
						errMsg = msgResetPasswordErrGrantExpired
					} else if errors.Is(err, guard.ErrGrantExhausted) {
						errMsg = msgResetPasswordErrGrantExhausted
					}
				}

				res = resetPasswordResponse(mpr, &reqpb, reqpb.GrantId, errMsg)
				log.Error().Err(err).Msg("")
			} else {
				res = &authpb.ResetPasswordResponse{
					Step: authpb.ResetPasswordStep_RESET_PASSWORD_STEP_SUCCESS,
					Redirect: &components.Action{
						Type:   components.Action_TYPE_REDIRECT,
						Screen: components.ScreenType_SCREEN_LOGIN,
						Url:    "/login",
					},
				}
			}
		} else {
			res = resetPasswordResponse(mpr, &reqpb, "", "")
		}

		httpx.Ok(w, httpx.WithMessage(res))
	}
}
