package database

import (
	"os"

	"github.com/domjesus/api-go-gin/models"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados(l *zap.SugaredLogger) error {

	// stringDeConexao := os.Getenv("DATABASE_URL")
	err := godotenv.Load()
	if err != nil {
		l.Error("Error loading .env file")
	}

	var sslmode string

	env := os.Getenv("ENV")

	if env == "local" {
		sslmode = "sslmode=disable"
	} else {
		sslmode = "sslmode=require"
	}

	var stringDeConexao string
	// fmt.Println("Ambiente: ", env)

	stringDeConexao = "host=" + os.Getenv("DATABASE_HOST") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PASSWORD") + " dbname=" + os.Getenv("DATABASE_NAME") + " port=5432 " + sslmode

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		return err
	}
	DB.AutoMigrate(&models.Aluno{})

	l.Info("DB connected")

	return nil
}
