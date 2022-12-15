package main

import (
	"finalCheckpoint/cmd/server/handler"
	"finalCheckpoint/internal/appointments"
	"finalCheckpoint/internal/dentists"
	"finalCheckpoint/internal/patients"
	"finalCheckpoint/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	sqlStorage := store.NewSQLStore()

	dentistRepo := dentists.NewRepository(sqlStorage)
	dentistService := dentists.NewService(dentistRepo)
	dentistHandler := handler.NewDentistsHandler(dentistService)

	patientRepo := patients.NewPatientsRepository(sqlStorage)
	patientService := patients.NewPatientsService(patientRepo)
	patientHandler := handler.NewPatientsHandler(patientService)

	appointmentRepo := appointments.NewAppointmentsRepository(sqlStorage)
	appointmentService := appointments.NewAppointmentsService(appointmentRepo)
	appointmentHandler := handler.NewAppointmentsHandler(appointmentService)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentists := r.Group("/dentists")
	{
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.POST("", dentistHandler.Post())
		dentists.DELETE(":id", dentistHandler.Delete())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.PUT(":id", dentistHandler.Put())
	}
	patients := r.Group("/patients")
	{
		patients.GET(":id", patientHandler.GetByID())
		patients.POST("", patientHandler.Post())
		patients.DELETE(":id", patientHandler.Delete())
		patients.PATCH(":id", patientHandler.Patch())
		patients.PUT(":id", patientHandler.Put())
	}
	appointments := r.Group("/appointments")
	{
		appointments.GET(":id", appointmentHandler.GetByID())
		appointments.POST("", appointmentHandler.Post())
		appointments.DELETE(":id", appointmentHandler.Delete())
		appointments.PATCH(":id", appointmentHandler.Patch())
		appointments.PUT(":id", appointmentHandler.Put())
	}

	r.Run(":8080")
}
