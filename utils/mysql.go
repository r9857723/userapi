package utils

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type DataBase struct {
	Conn *gorm.DB
}

var db *DataBase

func init() {
	log.Println("connect...")
	config := GetConfig()
	db = connect(config)
}

func GetDb() *DataBase {
	return db
}

func connect(c *Config) *DataBase {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=True",
		c.UserName,
		c.Password,
		c.HostName,
		c.Port,
		c.DataBase)
	conn, err := gorm.Open(mysql.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect mysql, err = ", err)
		return nil
	}
	db, _ := conn.DB()
	db.SetMaxOpenConns(c.MaxOpenConn)
	db.SetMaxIdleConns(c.MaxIdleConn)
	db.SetConnMaxIdleTime(time.Duration(c.MaxLifeTime) * time.Second)
	return &DataBase{Conn: conn}
}

func (db *DataBase) Insert(entity interface{}) (err error) {
	result := db.Conn.Debug().Create(entity)
	if result.Error != nil {
		err = errors.New("Failed to Insert data")
		return
	}
	if result.RowsAffected != 1 {
		err = errors.New("Failed to rowAffected number")
		return
	}
	log.Printf("%d \n", result.RowsAffected)
	return
}

func (db *DataBase) QueryByAny(entity interface{}, query interface{}, args ...interface{}) {
	db.Conn.Debug().Where(query, args).Find(entity)
}

func (db *DataBase) DeleteById(entity interface{}, query interface{}, args ...interface{}) {
	db.Conn.Debug().Where(query, args).Delete(entity)
}

func (db *DataBase) Update(entity interface{}) {
	db.Conn.Debug().Save(entity)
}

