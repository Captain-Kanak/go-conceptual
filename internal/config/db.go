package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB(env *Env) *gorm.DB {
	dsn := env.DSN

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database Connected Successfully!")

	return db
}
