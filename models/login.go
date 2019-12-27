package models

type loginError struct {
	msg string
}

type Login struct {
	phone string `json: "phone"`
	password string `json:"password"`
	msg loginError `json:"msg"`
}


func (l *Login) readError() *Login {
	return l
}

func (l *Login) writeError(c *Login) bool  {
	return true
}
