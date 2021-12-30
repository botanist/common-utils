package serviceutils

type Jwt struct {
	OK    bool   `json:"bool"`
	Error string `json:"error"`
	Jwt   string `json:"jwt"`
}
