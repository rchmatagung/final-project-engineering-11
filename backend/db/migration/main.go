package migration

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func Migrate() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	dbname := os.Getenv("DB_NAME")
	db, err := sql.Open("sqlite3", `db/migration/`+dbname+`?_loc=Local`)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id integer not null primary key AUTOINCREMENT,
		username varchar(255) not null,
		name varchar(255) not null,
		password varchar(255) not null,
		role varchar(255) not null,
		address varchar(255) not null,
		phone varchar(255) not null,
		email varchar(255) not null,
		created_at datetime not null
);

CREATE TABLE IF NOT EXISTS mentor (
    id integer not null primary key AUTOINCREMENT,
	skill varchar(255) not null,
	bio varchar(255) not null,
	name varchar(255) not null,
	address varchar(255) not null,
	phone varchar(255) not null,
	email varchar(255) not null,
	created_at datetime not null,
	image varchar(255) not null
);
CREATE TABLE IF NOT EXISTS artikel (
    id integer not null primary key AUTOINCREMENT,
	judul varchar(255) not null,
	content varchar(255) not null,
	created_at datetime not null
);

CREATE TABLE IF NOT EXISTS bookmentor (
	id integer not null primary key AUTOINCREMENT,
	member_id integer not null,
	mentor_id integer not null,
	bookid varchar(255) not null,
	status varchar(255) not null,
	created_at datetime not null
);


`)

	if err != nil {
		panic(err)
	}
}
