package context

import (
	"fmt"
	"log"
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

const (
	host     = ""
	port     = 0
	database = ""
	username = ""
	password = ""
)

func InitAppContext() AppContext {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", host, port, username, password, database)
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
