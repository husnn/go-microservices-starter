package auth

import (
	"context"
	"fmt"
	"github.com/luno/jettison/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"net"
	"net/http"
	"boilerplate/auth"
	"boilerplate/gateway/httpx"
	"boilerplate/gateway/state"
	"boilerplate/guard"
	authpb "boilerplate/proto/auth"
	"boilerplate/proto/components"
	"boilerplate/users"
)

func loginHandler(ctx context.Context, d state.Deps,
	ip net.IP, userAgent string, reqpb *authpb.LoginRequest,
	w http.ResponseWriter) (*authpb.LoginResponse, error) {
	mpr := message.NewPrinter(language.English)

	switch reqpb.Step {
	case authpb.LoginStep_LOGIN_STEP_UNSPECIFIED:
		return identificationResponse(mpr, reqpb, ""), nil
	case authpb.LoginStep_LOGIN_STEP_INITIAL:
		var sid, gid string
		var err error
		if reqpb.AuthMethod == authpb.AuthMethod_AUTH_METHOD_PHONE {
			sid, gid, err = d.AuthClient().LoginWithPhone(ctx,
				users.UserType(reqpb.UserType), reqpb.GetPhone(),
				reqpb.Password, ip, userAgent)
		} else {
			sid, gid, err = d.AuthClient().LoginWithEmail(ctx,
				users.UserType(reqpb.UserType), reqpb.GetEmail(),
				reqpb.Password, ip, userAgent)
		}
		if err != nil {
			errMsg := msgLoginErrGeneral

			if errors.Is(err, users.ErrInvalidEmail) {
				errMsg = msgLoginErrInvalidEmail
			} else if errors.Is(err, users.ErrInvalidPhone) {
				errMsg = msgLoginErrInvalidPhone
			} else if errors.Is(err, users.ErrUserNotFound) {
				errMsg = msgLoginErrUserNotFound
			} else if errors.Is(err, auth.ErrIncorrectPassword) {
				errMsg = msgLoginErrIncorrectPassword
			}

			return identificationResponse(mpr, reqpb,
				mpr.Sprintf(errMsg)), err
		}

		if sid == "" {
			return buildOTPResponse(mpr, reqpb, gid, ""), nil
		}

		return success(w, sid), nil
	case authpb.LoginStep_LOGIN_STEP_SUBMIT_OTP:
		sid, err := d.AuthClient().SubmitOTP(ctx,
			reqpb.GrantId, ip, userAgent, reqpb.Otp)
		if err != nil {
			if errors.Is(err, guard.ErrIncorrectOtp) {
				return buildOTPResponse(mpr, reqpb, reqpb.GrantId, msgLoginErrIncorrectOtp), err
			}

			return identificationResponse(mpr, reqpb, msgLoginErrGeneral), err
		}

		return success(w, sid), nil
	}

	return nil, errors.New("unhandled login stage")
}

func buildOTPResponse(mpr *message.Printer, req *authpb.LoginRequest,
	gid string, err message.Reference) *authpb.LoginResponse {
	method := authpb.AuthMethod_AUTH_METHOD_EMAIL
	subheading := msgLoginStepOTPSubheadingEmail

	usePhone := req.AuthMethod == authpb.AuthMethod_AUTH_METHOD_PHONE
	if usePhone {
		method = authpb.AuthMethod_AUTH_METHOD_PHONE
		subheading = msgLoginStepOTPSubheadingPhone
	}
	return &authpb.LoginResponse{
		Step: authpb.LoginStep_LOGIN_STEP_SUBMIT_OTP,
		Screen: &components.Screen{
			Type: components.ScreenType_SCREEN_LOGIN,
			Content: &components.Screen_Login_Screen{
				Login_Screen: &components.LoginScreen{
					Form: &components.Form{
						Heading:    mpr.Sprintf(msgLoginStepOTPHeading),
						Subheading: mpr.Sprintf(subheading),
						Groups: []*components.Form_Group{
							{
								Position: &components.GroupPosition{
									PosX: components.Position_POSITION_FILL,
									PosY: components.Position_POSITION_START,
								},
								Inputs: []*components.Input{
									{
										Type:  components.Input_NUMERIC,
										Name:  "otp",
										Label: mpr.Sprintf(msgLoginStepOTPInputLabel),
										Validation: &components.Input_Validation{
											MinLength: guard.OTPLength,
											MaxLength: guard.OTPLength,
										},
									},
								},
							},
						},
						Error: mpr.Sprintf(err),
						Actions: []*components.Action{
							{
								Type:     components.Action_TYPE_ENDPOINT,
								Style:    components.Style_STYLE_PRIMARY,
								Element:  components.Element_ELEMENT_BUTTON,
								Text:     mpr.Sprintf(msgLoginStepOTPSubmit),
								Endpoint: components.Action_ENDPOINT_AUTH_LOGIN,
							},
							{
								Type:     components.Action_TYPE_BACK,
								Style:    components.Style_STYLE_SECONDARY,
								Element:  components.Element_ELEMENT_BUTTON,
								Text:     mpr.Sprintf(msgLoginStepOTPBack),
								Endpoint: components.Action_ENDPOINT_AUTH_LOGIN,
							},
						},
					},
				},
			},
		},
		GrantId:    gid,
		AuthMethod: method,
	}
}

func identificationResponse(mpr *message.Printer, req *authpb.LoginRequest,
	err message.Reference) *authpb.LoginResponse {

	method := authpb.AuthMethod_AUTH_METHOD_EMAIL
	subheading := msgLoginSubheadingEmail
	toggleMessage := msgLoginUsePhone
	identifierInput := &components.Input{
		Type:        components.Input_EMAIL,
		Name:        "email",
		Placeholder: "wayne@boilerplate.pk",
		Value:       req.GetEmail(),
	}

	usePhone := req.AuthMethod == authpb.AuthMethod_AUTH_METHOD_PHONE
	if usePhone {
		method = authpb.AuthMethod_AUTH_METHOD_PHONE
		subheading = msgLoginSubheadingPhone
		toggleMessage = msgLoginUseEmail
		identifierInput = &components.Input{
			Type:        components.Input_PHONE,
			Name:        "phone",
			Placeholder: "+92 3xx xxxx xxx",
			Value:       req.GetPhone(),
		}
	}

	return &authpb.LoginResponse{
		Step: authpb.LoginStep_LOGIN_STEP_INITIAL,
		Screen: &components.Screen{
			Title: mpr.Sprintf(msgLoginTitle),
			Type:  components.ScreenType_SCREEN_LOGIN,
			ActionsTop: &components.ActionGroup{
				Actions: []*components.Action{
					{
						Type:     components.Action_TYPE_TOGGLE,
						Style:    components.Style_STYLE_PRIMARY,
						Element:  components.Element_ELEMENT_LINK,
						Text:     mpr.Sprintf(toggleMessage),
						Endpoint: components.Action_ENDPOINT_AUTH_LOGIN,
						Toggled:  usePhone,
					},
				},
				Position: &components.GroupPosition{
					PosX:    components.Position_POSITION_FILL,
					PosY:    components.Position_POSITION_START,
					IsRow:   true,
					Reverse: true,
				},
			},
			Content: &components.Screen_Login_Screen{
				Login_Screen: &components.LoginScreen{
					Form: &components.Form{
						Heading:    mpr.Sprintf(msgLoginHeading),
						Subheading: mpr.Sprintf(subheading),
						Groups: []*components.Form_Group{
							{
								Inputs: []*components.Input{
									identifierInput,
									{
										Type:        components.Input_PASSWORD,
										Name:        "password",
										Placeholder: "••••••••",
									},
								},
								Actions: []*components.Action{
									{
										Type:    components.Action_TYPE_REDIRECT,
										Style:   components.Style_STYLE_SECONDARY,
										Element: components.Element_ELEMENT_LINK,
										Text:    mpr.Sprintf(msgLoginStepResetPassword),
										Screen:  components.ScreenType_SCREEN_RESET_PASSWORD,
										Url:     "/password_reset",
									},
								},
								Position: &components.GroupPosition{
									PosX: components.Position_POSITION_FILL,
									PosY: components.Position_POSITION_START,
								},
							},
						},
						Error: mpr.Sprintf(err),
						Actions: []*components.Action{
							{
								Type:     components.Action_TYPE_ENDPOINT,
								Style:    components.Style_STYLE_PRIMARY,
								Element:  components.Element_ELEMENT_BUTTON,
								Text:     mpr.Sprintf(msgLoginStepSubmit),
								Endpoint: components.Action_ENDPOINT_AUTH_LOGIN,
							},
							{
								Type:    components.Action_TYPE_REDIRECT,
								Style:   components.Style_STYLE_SECONDARY,
								Element: components.Element_ELEMENT_BUTTON,
								Text:    mpr.Sprintf(msgLoginStepRegister),
								Screen:  components.ScreenType_SCREEN_SIGNUP,
								Url:     "/signup",
							},
						},
					},
				},
			},
		},
		AuthMethod: method,
	}
}

func success(w http.ResponseWriter, sessionId string) *authpb.LoginResponse {
	httpx.DefaultSessionCookie.SetSession(w, sessionId)

	return &authpb.LoginResponse{
		Step: authpb.LoginStep_LOGIN_STEP_SUCCESS,
		AuthSession: &authpb.AuthSession{
			Token:     sessionId,
			ExpiresAt: 0,
		},
		Redirect: &components.Action{
			Type:   components.Action_TYPE_REDIRECT,
			Screen: components.ScreenType_SCREEN_UNSPECIFIED,
			Url:    "/",
		},
	}
}

func Login(deps state.Deps) httpx.Handler {
	return func(w http.ResponseWriter, r *httpx.Request) {
		var reqpb authpb.LoginRequest
		if err := httpx.ParseJSON(r, &reqpb); err != nil {
			httpx.Fail(w, err)
			return
		}

		respb, err := loginHandler(r.Context(), deps,
			r.IP, r.UserAgent(), &reqpb, w)
		if err != nil {
			err = fmt.Errorf("could not login user: %v", err)
			if respb == nil {
				httpx.Fail(w, err)
				return
			}
			log.Error().Err(err).Msg("")
		}

		httpx.Ok(w, httpx.WithMessage(respb))
	}
}
