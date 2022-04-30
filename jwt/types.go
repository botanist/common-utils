package jwt

type Jwt struct {
	OK    bool   `json:"ok"`
	Error string `json:"error"`
	Jwt   string `json:"jwt"`
}
