package entity

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("se-64.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	database.AutoMigrate(
		// User
		&UserRole{},
		&UserLogin{},
		&Gender{},
		&UserPrefix{},
		&UserDetail{},
		&Member{},
		&Employee{},
		&EmployeePosition{},

		// ManageSalary

		// ManageWorkSchedule

		// ProductStock
		&Productstock{},
		&Typeproduct{},
		&Product{}, 
		
		// ManagePromotion
		&Employee{},
		&NamePromotion{},
		&PromotionPeriod{},
		&PromotionType{},
		&ManagePromotion{},

		// PremiumMember

		// PurchaseOrder
		&PaymentMethod{},
		&PurchaseOrder{},
		&PurchaseOrderItem{},
	)

	db = database
}
