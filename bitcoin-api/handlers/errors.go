package handlers

func BadRequestErrorResp(message string) *CommonErrorObject {
	return &CommonErrorObject{
		Message: message,
		Text:    "BAD_REQUEST",
	}
}
