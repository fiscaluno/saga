package student

import (
	"github.com/fiscaluno/pandorabox/db"
)

// Student is a Entity
type Student struct {
	Name       string     `json:"student_name"`
	UserID     int        `json:"user_id"`
	CourseInfo CourseInfo `json:"course_info"`
}

// CourseInfo ...
type CourseInfo struct {
	CourseID            int    `json:"course_id"`
	ReviewID            int    `json:"review_id"`
	InstitutionID       int    `json:"institution_id"`
	Period              string `json:"period"`
	Semester            int    `json:"semester"`
	MonthlyPaymentValue int    `json:"monthly_payment_value"`
}

// Entity is a course
type Entity struct {
	Student
	db.CommonModelFields
}

// Entitys is Entity slice
type Entitys []Entity

// GetAll Entitys
func GetAll() Entitys {
	db := db.Conn()
	defer db.Close()
	var entitys Entitys
	db.Find(&entitys)
	return entitys
}

// Save a Entity
func (entity Entity) Save() (Entity, error) {
	db := db.Conn()
	defer db.Close()

	db.Create(&entity)

	return entity, nil
}

// GetByID a Entity
func GetByID(id int) Entity {
	db := db.Conn()
	defer db.Close()

	var entity Entity

	db.Find(&entity, id)

	return entity
}

// GetByQuery a Entity
func GetByQuery(query string, value interface{}) Entitys {
	db := db.Conn()
	defer db.Close()

	var entitys Entitys

	db.Find(&entitys, query, value)
	return entitys
}
