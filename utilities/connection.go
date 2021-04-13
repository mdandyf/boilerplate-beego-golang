package utilities

import (
	"github.com/beego/beego/v2/server/web"
	"gorm.io/gorm"
)

var (
	// this one gonna contain an open connection
	// make sure to call Connect() before using it
	Conn *gorm.DB
)

// GetDB is connection to the Apps DB
func GetDB() *gorm.DB {
	dbURL, _ := web.AppConfig.String("dburl")
	dbPort, _ := web.AppConfig.String("dbport")
	dbUser, _ := web.AppConfig.String("dbuser")
	dbPass, _ := web.AppConfig.String("dbpass")
	dbName, _ := web.AppConfig.String("dbname")
	dbSSL, _ := web.AppConfig.String("dbssl")
	dbZone, _ := web.AppConfig.String("dbzone")
	dbType, _ := web.AppConfig.String("dbtype")

	switch dbType {
	case "postgresql":
		appsDB, err := connectPqSQL(
			dbURL, dbPort,
			dbUser, dbPass,
			dbName, dbSSL,
			dbZone,
		)

		if err != nil {
			panic("Error while connecting to App DB")
		}

		Conn = appsDB

		return appsDB
	case "mysql":
		appsDB, err := connectMySQL(
			dbURL, dbPort,
			dbUser, dbPass,
			dbName, dbSSL,
			"UTC",
		)

		if err != nil {
			panic("Error while connecting to App DB")
		}

		Conn = appsDB

		return appsDB

	case "mssql":
		appsDB, err := connectMsSQL(
			dbURL, dbPort,
			dbUser, dbPass,
			dbName, dbSSL,
			"UTC",
		)
		if err != nil {
			panic("Error while connecting to App DB")
		}

		Conn = appsDB

		return appsDB
	}

	return nil
}
