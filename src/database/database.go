package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var dblogger *log.Logger = log.New(os.Stdout, "Database: ", log.LstdFlags|log.Lmsgprefix)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	Path     string // Sql .db file`s path
	Driver   string // Sql driver`s name
}

type DB struct {
	db     *sql.DB
	config Config
}

func envDBConfig() Config {
	return Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Path:     os.Getenv("DB_PATH"),
		Driver:   os.Getenv("DB_DRIVER"),
	}
}

var db *DB // Global object

func New(config ...Config) *DB {
	conn := &DB{
		config: Config{},
	}

	if len(config) > 0 {
		conn.config = config[0]
	} else { // if no config was provided then env confing used
		conn.config = envDBConfig()
	}

	if db != nil {
		dblogger.Println("")
	}

	db = conn
	return conn
}

func GetDB() *DB {
	if db == nil {
		dblogger.Panicln("db object is nil. Initilize db via database.New()!")
	}
	return db
}

func (db *DB) Connect() error {
	_db, err := sql.Open(db.config.Driver, db.config.Path)
	db.db = _db

	//! TEMPORARY
	f, err := os.ReadFile("D:\\Files\\Projects\\GO\\bid2bless-auction\\db\\scripts\\init_db.sql")
	if err != nil {
		dblogger.Fatalln(err)
	}
	_, err = _db.Exec(string(f))
	if err != nil {
		dblogger.Fatalln(err)
	}
	//! END TEMPORARY

	if err != nil {
		dblogger.Println("DB connection error: ", err)
		return err
	}
	err = _db.Ping()
	if err != nil {
		dblogger.Println("DB ping error: ", err)
		return err
	}

	dblogger.Println("Successfully connected to database!")
	return nil
}

func (db *DB) RawQuery(query string) error {
	_, err := db.db.Query("")
	return err
}

func (db *DB) Close() {
	db.db.Close()
}
