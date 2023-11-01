package main

import (
	"WB_Intern_L0"
	"WB_Intern_L0/internal/handler"
	"WB_Intern_L0/internal/repo"
	"WB_Intern_L0/internal/usecase"
	"WB_Intern_L0/pkg/cache"
	natsService "WB_Intern_L0/pkg/nats"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"os"
)

func main() {
	zap.ReplaceGlobals(zap.Must(zap.NewProduction())) // Zap

	if err := initConfig(); err != nil {
		zap.L().Error("error initializing configs", zap.Error(err))
		return
	}

	if err := godotenv.Load(); err != nil {
		zap.L().Error("env error", zap.Error(err))
		return
	}

	db, err := repo.NewPostgresDB(repo.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	if err != nil {
		zap.L().Error("db initialize error", zap.Error(err))
		return
	}

	repository := repo.NewRepo(db)
	services := usecase.NewUseCase(repository)
	cacheService := cache.NewCache(services)
	handlers := handler.NewHandler(services, cacheService)

	// NATS Setup
	nc, err := nats.Connect(nats.DefaultURL)

	natsHandler := natsService.NewNats(services, nc, cacheService)

	nc.Subscribe("newOrder", func(m *nats.Msg) {
		natsHandler.CreateOrderNATS(m)
	})

	nc.Subscribe("getOrder", func(m *nats.Msg) {
		natsHandler.GetOrderByIdNATS(m)
	})

	if err != nil {
		zap.L().Error("nats connect error", zap.Error(err))
		return
	}

	// Cache Setup
	err = cacheService.LoadDataFromDatabase()
	if err != nil {
		zap.L().Error("cache reload error", zap.Error(err))
		return
	}

	// Web Setup
	srv := new(WB_Intern_L0.Server)
	if err := srv.Run(viper.GetString("ip")+":"+viper.GetString("port"), handlers.InitRoute()); err != nil {
		fmt.Println(err)
		zap.L().Error("Error starting server")
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
