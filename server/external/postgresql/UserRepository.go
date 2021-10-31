package postgresql

import (
	"database/sql"
	"fmt"
	"log"
	"github.com/jmoiron/sqlx"
	"server/core/infrastructure"
);

type userRepository struct {
	db sqlx.Connect()

}

func New() infrastructure.UserRepository {
	
}
