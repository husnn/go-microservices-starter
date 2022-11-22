package auth

import (
	"boilerplate/gateway/httpx"
	"boilerplate/gateway/state"
	authpb "boilerplate/proto/auth"
	"boilerplate/proto/components"
	"boilerplate/users"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang.org/x/text/message"
	"net/http"
)

func Signup(d state.Deps, userType users.UserType) httpx.Handler {
	return func(w http.ResponseWriter, r *httpx.Request) {
		var reqpb authpb.SignupRequest
		if err := httpx.ParseJSON(r, &reqpb); err != nil {
			httpx.Fail(w, err)
			return
		}

		ctx := r.Context()
		mpr := message.NewPrinter(r.Lang)

		if reqpb.Step == authpb.SignupStep_SIGNUP_STEP_UNSPECIFIED {
			httpx.Ok(w, httpx.WithMessage(signupResponse(mpr, &reqpb, "")))
			return
		}

		uid, err := d.UsersClient().Signup(ctx, userType,
			reqpb.GetEmail(), reqpb.GetPhone(), reqpb.Password, r.IP)
		if err != nil {
			errMsg := msgSignupErrGeneral

			if errors.Is(err, users.ErrInvalidEmail) {
				errMsg = msgLoginErrInvalidEmail
			} else if errors.Is(err, users.ErrInvalidPhone) {
				errMsg = msgLoginErrInvalidPhone
			} else if errors.Is(err, users.ErrInvalidPassword) {
				errMsg = msgLoginErrInvalidPassword
			} else if errors.Is(err, users.ErrUserAlreadyExists) {
				errMsg = msgSignupErrUserExists
			}

			err = fmt.Errorf("could not signup user: %v", err)
			log.Error().Err(err).Msg("")

			httpx.Ok(w, httpx.WithMessage(signupResponse(mpr,
				&reqpb, mpr.Sprintf(errMsg))))
			return
		}

		sid, err := d.AuthClient().LoginUnsafe(ctx, uid, r.IP, r.UserAgent())
		if err != nil {
			httpx.Fail(w, fmt.Errorf(
				"could not login newly signed up user: %v", err))
			return
		}

		httpx.DefaultSessionCookie.SetSession(w, sid)

		res := &authpb.SignupResponse{
			Step: authpb.SignupStep_SIGNUP_STEP_SUCCESS,
			AuthSession: &authpb.AuthSession{
				Token:     sid,
				ExpiresAt: 0,
			},
			Redirect: &components.Action{
				Screen: components.ScreenType_SCREEN_UNSPECIFIED,
				Url:    "/",
			},
		}

		httpx.Ok(w, httpx.WithMessage(res))
	}
}

func signupResponse(mpr *message.Printer, req *authpb.SignupRequest, err message.Reference) *authpb.SignupResponse {
	method := authpb.AuthMethod_AUTH_METHOD_EMAIL
	subheading := msgSignupSubheadingEmail
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
		subheading = msgSignupSubheadingPhone
		toggleMessage = msgLoginUseEmail
		identifierInput = &components.Input{
			Type:        components.Input_PHONE,
			Name:        "phone",
			Placeholder: "+44 xxx xxxx xxx",
			Value:       req.GetPhone(),
		}
	}

	return &authpb.SignupResponse{
		Step: authpb.SignupStep_SIGNUP_STEP_INITIAL,
		Screen: &components.Screen{
			Type: components.ScreenType_SCREEN_SIGNUP,
			ActionsTop: &components.ActionGroup{
				Actions: []*components.Action{
					{
						Type:     components.Action_TYPE_TOGGLE,
						Style:    components.Style_STYLE_PRIMARY,
						Element:  components.Element_ELEMENT_LINK,
						Text:     mpr.Sprintf(toggleMessage),
						Endpoint: components.Action_ENDPOINT_AUTH_SIGNUP,
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
			Content: &components.Screen_SignupScreen{
				SignupScreen: &components.SignupScreen{
					Form: &components.Form{
						Heading:    mpr.Sprintf(msgSignupHeading),
						Subheading: mpr.Sprintf(subheading),
						Groups: []*components.Form_Group{
							{
								Inputs: []*components.Input{
									identifierInput,
									{
										Type:        components.Input_PASSWORD,
										Name:        "password",
										Placeholder: "••••••••",
										Value:       "",
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
								Text:     mpr.Sprintf(msgSignupCreateAccount),
								Endpoint: components.Action_ENDPOINT_AUTH_SIGNUP,
							},
						},
						Footer: &components.Text{
							Type:      components.Text_TYPE_BODY,
							Alignment: components.Position_POSITION_MIDDLE,
							Value:     mpr.Sprintf(msgSignupFooter),
						},
					},
				},
			},
		},
		AuthMethod: method,
	}
}
