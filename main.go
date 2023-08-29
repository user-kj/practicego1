package main

import (
	"time"

	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DataField struct {
	ID          uint   `gorm:"primaryKey"`
	UUID        string `gorm:"unique"`
	Name        string
	Age         int
	DateOfBirth time.Time `gorm:"column:date_of_birth"`
	Email       string    `gorm:"unique"`
}

func main() {
	// Replace with your connection string.
	dsn := "root:@tcp(localhost:3306)/my_yoga_teacher?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatal("Failed to connect to database!", err)
	}

	// Getting the first record.
	var dataField DataField
	result := db.First(&dataField)

	var dataFieldUpdate DataField
	update_result := db.First(&dataFieldUpdate)

	if update_result.Error != nil {
		logrus.Error("Error reading record: ", result.Error)
	} else {
		logrus.Info("Record: ", dataFieldUpdate)
	}

	// Check for errors.
	if result.Error != nil {
		logrus.Error("Error reading record: ", result.Error)
	} else {
		logrus.Info("Record: ", dataField)
	}

	// data_enter := &DataField{56, "some-unique-id1", "John Doe", 25, time.Date(1998, 7, 5, 0, 0, 0, 0, time.UTC), "john.doe@example.com"}
	// result1 := write(db, data_enter)
	// if result1.Error != nil {
	// 	fmt.Printf("Errt: %e", result1.Error)
	// } else {
	// 	logrus.Info("Record: ", dataField)
	// }
	dataUpdate := &DataField{Email: "john.doe@exampleupdate.com"}

	update_result1 := update(db, dataUpdate)

	if update_result1.Error != nil {
		fmt.Println(update_result1.Error)
	} else {
		logrus.Info("Record: ", dataField)
	}

}

func write(db *gorm.DB, data *DataField) *gorm.DB {
	//dataField := DataField{
	//UUID:        "some-unique-id",
	//Name:        "John Doe",
	//Age:         25,
	//DateOfBirth: time.Date(1998, 7, 5, 0, 0, 0, 0, time.UTC),
	//Email:       "john.doe@example.com",
	//}

	result := db.Create(&data)
	return result
}

func update(db *gorm.DB, data *DataField) *gorm.DB {

	resultUpt := db.Where("name = 'Mike Johnson'").Updates(&data)
	return resultUpt

}
