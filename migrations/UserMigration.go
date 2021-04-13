package migrations

import (
	"boilerplate-golang/models"
	"database/sql"

	"gorm.io/gorm"
)

func MigrateUserDB(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func SeedUserDB(db *gorm.DB) {
	var users = []models.User{
		{Username: "husniana", UserRealName: "Husnia Nur Annisa", Userpass: "password1234", UserNik: "H1234", UserAddress: sql.NullString{"SG2", true}, UserPosition: sql.NullString{"Marketing Manager", true}, CreatedBy: "GORM", UpdatedBy: "GORM"},
		{Username: "mdandy_f", UserRealName: "Mochamad Dandy Firmansyah", Userpass: "password4321", UserNik: "M3953", UserAddress: sql.NullString{"SG2", true}, UserPosition: sql.NullString{"Tech Lead", true}, CreatedBy: "GORM", UpdatedBy: "GORM"},
	}
	db.Create(&users)
}
