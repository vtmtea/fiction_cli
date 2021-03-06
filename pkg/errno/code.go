package errno

var (
	// OK Common error
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// ErrUserNotFound user errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
	ErrValidation   = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase     = &Errno{Code: 20002, Message: "Database error."}
	ErrToken        = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// ErrEncrypt user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	ErrGetBookId        = &Errno{Code: 20300, Message: "书本ID获取失败"}
	ErrGetBookChapterId = &Errno{Code: 20301, Message: "章节ID获取失败"}
	ErrGetSourceId      = &Errno{Code: 20302, Message: "来源ID获取失败"}
)
