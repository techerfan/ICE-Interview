package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	db *gorm.DB
}

func New() *Mysql {

	// MySQL connection string
	// TODO: these attributes must be read from environment variables
	dsn := "ice_user:9xz3jrd8wf@tcp(localhost:4001)/ice_db?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Mysql{
		db: db,
	}
}

func (m *Mysql) MigrateDatabase() {
	// AutoMigrate will create or update the tables based on the models
	err := m.db.AutoMigrate(&Cart{}, &CartItem{})
	if err != nil {
		panic(err)
	}
}
