package handler

import (
	"errors"
	"os"
	"strconv"

	"finalCheckpoint/internal/domain"
	"finalCheckpoint/internal/patients"
	"finalCheckpoint/pkg/web"

	"github.com/gin-gonic/gin"
)

type patientsHandler struct {
	j patients.Service
}

func NewPatientsHandler(j patients.Service) *patientsHandler {
	return &patientsHandler{j: j}
}

func (h2 *patientsHandler) GetByID() gin.HandlerFunc {
	return func(c2 *gin.Context) {
		idParam2 := c2.Param("id")
		id2, err := strconv.Atoi(idParam2)
		if err != nil {
			web.Failure(c2, 400, errors.New("invalid id"))
			return
		}
		patients, err := h2.j.GetByID(id2)
		if err != nil {
			web.Failure(c2, 404, errors.New("patients not found"))
			return
		}
		web.Success(c2, 200, patients)
	}
}

func validateEmptys(patients *domain.Patients) (bool, error) {
	switch {
	case patients.Name == "" || patients.Lastname == "" || patients.Document == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func (h2 *patientsHandler) Post() gin.HandlerFunc {
	return func(c2 *gin.Context) {
		var patients domain.Patients
		err := c2.ShouldBindJSON(&patients)
		if err != nil {
			web.Failure(c2, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&patients)
		if !valid {
			web.Failure(c2, 400, err)
			return
		}
		p2, err := h2.j.Create(patients)
		if err != nil {
			web.Failure(c2, 400, err)
			return
		}
		web.Success(c2, 201, p2)
	}
}

func (h2 *patientsHandler) Delete() gin.HandlerFunc {
	return func(c2 *gin.Context) {
		token := c2.GetHeader("TOKEN")
		if token == "" {
			web.Failure(c2, 401, errors.New("token not found"))
			return
		}
		if token != os.Getenv("TOKEN") {
			web.Failure(c2, 401, errors.New("invalid token"))
			return
		}
		idParam := c2.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c2, 400, errors.New("invalid id"))
			return
		}
		err = h2.j.Delete(id)
		if err != nil {
			web.Failure(c2, 404, err)
			return
		}
		web.Success(c2, 204, nil)
	}
}

func (h2 *patientsHandler) Put() gin.HandlerFunc {
	return func(c2 *gin.Context) {
		idParam := c2.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c2, 400, errors.New("invalid id"))
			return
		}
		_, err = h2.j.GetByID(id)
		if err != nil {
			web.Failure(c2, 404, errors.New("patients not found"))
			return
		}
		if err != nil {
			web.Failure(c2, 409, err)
			return
		}
		var patients domain.Patients
		err = c2.ShouldBindJSON(&patients)
		if err != nil {
			web.Failure(c2, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&patients)
		if !valid {
			web.Failure(c2, 400, err)
			return
		}
		p, err := h2.j.Update(id, patients)
		if err != nil {
			web.Failure(c2, 409, err)
			return
		}
		web.Success(c2, 200, p)
	}
}

func (h *patientsHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name     string `json:"name,omitempty"`
		Lastname string `json:"lastname,omitempty"`
		Document string `json:"document,omitempty"`
		Reg_Date string `json:"reg_date,omitempty"`
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
		_, err = h.j.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patients not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Patients{
			Name:     r.Name,
			Lastname: r.Lastname,
			Document: r.Document,
			Reg_Date: r.Reg_Date,
		}
		p, err := h.j.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
