package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"os"
	"todo_bygolang/config"

	"github.com/google/uuid"
	"github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

// sqlite3では使う
// const (
// 	tableNameUser    = "users"
// 	tableNameTodo    = "todos"
// 	tableNameSession = "sessions"
// )

// postgres版のinit関数
func init() {
	url := os.Getenv("DATABASE_URL") // herokuのURLの環境変数
	connection, _ := pq.ParseURL(url)
	connection += "sslmode=require"
	Db, err = sql.Open(config.Config.SQLDriver, connection)
	if err != nil {
		log.Println(err)
	}
}

// sqlite3版のinit関数
// func init() {
// 	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	// ユーザーのテーブル
// 	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		uuid STRING NOT NULL UNIQUE,
// 		name STIRNG,
// 		email STRING,
// 		password STRING,
// 		created_at DATETIME)`, tableNameUser)

// 	Db.Exec(cmdU)

// 	// Todoのテーブル
// 	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		content TEXT,
// 		user_id INTEGER,
// 		created_at DATETIME)`, tableNameTodo)

// 	Db.Exec(cmdT)

// 	// Sessionのテーブル
// 	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
// 		id INTEGER PRIMARY KEY AUTOINCREMENT,
// 		uuid STRING NOT NULL UNIQUE,
// 		email STRING,
// 		user_id INTEGER,
// 		created_at DATETIME)`, tableNameSession)

// 	Db.Exec(cmdS)

// }

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
