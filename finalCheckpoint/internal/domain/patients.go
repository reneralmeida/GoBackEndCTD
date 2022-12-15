package domain

type Patients struct {
	Id       int    `json:"id"`
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Document string `json:"document" binding:"required"`
	Reg_Date string `json:"reg_date" binding:"required"`
}
