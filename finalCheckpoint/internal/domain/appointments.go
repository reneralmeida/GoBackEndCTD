package domain

type Appointments struct {
	Id                    int    `json:"id"`
	Description           string `json:"description" binding:"required"`
	Date_And_Time         string `json:"date_and_time" binding:"required"`
	Dentists_Registration int    `json:"dentists_registration" binding:"required"`
	Patients_Id           int    `json:"patients_id" binding:"required"`
}
