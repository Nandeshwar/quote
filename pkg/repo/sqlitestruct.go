package repo

import (
	"database/sql"
	"fmt"
	"quote/pkg/fileutil"
	"quote/pkg/newrelicwrapper"
)

type SQLite3Repo struct {
	DB *sql.DB
}

func NewSqlite3Repo(sqlite3FileName string) (SQLite3Repo, error) {
	txn := newrelicwrapper.StartTransaction("Connect To Cassandra")

	if !fileutil.FileExists(sqlite3FileName) {
		txn.NoticeError(fmt.Errorf("db connection failed. file=%s does not exist", sqlite3FileName))
		return SQLite3Repo{}, fmt.Errorf("db connection failed. file=%s does not exist", sqlite3FileName)
	}
	db, err := sql.Open("sqlite3", sqlite3FileName)
	if err != nil {
		return SQLite3Repo{}, fmt.Errorf("error opening file=%v, error=%v", sqlite3FileName, err)
	}
	txn.End()
	return SQLite3Repo{DB: db}, nil
}
