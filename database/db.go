package database

import (
	"log"
	"os"

	"github.com/guilhermelribeiroo/go-continuousintegration/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	stringDeConexao := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}
