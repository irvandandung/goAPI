package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"strings"
	"time"
	"github.com/joho/godotenv"
    "os"
)

func ConnectDB() *sql.DB {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
	   log.Fatal("Error loading .env file")
	}
	errorList := []string{}
	errorFormat := "ERROR - %s is not found in environment variable"
	HostDatabase := os.Getenv("HOST_DATABASE")
	if HostDatabase == "" {
		errorList = append(errorList, fmt.Sprintf(errorFormat, "HOST_DATABASE"))
	}
	PortDatabase := os.Getenv("PORT_DATABASE")
	if PortDatabase == "" {
		errorList = append(errorList, fmt.Sprintf(errorFormat, "PORT_DATABASE"))
	}
	UserDatabase := os.Getenv("USER_DATABASE")
	if UserDatabase == "" {
		errorList = append(errorList, fmt.Sprintf(errorFormat, "USER_DATABASE"))
	}
	PasswordDatabase := os.Getenv("PASSWORD_DATABASE")
	if PasswordDatabase == "" {
		errorList = append(errorList, fmt.Sprintf(errorFormat, "PASSWORD_DATABASE"))
	}
	NameDatabase := os.Getenv("NAME_DATABASE")
	if NameDatabase == "" {
		errorList = append(errorList, fmt.Sprintf(errorFormat, "NAME_DATABASE"))
	}
	if(len(errorList) != 0){
		errorMessage := strings.Join(errorList, "\n")
		log.Println(errorMessage)
		os.Exit(1)
	}
	db, err := sql.Open("mysql", UserDatabase+":"+PasswordDatabase+"@tcp("+HostDatabase+":"+PortDatabase+")/"+NameDatabase)
	if err != nil {
		log.Fatal(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}