package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/nvlhnn/go-plesir/database/seeders"
	"github.com/nvlhnn/go-plesir/model/domain"

	// "github.com/ydhnwb/golang_api/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func OpenConnection() *gorm.DB{

	ssl := "require"
	env := os.Getenv("env")
	if env != "production" {
		ssl = "disable"
		err := godotenv.Load()
		if err != nil {
			panic("failed to load env file")
		}		
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")


	// dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=%s TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName, ssl)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to create a connection to database")
	}

	fmt.Println("migrating...")

	
	seed, err := strconv.ParseBool(os.Getenv("SEEDING"))
	if err != nil {
		seed = false
	}

	log.Println(seed)


	if seed {
		err = db.Migrator().DropTable(&domain.Place{}, &domain.PlaceDays{}, &domain.Order{})
		if err != nil {
			log.Panicln(err)
		}
	}

	
	
	err = db.AutoMigrate(&domain.User{}, &domain.Place{}, &domain.Day{}, &domain.PlaceDays{}, &domain.Order{})

	if err == nil {
		seeders.SeedDay(db)  
		if seed {
			seeders.PlaceSeed(db)				
		}
	}
	db.AutoMigrate(&domain.User{}, &domain.Place{}, &domain.Day{})
	// db.SetupJoinTable(&domain.Place{}, "Days", &domain.PlaceDays{})

	// db.AutoMigrate(&entity.Book{}, &entity.User{})


	return db
}

func CloseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}