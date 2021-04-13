package models

import (
	"database/sql"
	"time"
)

type User struct {
	ID           int            `db:"id" gorm:"primaryKey,autoIncrement" json:"id"`
	Username     string         `db:"username" json:"username"`
	Userpass     string         `db:"userpass" json:"userpass"`
	UserRealName string         `db:"name" json:"name"`
	UserNik      string         `db:"nik" json:"nik"`
	UserAddress  sql.NullString `db:"address" json:"address"`
	UserPosition sql.NullString `db:"position" json:"position"`
	CreatedDate  time.Time      `db:"created_date" sql:"DEFAULT:current_timestamp" json:"created_date"`
	CreatedBy    string         `db:"created_by" json:"created_by"`
	UpdatedDate  time.Time      `db:"updated_date" sql:"DEFAULT:current_timestamp" json:"updated_date"`
	UpdatedBy    string         `db:"updated_by" json:"updated_by"`
}
