package context

import (
	"fmt"
	"log"
	"os"
	"server/core/infrastructure"
	"server/external/postgresql"
	"server/external/s3"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type AppContext struct {
	Db                    sqlx.DB
	UserRepository        infrastructure.UserRepository
	ParaphraseRespository infrastructure.ParaphraseRespository
	//RatingRepository infrastructure.RatingRepository,
	Storage infrastructure.StorageClient
}

func initDb() sqlx.DB {

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

	return *db
}

func initStorge() infrastructure.StorageClient {
	sesh, err := session.NewSession()

	if err != nil {
		log.Fatalf("[InitStorage] Unable to create connection to S3 storage.")
	}

	return s3.New(sesh)
}

func InitAppContext() AppContext {
	db := initDb()
	storage := initStorge()

	return AppContext{
		Db:                    db,
		UserRepository:        postgresql.NewUserRepository(db),
		ParaphraseRespository: postgresql.NewParaphraseRepository(db),
		Storage:               storage,
	}
}
