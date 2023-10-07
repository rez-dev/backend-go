package models

type Term struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Term_group int    `json:"term_group"`
}

type Terms []Term

type Test struct {
	Category     string `json:"category"`
	Description  string `json:"description"`
	Post_content string `json:"post_content"`
	Post_title   string `json:"post_title"`
	Meta_value   string `json:"meta_value"`
}

type Tests []Test
