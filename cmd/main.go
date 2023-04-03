package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"shopChallenge/delivery"
	"shopChallenge/domain"
	"shopChallenge/repository"
	third_domain "shopChallenge/thirdparty/domain"
	"shopChallenge/thirdparty/party"
	"shopChallenge/usecase"
)

func main() {
	setUpConfig()

	db := getGormDb()
	repo := repository.NewRepo(db)

	thirdparyConfig := third_domain.Config{}
	err := viper.UnmarshalKey("thirdparty", &thirdparyConfig)
	if err != nil {
		panic("cannot unmarshal thirdparty config")
	}
	fmt.Println(thirdparyConfig, " !!!")
	sms := party.InitSMS(thirdparyConfig)

	useCase := usecase.NewUseCase(repo, sms)

	handler := delivery.NewDataHandler(useCase)
	gin.SetMode(viper.GetString("gin_mode"))
	executeDbMigrations()
	router := gin.New()
	setRoute(router, &handler)

	router.Run(":9001")
}

func setUpConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error in reading config file: %w \n", err))
	}
}

func executeDbMigrations() {
	log.Info("DB migrations is Running...")
	m, err := migrate.New(
		"file://db/migration",
		viper.GetString("pg.migrate_src"))

	if err != nil {
		log.Error(errors.Wrap(err, "failed to find migration files"))
		panic("failed to find migration files")
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error(errors.Wrap(err, "error on apply migrations"))
		panic("failed to apply migrations")
	}
}

func getGormDb() *gorm.DB {
	db, err := gorm.Open(postgres.Open(viper.GetString("pg.src")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		log.Error(errors.Wrap(err, "failed to initial gorm db"))
		panic("failed to initial gorm db")
	}
	return db
}

func setRoute(g *gin.Engine, handler *domain.DataHandler) {
	g.GET("/customer/active/transactions",
		(*handler).GetLatestTXNsOfMostActiveUsers)
	g.POST("/card/transfer", (*handler).TransferMoney)
}
