package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"

	// driver dialect for gorm
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"go.elastic.co/apm/module/apmgorm"
)

// GormDB is struct for Gorm Connection
type GormDB struct{}

func (d *GormDB) buildConnection() (*gorm.DB, error) {
	driver := fmt.Sprint(os.Getenv("driver"))

	switch driver {
	case "mysql":
		return d.mysqlConnection()
	default:
		return nil, errors.New("invalid database driver, please use mysql or sqlite")
	}
}

func (d *GormDB) mysqlConnection() (*gorm.DB, error) {
	usernameAndPassword := fmt.Sprint(os.Getenv("db_user")) + ":" + fmt.Sprint(os.Getenv("db_password"))
	hostName := "tcp(" + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + ")"

	log.Println("Connecting to DB Server " + fmt.Sprint(os.Getenv("db_host")) + ":" + fmt.Sprint(os.Getenv("db_port")) + "...")
	urlConnection := usernameAndPassword + "@" + hostName + "/" + fmt.Sprint(os.Getenv("db_database")) +
		"?charset=utf8&parseTime=true&loc=UTC"

	db, err := apmgorm.Open("mysql", urlConnection)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	log.Println("DB Server is connected!")

	return db, nil
}

// GetGormConnection create connection for gorm
func GetGormConnection() (*gorm.DB, error) {
	db := GormDB{}
	return db.buildConnection()
}
