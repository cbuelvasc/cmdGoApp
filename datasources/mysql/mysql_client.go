package mysql

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var ClientDB *sql.DB

func OpenConnection() {
	errVar := godotenv.Load()
	if errVar != nil {
		panic(errVar)
	}

	conection, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_SERVER")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))

	if err != nil {
		panic(err)
	}
	ClientDB = conection
}

func CloseConnection() {
	ClientDB.Close()
}
