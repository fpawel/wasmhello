package datatype

type AccProfile struct {
	Acc string
}

type AccPass struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

type Miner struct {
	Name               string  `json:"name"`
	TerraBytes         string  `json:"terra-bytes"`
	Commit             string  `json:"commit"`
	Factor             float64 `json:"factor"`
	ContributorsNumber int     `json:"contributors-number"`
}
