package model

type Book struct {
	Name     string `json:"name"`
	Url      string `json:"url"`
	Kind     string `json:"kind"`
	Author   string `json:"author"`
	Platform string `json:"platform"`
	Info     string `json:"info"`
}
