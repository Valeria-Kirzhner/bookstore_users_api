package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

const (
	mysql_users_username = "mysql_users_username"
	mysql_users_password = "mysql_users_password"
	mysql_users_host = "mysql_users_host"
	mysql_users_schema = "mysql_users_schema"
)

var (
	Client *sql.DB

	username string
	password string
	host string
	schema string
)

// the init function is called by first time this users_db package is called.
func init(){

	getEnvVariables()
	fmt.Printf("godotenv : %s = %s \n", "mysql_users_host", host)

	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, host, schema) // the layout i should use in order to connect the db (user, password, host)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)// I don't want to start the application if the db connection is not succeeded
	}
	if err = Client.Ping(); err != nil {
		panic(err)// Check if it really works.
	}
	log.Println("database succesfully configured")
}

func getEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	  }
	 host = os.Getenv(mysql_users_host)
	 schema = os.Getenv(mysql_users_schema)
	 username = os.Getenv(mysql_users_username)
	 password = os.Getenv(mysql_users_password)
}