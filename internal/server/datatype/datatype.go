package datatype

type AccProfile struct {
	Acc string
}

type UserPass struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

type Miner struct {
	Name               string  `json:"name"`
	TerraBytes         string  `json:"terra-bytes"`
	Commit             string  `json:"commit"`
	Factor             float64 `json:"factor"`
	ContributorsNumber int     `json:"contributors-number"`
}
