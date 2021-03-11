package configSql

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func ConnectSqlServer() *gorm.DB {
	dsn := "sqlserver://sa:Password01.@localhost:1433?database=JWTAPP"
	// dsn := "sqlserver://sa:Password01.@localhost:1433?database=JWTAPP"
	database, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return database
}
