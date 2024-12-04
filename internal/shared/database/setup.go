package database

import (
	"fmt"
	"http-server/internal/shared/env"
	"log"

	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection(env *env.Env) *gorm.DB {
	DB, err := gorm.Open(postgres.Open(env.DB_URL), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}) // add env
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
		return nil
	}
	fmt.Println("Database connected!")
	pg, _ := DB.DB()
	pg.SetMaxOpenConns(20)
	return DB
}

var DBConnection = fx.Options(fx.Provide(NewDatabaseConnection))
