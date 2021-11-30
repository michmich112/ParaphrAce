package context

import (
	"fmt"
	"log"
	"os"
	"server/core/infrastructure"
	aimodel "server/external/ai-model"
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
	WithStorage           bool
	Storage               infrastructure.StorageClient
	ParaphraseApi         infrastructure.ParaphrasingApi
}

func initDb() sqlx.DB {

	var host string = os.Getenv("POSTGRES_HOST")
	var port string = os.Getenv("POSTGRES_PORT")
	var database string = os.Getenv("POSTGRES_DB_NAME")
	var username string = os.Getenv("POSTGRES_USER")
	var password string = os.Getenv("POSTGRES_PASSWORD")
	var ssl string = os.Getenv("POSTGRES_SSL")

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, username, password, database)
	if ssl != "" {
		sqlInfo = fmt.Sprintf("%s sslmode=%s", sqlInfo, ssl)
	}
	db, err := sqlx.Connect("postgres", sqlInfo)
	if err != nil {
		log.Println(err)
		log.Fatalf("Could not connect to postgres instance")
	}

	return *db
}

func initStorge() (infrastructure.StorageClient, bool) {
	sesh, err := session.NewSession()

	if err != nil {
		log.Println("[InitStorage] Unable to create connection to S3 storage. Continuing without storage.")
		return s3.S3Client{}, false
	}

	return s3.New(sesh), true
}

func initAiApi() infrastructure.ParaphrasingApi {
	return aimodel.New()
}

func InitAppContext() AppContext {
	db := initDb()
	storage, withStorage := initStorge()
	paraphraseApi := initAiApi()

	return AppContext{
		Db:                    db,
		UserRepository:        postgresql.NewUserRepository(db),
		ParaphraseRespository: postgresql.NewParaphraseRepository(db),
		WithStorage:           withStorage,
		Storage:               storage,
		ParaphraseApi:         paraphraseApi,
	}
}
