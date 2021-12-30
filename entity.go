package serviceutils

type Entity struct {
	Ok       bool   `json:"ok"`
	Error    bool   `json:"error"`
	ID       string `json:"id"`
	Site     string `json:"site"`
	Type     string `json:"type"`
	ParentID string `json:"parent"`
}
