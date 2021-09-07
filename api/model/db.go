package model

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Cinema struct {
	ID      int    `gorm:"column:id_cinema" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Address string `gorm:"column:address" json:"address"`
	Status  string `gorm:"column:status" json:"status"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name
func (Cinema) TableName() string {
	return "cinemas"
}

func newConnection() *gorm.DB {
	USER := os.Getenv("DB_USER")
	PASS := os.Getenv("DB_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)
	db, err := gorm.Open(mysql.Open(URL))
	if err != nil {
		panic("Failed to connect to database!")
	}
	log.Println("Database connection established")

	return db
}

func GetAllCinemas() []Cinema {

	dbConnection := newConnection()

	var cinemas []Cinema
	dbConnection.Where("status = ?", "enabled").Find(&cinemas)

	return cinemas
}

func FindCinema(cinemaID int) Cinema {
	dbConnection := newConnection()

	var cinema Cinema
	dbConnection.Where("id_cinema = ?", cinemaID).Find(&cinema)

	return cinema
}

func NewCinema(cinema Cinema) int {
	dbConnection := newConnection()
	dbConnection.Create(&cinema)

	return cinema.ID
}

func UpdateCinema(cinemaID int, cinema Cinema) {
	dbConnection := newConnection()

	var cinemaOrig Cinema
	dbConnection.Where("id_cinema = ?", cinemaID).Find(&cinemaOrig)

	cinemaOrig.Address = cinema.Address
	cinemaOrig.Name = cinema.Name
	cinemaOrig.Status = cinema.Status

	dbConnection.Save(&cinemaOrig)
}

func DeleteCinema(cinemaID int) {
	dbConnection := newConnection()

	var cinema Cinema
	dbConnection.Where("id_cinema = ?", cinemaID).Find(&cinema)
	dbConnection.Delete(&cinema)
}
