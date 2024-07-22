package db

import(
	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

var DB *sql.DB
func InitDB(){
	var err error
	DB ,err=sql.Open("sqlite3","api.db")
	if err!=nil{
		panic("Could not connect")
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()

}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS USERS(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	EMAIL TEXT NOT NULL UNIQUE,
	PASSWORD TEXT NOT NULL
	)`
	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("could not create users table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime DATETIME,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	)`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table")
	}

	createRegistrationsTable:=`
	CREATE TABLE IF NOT EXISTS REGISTRATIONS(
	ID INTEGER PRIMARY KEY AUTOINCREMENT,
	EVENT_ID INTEGER,
	USER_ID INTEGER,
	FOREIGN KEY(EVENT_ID) REFERENCES EVENTS(ID),
	FOREIGN KEY (USER_ID) REFERENCES USERS(ID) )`

	_,err=DB.Exec(createRegistrationsTable)
	if err != nil {
		panic("could not create events table")
	}
}
