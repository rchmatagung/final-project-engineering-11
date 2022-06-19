package model

type ArtikelList struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

type ArtikelDetail struct {
	Title   string `json:"judul"`
	Date    string `json:"date"`
	Content string `json:"content"`
}

type Artikel struct {
	Title   string `json:"judul"`
	Content string `json:"content"`
}
