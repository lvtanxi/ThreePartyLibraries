package main

import (
	"github.com/jinzhu/gorm"
	"time"
)

var users []User = []User{
	{Username: "foobar", FirstName: "Foo", LastName: "Bar", Salary: 200},
	{Username: "helloworld", FirstName: "Hello", LastName: "World", Salary: 200},
	{Username: "john", FirstName: "John", Salary: 200},
}

type User struct {
	gorm.Model
	Salary int

	// Set column type
	Username  string `sql:"type:VARCHAR(255)"`
	FirstName string

	// Set default value
	LastName string `sql:"DEFAULT:'Smith'"`

	// Ignored attribute will be treated as attr instead of column
	IgnoredField bool `sql:"-"`

	// Relationship
	Calendar     Calendar
	Appointments []Appointment `gorm:"many2many:appointment_user"`


}

type Calendar struct {
	ID           uint
	Name         string
	UserID       uint
	Appointments []Appointment `gorm:"polymorphic:owner"`
}

type Appointment struct {
	gorm.Model
	Subject     string
	Description string
	StartTime   time.Time
	Length      uint
	OwnerID     uint
	OwnerType   string
	Attendees   []User `gorm:"many2many:appointment_user"`
}

type TaskList struct {
	gorm.Model
	Appointments []Appointment `gorm:"polymorphic:owner"`
}

// NotFound checks if a record exists in the database
func (u *User) NotFound() bool {
	return u.Model.ID == 0
}
