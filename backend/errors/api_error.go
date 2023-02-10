package api_error

const (
	ErrUnknown             = "ERROR_UNKNOWN"
	ErrWrongPassword       = "ERROR_WRONG_PASSWORD"
	ErrNoUser              = "ERROR_NO_USER"
	ErrRegisterFail        = "ERROR_REGISTER_FAILURE"
	ErrCommunityExists     = "ERROR_COMMUNITY_EXISTS"
	ErrCommunityCreateFail = "ERROR_COMMUNITY_CREATE_FAILURE"
	ErrNoCommunity         = "ERROR_NO_COMMUNITY"
	ErrAuthFail            = "ERROR_AUTH_FAILURE"
)

var GetErrorMessage map[string]string = map[string]string{
	ErrUnknown:             "An unknown error occured.",
	ErrWrongPassword:       "The password you entered is incorrect.",
	ErrNoUser:              "No user with that email exists.",
	ErrRegisterFail:        "Could not register user with provided details.",
	ErrCommunityExists:     "A community with that address already exists.",
	ErrCommunityCreateFail: "Could not create a community with provided details.",
	ErrNoCommunity:         "Could not find this community.",
	ErrAuthFail:            "Could not authenticate user with provided details.",
}
