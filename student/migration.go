package student

import "github.com/fiscaluno/pandorabox/db"

// Migrate migration entity BD
func Migrate() {
	db := db.Conn()
	defer db.Close()

	entity := new(Entity)

	entity.Name = "JC"
	entity.CourseInfo = CourseInfo{
		CourseID: 1,
	}

	// Migrate the schema
	db.AutoMigrate(&Entity{})

	// Create
	db.Create(entity)

	// Read
	// var entity Entity
	db.First(&entity, 1) // find entity with id 1
	// db.First(&entity, "name = ?", "JC") // find entity with name JC

	// Update - update entity's Name to SI
	db.Model(&entity).Update("Name", "SI")

	// Delete - delete entity
	db.Delete(&entity)
}
