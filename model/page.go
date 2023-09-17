package model

type Page struct {
	Title        string `json:"title"`
	Headers      string `json:"headers"`
	Signature    string `json:"signature"`
	Description  string `json:"description"`
	Example      string `json:"example"`
}
