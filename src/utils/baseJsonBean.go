package utils


type BaseJsonBean struct {
	Code     int          `json:"error_code"`
	Message  string       `json:"error_msg"`
}

func NewBaseJsonBean() *BaseJsonBean {
	return &BaseJsonBean{}
}
