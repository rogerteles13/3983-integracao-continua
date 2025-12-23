package database

import (
	"log"
	"os"
	"time"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	endereco := os.Getenv("DB_HOST")
	usuario := os.Getenv("DB_USER")
	senha := os.Getenv("DB_PASSWORD")
	nomeBanco := os.Getenv("DB_NAME")
	portaBanco := os.Getenv("DB_PORT")
	dsn := "host=" + endereco + " user=" + usuario + " password=" + senha + " dbname=" + nomeBanco + " port=" + portaBanco + " sslmode=disable"

	maxAttempts := 20
	retryDelay := 3 * time.Second
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		DB, err = gorm.Open(postgres.Open(dsn))
		if err == nil {
			break
		}
		log.Printf("Tentativa %d/%d: erro ao conectar com banco: %v", attempt, maxAttempts, err)
		time.Sleep(retryDelay)
	}

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados após tentativas")
	}

	_ = DB.AutoMigrate(&models.Aluno{})
}
