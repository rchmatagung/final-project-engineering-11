package model

type BookMentor struct {
	ID        int    `json:"id"`
	BookID    string `json:"book_id"`
	UserID    int    `json:"user_id"`
	MentorID  int    `json:"mentor_id"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

type BookListStatus struct {
	BookID     string `json:"book_id"`
	MentorName string `json:"mentor_name"`
	Status     string `json:"status"`
}
