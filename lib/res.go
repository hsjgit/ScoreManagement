package lib

type ResGetStudentScore struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
	Page  int         `json:"page"`
}

type ResUploadStudentScore struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
