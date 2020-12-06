package setting

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var err error
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable&search_path=%s", os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_DB"), "public")
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	loadSqlFile(db)
	return db
}

func loadSqlFile(db *sql.DB) {
    // Read file
    file, err := ioutil.ReadFile("./database.sql")
    if err != nil {
        fmt.Println(err.Error())
    }

    // Execute all
    _, err = db.Exec(string(file))
    if err != nil {
        fmt.Println(err.Error())
    }
}
