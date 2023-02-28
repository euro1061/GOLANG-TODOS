package main

import (
	"GOLANG-TODOS/handler"
	"GOLANG-TODOS/repository"
	"GOLANG-TODOS/service"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("sql: %v\n=======================================\n", sql)
}

func main() {
	initConfig()
	db := initDB()
	db.AutoMigrate(&repository.Todo{})
	app := fiber.New()

	todoRepo := repository.NewTodoRepositoryDB(db)
	todoService := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoService)

	app.Get("/todos", todoHandler.GetTodos)
	app.Get("/todos/:id", todoHandler.GetTodo)
	app.Post("/todos", todoHandler.NewTodo)
	app.Put("/todos", todoHandler.UpdateTodo)
	app.Delete("/todos/:id", todoHandler.DeleteTodo)

	err := app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
	if err != nil {
		panic(err)
	}
}

func initConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=%v TimeZone=%v",
		viper.GetString("db.host"),
		viper.GetString("db.user"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
		viper.GetInt("db.port"),
		viper.GetString("db.sslmode"),
		viper.GetString("db.timezone"),
	)

	dial := postgres.Open(dsn)
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
	})

	if err != nil {
		panic(err)
	}

	return db
}
