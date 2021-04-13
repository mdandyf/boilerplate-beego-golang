package utilities

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// ConnectPqSQL to PostgreSQL Using GORM
func connectPqSQL(dburl string, dbport string, dbuser string, dbpass string, dbname string, dbssl string, dbtime string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s", dburl, dbport, dbuser, dbpass, dbname, dbssl, dbtime),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Println(fmt.Sprintf("%s %s", "Connection failed to PostgreSQL ", err.Error()))
		return nil, err
	}

	return db, nil
}

// ConnectMsSQL to SQL Server Using GORM
func connectMsSQL(dburl string, dbport string, dbuser string, dbpass string, dbname string, dbssl string, dbtime string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", dbuser, dbpass, dburl, dbport, dbname)
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(fmt.Sprintf("%s %s", "Connection failed to SQL Server ", err.Error()))
		return nil, err
	}

	return db, nil
}

// ConnectMySQL to My Server Using GORM
func connectMySQL(dburl string, dbport string, dbuser string, dbpass string, dbname string, dbssl string, dbtime string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpass, dburl, dbport, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Println(fmt.Sprintf("%s %s", "Connection failed to SQL Server ", err.Error()))
		return nil, err
	}

	return db, nil
}
