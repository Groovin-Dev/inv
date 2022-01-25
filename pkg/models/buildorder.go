package models

type BuildOrder struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CreatedFor  int     `json:"createdFor"`
	Assets      []Asset `json:"assets"`
}
