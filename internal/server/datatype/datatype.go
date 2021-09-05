package datatype

type AccProfile struct {
	Acc string
}

type AccPass struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
