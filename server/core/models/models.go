package models

import (
	"database/sql"
	"time"
)

type User struct {
	Id           int64  `json:"id" db:"id"`
	SessionToken string `json:"session_token" db:"session_token"`
}

type Paraphrase struct {
	Id              int64          `json:"id" db:"id"`
	UserId          int64          `json:"user_id" db:"user_id"`
	RatingId        sql.NullInt64  `json:"rating_id" db:"rating_id"`
	Timestamp       time.Time      `json:"timestamp" db:"timestamp"`
	StartTime       sql.NullTime   `json:"start_time" db:"start_time"`
	EndTime         sql.NullTime   `json:"end_time" db:"end_time"`
	OriginalFileUri string         `json:"original_file_uri" db:"original_file_uri"`
	ResultFileUri   sql.NullString `json:"result_file_uri" db:"result_file_uri"`
}

type Rating struct {
	Id           int64          `json:"id" db:"id"`
	ParaphraseId int64          `json:"paraphrase_id" db:"paraphrase_id"`
	Timestamp    time.Time      `json:"timestamp" db:"timestamp"`
	Comment      sql.NullString `json:"comment" db:"comment"`
	Value        int16          `json:"value" db:"value"`
}
