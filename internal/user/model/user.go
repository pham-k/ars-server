package model

const (
	ObjUser       = "user"
	PIDPrefixUser = "usr"
)

type User struct {
	PID    string `json:"pid"`
	Object string `json:"object"`
	Phone  string `json:"phone"`
}
