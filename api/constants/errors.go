package constants

type Error struct {
	Message       string `json:"message"`
	PrettyMessage string `json:"prettyMessage"`
	Status        int    `json:"status"`
}

var (
	ErrorTokenInvalid = Error{
		Message:       "Token Invalid",
		PrettyMessage: "The token provided is invalid.",
		Status:        401,
	}
	ErrorEmailExists = Error{
		Message:       "Email Exists",
		PrettyMessage: "The email address is already in use.",
		Status:        409,
	}
	ErrorMissingFields = Error{
		Message:       "Missing Fields",
		PrettyMessage: "The request is missing required fields.",
		Status:        400,
	}
	ErrorBadRequest = Error{
		Message:       "Bad Request",
		PrettyMessage: "The request could not be understood by the server due to malformed syntax.",
		Status:        400,
	}
	ErrorUnauthorized = Error{
		Message:       "Unauthorized",
		PrettyMessage: "The request has not been applied because it lacks valid authentication credentials for the target resource.",
		Status:        401,
	}
	ErrorForbidden = Error{
		Message:       "Forbidden",
		PrettyMessage: "The server understood the request, but is refusing to fulfill it.",
		Status:        403,
	}
	ErrorNotFound = Error{
		Message:       "Not Found",
		PrettyMessage: "The server has not found anything matching the Request-URI.",
		Status:        404,
	}
	ErrorMethodNotAllowed = Error{
		Message:       "Method Not Allowed",
		PrettyMessage: "The method specified in the Request-Line is not allowed for the resource identified by the Request-URI.",
		Status:        405,
	}
	ErrorConflict = Error{
		Message:       "Conflict",
		PrettyMessage: "The request could not be completed due to a conflict with the current state of the resource.",
		Status:        409,
	}
	ErrorInternalServerError = Error{
		Message:       "Internal Server Error",
		PrettyMessage: "The server encountered an unexpected condition which prevented it from fulfilling the request.",
		Status:        500,
	}
	ErrorServiceUnavailable = Error{
		Message:       "Service Unavailable",
		PrettyMessage: "The server is currently unable to handle the request due to a temporary overloading or maintenance of the server.",
		Status:        503,
	}
)
