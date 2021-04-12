package config

import (
	"fmt"
)

func ENGINE() string {
	return "postgres"
}

func DBCONFIG() string {
	host := "localhost"
	port := "5432"
	user := "postgres"
	pass := ""
	dbname := "golang"
	ssl := "disable"

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host,
		port,
		user,
		pass,
		dbname,
		ssl,
	)
}
