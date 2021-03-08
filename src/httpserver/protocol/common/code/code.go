package code

const Success = "success"
const Fail = "error"
const AuthRequired = "auth_required"

var messageDict = map[string]string{
	Fail:         "API request error occurred.",
	AuthRequired: "Requiring auth. Please login at first.",
}

func GetMessage(code string) string {
	message, ok := messageDict[code]
	if ok {
		return message
	} else {
		return "Unknown error occurred."
	}
}
