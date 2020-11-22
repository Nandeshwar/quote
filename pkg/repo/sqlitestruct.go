package repo

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"quote/pkg/fileutil"
	"quote/pkg/newrelicwrapper"
)

type SQLite3Repo struct {
	DB     *sql.DB
	GORMDB *gorm.DB
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

	gormDB, err := gorm.Open(sqlite.Open(sqlite3FileName), &gorm.Config{})
	if err != nil {
		return SQLite3Repo{}, fmt.Errorf("error connecting db file=%v using GORM, error=%v", sqlite3FileName, err)
	}

	txn.End()
	return SQLite3Repo{DB: db, GORMDB: gormDB}, nil
}
