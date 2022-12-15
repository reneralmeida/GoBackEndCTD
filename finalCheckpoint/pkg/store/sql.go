package store

import (
	"database/sql"
	"fmt"
	"log"

	"finalCheckpoint/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type sqlStore struct {
	db *sql.DB
}

func NewSQLStore() StoreInterface {
	database, err := sql.Open("mysql", "root:@tcp(localhost:3306)/checkpoint_db")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStore{
		db: database,
	}
}

func (s *sqlStore) DeleteDentist(id int) error {
	_, err := s.db.Exec("DELETE FROM dentists WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlStore) ReadDentist(id int) (domain.Dentists, error) {
	dentist := domain.Dentists{}

	rows, err := s.db.Query(`SELECT * from dentists WHERE id=?`, id)
	if err != nil {
		return domain.Dentists{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&dentist.Id,
			&dentist.Name,
			&dentist.Lastname,
			&dentist.Registration,
			&dentist.Email,
		)
		if err != nil {
			return domain.Dentists{}, err
		}
	}
	return dentist, nil
}

func (s *sqlStore) UpdateDentist(dentist domain.Dentists) error {
	fmt.Println("updating dentist")
	_, err := s.db.Exec(
		`UPDATE dentists
		SET
			name = ?,
			lastname = ?,
			registration = ?,
			email = ?,
		WHERE id = ?`,
		dentist.Name,
		dentist.Lastname,
		dentist.Registration,
		dentist.Email,
		dentist.Id,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *sqlStore) ExistsDentist(email string) bool {
	return false
}

func (s *sqlStore) CreateDentist(dentist domain.Dentists) error {
	_, err := s.db.Exec(
		`INSERT INTO dentists
		(name, lastname, registration, email)
		VALUES
		(?, ?, ?, ?, ?, ?)`,
		dentist.Name,
		dentist.Lastname,
		dentist.Registration,
		dentist.Email,
	)
	if err != nil {
		return err
	}
	return nil
}

func (s2 *sqlStore) DeletePatient(id int) error {
	_, err := s2.db.Exec("DELETE FROM patients WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s2 *sqlStore) ReadPatient(id int) (domain.Patients, error) {
	patient := domain.Patients{}

	rows, err := s2.db.Query(`SELECT * from patients WHERE id=?`, id)
	if err != nil {
		return domain.Patients{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&patient.Id,
			&patient.Name,
			&patient.Lastname,
			&patient.Document,
			&patient.Reg_Date,
		)
		if err != nil {
			return domain.Patients{}, err
		}
	}
	return patient, nil
}

func (s2 *sqlStore) UpdatePatient(patient domain.Patients) error {
	fmt.Println("updating patient")
	_, err := s2.db.Exec(
		`UPDATE patients
		SET
			name = ?,
			lastname = ?,
			document = ?,
			reg_date = ?,
		WHERE id = ?`,
		patient.Name,
		patient.Lastname,
		patient.Document,
		patient.Reg_Date,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s2 *sqlStore) ExistsPatient(email string) bool {
	return false
}

func (s2 *sqlStore) CreatePatient(patient domain.Patients) error {
	_, err := s2.db.Exec(
		`INSERT INTO patients
		(name, lastname, document, reg_date)
		VALUES
		(?, ?, ?, ?, ?, ?)`,
		patient.Name,
		patient.Lastname,
		patient.Document,
		patient.Reg_Date,
	)
	if err != nil {
		return err
	}
	return nil
}

func (x *sqlStore) DeleteAppointment(id int) error {
	_, err := x.db.Exec("DELETE FROM appointments WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (x *sqlStore) ReadAppointment(id int) (domain.Appointments, error) {
	appointment := domain.Appointments{}

	rows, err := x.db.Query(`SELECT * from appointments WHERE id=?`, id)
	if err != nil {
		return domain.Appointments{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&appointment.Id,
			&appointment.Description,
			&appointment.Date_And_Time,
			&appointment.Dentists_Registration,
			&appointment.Patients_Id,
		)
		if err != nil {
			return domain.Appointments{}, err
		}
	}
	return appointment, nil
}

func (x *sqlStore) UpdateAppointment(appointment domain.Appointments) error {
	fmt.Println("updating appointment")
	_, err := x.db.Exec(
		`UPDATE appointments
		SET
			description = ?,
			date_and_time = ?,
			dentists_registration = ?,
			patients_id = ?,
		WHERE id = ?`,
		appointment.Description,
		appointment.Date_And_Time,
		appointment.Dentists_Registration,
		appointment.Patients_Id,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (x *sqlStore) ExistsAppointment(description string) bool {
	return false
}

func (x *sqlStore) CreateAppointment(appointment domain.Appointments) error {
	_, err := x.db.Exec(
		`INSERT INTO appointments
		(description, date_and_time, dentists_registration, patients_id)
		VALUES
		(?, ?, ?, ?, ?, ?)`,
		appointment.Description,
		appointment.Date_And_Time,
		appointment.Dentists_Registration,
		appointment.Patients_Id,
	)
	if err != nil {
		return err
	}
	return nil
}