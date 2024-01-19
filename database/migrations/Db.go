package migrations

import (
	"GOproject/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	// Replace with your actual MySQL connection details
	connection := "root@tcp(localhost:3306)/GObackend?charset=utf8mb4&parseTime=True&loc=Local"

	// Create a MySQL dialector with the connection string
	dialector := mysql.Open(connection)

	// Use gorm.Open with the dialector
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Customer{}, &models.Address{}) // Add other models as needed
	if err != nil {
		return err
	}

	// Additional migrations if needed

	return nil
}
