package constants

type Response struct {
	Message       string `json:"message"`
	PrettyMessage string `json:"prettyMessage"`
	Status        int    `json:"status"`
	Data          any    `json:"data,omitempty"`
}
