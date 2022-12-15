package domain

type Dentists struct {
	Id           int    `json:"id"`
	Name         string `json:"name" binding:"required"`
	Lastname     string `json:"lastname" binding:"required"`
	Registration int    `json:"registration" binding:"required"`
	Email        string `json:"email"`
}
