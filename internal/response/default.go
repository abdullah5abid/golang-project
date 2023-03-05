package response

type Response struct {
	Id      uint64 `json:"id,omitempty"`
	Status  uint   `json:"status"`
	Message string `json:"message"`
}
