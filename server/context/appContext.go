package context

import (
	"fmt"
	"log"
	"os"
	"server/core/infrastructure"
	"server/external/postgresql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AppContext struct {
	Db                    sqlx.DB
	UserRepository        infrastructure.UserRepository
	ParaphraseRespository infrastructure.ParaphraseRespository
	//RatingRepository infrastructure.RatingRepository,
}

func InitAppContext() AppContext {

	var host string = os.Getenv("POSTGRES_HOST")
	var port string = os.Getenv("POSTGRES_PORT")
	var database string = os.Getenv("POSTGRES_DB_NAME")
	var username string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, username, password, database)
	db, err := sqlx.Connect("postgres", sqlInfo)
	if err != nil {
		log.Println(err)
		log.Fatalf("Could not connect to postgres instance")
	}
	return AppContext{
		Db:                    *db,
		UserRepository:        postgresql.NewUserRepository(*db),
		ParaphraseRespository: postgresql.NewParaphraseRepository(*db),
	}
}
