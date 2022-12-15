package handler

import (
	"errors"
	"os"
	"strconv"

	"finalCheckpoint/internal/dentists"
	"finalCheckpoint/internal/domain"
	"finalCheckpoint/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistsHandler struct {
	s dentists.Service
}

func NewDentistsHandler(s dentists.Service) *dentistsHandler {
	return &dentistsHandler{
		s: s,
	}
}

func (h *dentistsHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentists, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentists not found"))
			return
		}
		web.Success(c, 200, dentists)
	}
}

func validateEmptysDentist(dentists *domain.Dentists) (bool, error) {
	switch {
	case dentists.Name == "" || dentists.Lastname == "" || dentists.Email == "":
		return false, errors.New("fields can't be empty")
	case dentists.Registration <= 0:
		return false, errors.New("value must be greater than 0")
	}
	return true, nil
}

func (h *dentistsHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentists domain.Dentists
		err := c.ShouldBindJSON(&dentists)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentist(&dentists)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(dentists)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *dentistsHandler) Delete() gin.HandlerFunc {
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
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

func (h *dentistsHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentists not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentists domain.Dentists
		err = c.ShouldBindJSON(&dentists)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentist(&dentists)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, dentists)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *dentistsHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name         string `json:"name,omitempty"`
		Lastname     string `json:"lastname,omitempty"`
		Registration int    `json:"registration,omitempty"`
		Email        string `json:"email,omitempty"`
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
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentists not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentists{
			Name:         r.Name,
			Lastname:     r.Lastname,
			Registration: r.Registration,
			Email:        r.Email,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
