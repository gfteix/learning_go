package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	config := mysql.Config{
		User:                 Envs.DBUser,
		Passwd:               Envs.DBPassword,
		Addr:                 Envs.DBAddress,
		DBName:               Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStorage(config)

	db, err := sqlStorage.Init()

	if err != nil {
		log.Fatalf("sqlStorage.Init failed! Error: %s", err)
	}

	store := NewStore(db)

	api := NewAPIServer(":3000", store)
	api.Serve()

}
