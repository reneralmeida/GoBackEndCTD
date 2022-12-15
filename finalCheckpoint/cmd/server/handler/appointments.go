package handler

import (
	"errors"
	"os"
	"strconv"

	"finalCheckpoint/internal/appointments"
	"finalCheckpoint/internal/domain"
	"finalCheckpoint/pkg/web"

	"github.com/gin-gonic/gin"
)

type appointmentsHandler struct {
	x appointments.Service
}

func NewAppointmentsHandler(x appointments.Service) *appointmentsHandler {
	return &appointmentsHandler{
		x: x,
	}
}

func (h *appointmentsHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		appointments, err := h.x.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointments not found"))
			return
		}
		web.Success(c, 200, appointments)
	}
}

func validateEmptAppointments(appointments *domain.Appointments) (bool, error) {
	switch {
	case appointments.Description == "" || appointments.Date_And_Time == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func (h *appointmentsHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointments domain.Appointments
		err := c.ShouldBindJSON(&appointments)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptAppointments(&appointments)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.x.Create(appointments)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *appointmentsHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.x.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

func (h *appointmentsHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.x.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patients not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var appointments domain.Appointments
		err = c.ShouldBindJSON(&appointments)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptAppointments(&appointments)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.x.Update(id, appointments)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *appointmentsHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Description           string `json:"description,omitempty"`
		Date_And_Time         string `json:"date_and_time,omitempty"`
		Dentists_Registration int    `json:"dentists,omitempty"`
		Patients_Id           int    `json:"patients,omitempty"`
	}
	return func(c *gin.Context) {
		token := c.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c, 401, errors.New("invalid token"))
			return
		}
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.x.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointments not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Appointments{
			Description:           r.Description,
			Date_And_Time:         r.Date_And_Time,
			Dentists_Registration: r.Dentists_Registration,
			Patients_Id:           r.Patients_Id,
		}
		p, err := h.x.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
