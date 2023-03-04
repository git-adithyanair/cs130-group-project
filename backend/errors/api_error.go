package api_error

const (
	ErrUnknown               = "ERROR_UNKNOWN"
	ErrWrongPassword         = "ERROR_WRONG_PASSWORD"
	ErrNoUser                = "ERROR_NO_USER"
	ErrRegisterFail          = "ERROR_REGISTER_FAILURE"
	ErrCommunityExists       = "ERROR_COMMUNITY_EXISTS"
	ErrCommunityCreateFail   = "ERROR_COMMUNITY_CREATE_FAILURE"
	ErrNoCommunity           = "ERROR_NO_COMMUNITY"
	ErrAuthFail              = "ERROR_AUTH_FAILURE"
	ErrRequestNoItems        = "ERROR_REQUEST_NO_ITEMS"
	ErrNoStore               = "ERROR_NO_STORE"
	ErrNoRequest             = "ERROR_NO_REQUEST"
	ErrInvalidUserForRequest = "ERROR_INVALID_USER_FOR_REQUEST"
	ErrNoMember              = "ERROR_NO_MEMBER"
	ErrNoItem                = "ERROR_NO_ITEM"
	ErrErrandNoRequests      = "ERROR_ERRAND_NO_REQUESTS"
	ErrWrongCommunity        = "ERROR_WRONG_COMMUNITY"
	ErrNoErrand              = "ERROR_NO_ERRAND"
	ErrInvalidUserForErrand  = "ERROR_INVALID_USER_FOR_ERRAND"
	ErrTwilio                = "ERROR_TWILIO"
	ErrActiveErrorExists     = "ERROR_ACTIVE_ERROR_EXISTS"
	ErrRequestHasErrand      = "ERROR_REQUEST_HAS_ERRAND"
	ErrUserOwnsRequest       = "ERROR_USER_OWNS_REQUEST"
)

var GetErrorMessage map[string]string = map[string]string{
	ErrUnknown:               "An unknown error occured.",
	ErrWrongPassword:         "The password you entered is incorrect.",
	ErrNoUser:                "No user with that email exists.",
	ErrRegisterFail:          "Could not register user with provided details.",
	ErrCommunityExists:       "A community with that address already exists.",
	ErrCommunityCreateFail:   "Could not create a community with provided details.",
	ErrNoCommunity:           "Could not find this community.",
	ErrAuthFail:              "Could not authenticate user with provided details.",
	ErrRequestNoItems:        "Request must contain at least one item.",
	ErrNoStore:               "Could not find this store.",
	ErrNoRequest:             "Could not find this request.",
	ErrInvalidUserForRequest: "Request does not belong to user.",
	ErrNoMember:              "User does not exists in given community.",
	ErrNoItem:                "Could not find this item.",
	ErrErrandNoRequests:      "Errand must contain at least one request.",
	ErrWrongCommunity:        "Request does not belong to the given community.",
	ErrNoErrand:              "Could not find this errand.",
	ErrInvalidUserForErrand:  "Errand does not belong to user.",
	ErrTwilio:                "Text message failed to send.",
	ErrActiveErrorExists:     "User already has an active errand.",
	ErrRequestHasErrand:      "Request already belongs to another errand.",
	ErrUserOwnsRequest:       "Cannot add request to errand, request belongs to the current user.",
}
