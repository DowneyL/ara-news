package response

type ErrorCode int64

type Errors struct {
	Code    int64
	Message string
}

type Status struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
	Errors  []string  `json:"errors,omitempty"`
	Time    string    `json:"time"`
}

type Response struct {
	Status Status      `json:"status"`
	Data   interface{} `json:"data"`
}
