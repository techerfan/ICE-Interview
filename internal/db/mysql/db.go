package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
}

type Mysql struct {
	db *gorm.DB
}

func New(config Config) *Mysql {

	// MySQL connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

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
