package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	RollNo uint
	Name   string
	Marks  uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Student{})

	// Create
	db.Create(&Student{RollNo: 423001, Name: "Tony Stark", Marks: 84})
	db.Create(&Student{RollNo: 201941, Name: "Captain America", Marks: 76})

	// Read
	var stu1, stu2 Student
	db.First(&stu1) // returns the first record

	db.Find(&stu2, "roll_no = ?", 201941) // find student with RollNo 201941

	// Update
	db.Model(&stu2).Update("Marks", 86) // update students's marks to 86

	//Delete
	db.Unscoped().Delete(&stu1) // Permanent deletion
	db.Delete(&stu2)            // Soft deletion
}
