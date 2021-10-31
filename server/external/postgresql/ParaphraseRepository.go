package postgresql

import (
	"server/core/infrastructure"
	"server/core/models"

	"github.com/jmoiron/sqlx"
)

type ParaphraseRespository struct {
	db    sqlx.DB
	table string
}

func NewParaphraseRepository(db sqlx.DB) infrastructure.ParaphraseRespository {
	return ParaphraseRespository{
		db:    db,
		table: "paraphrase",
	}
}

func (r ParaphraseRespository) GetById(id int64) (models.Paraphrase, error) {
	paraphrase := models.Paraphrase{}
	err := r.db.Get(&paraphrase, "SELECT * FROM paraphrase WHERE id=$1", id)
	return paraphrase, err
}

func (r ParaphraseRespository) Create(paraphrase models.Paraphrase) (models.Paraphrase, error) {
	res, err := r.db.NamedExec(`INSERT INTO paraphrase (user_id, 
		rating_id, 
		timestamp,
		start_time, 
		end_time, 
		original_file_uri, 
		result_file_uri) VALUES (:user_id, 
		:rating_id, 
		:timestamp, 
		:start_time,
		:end_time, 
		:original_file_uri, 
		:result_file_uri)`, paraphrase)

	if err == nil {
		id, _ := res.LastInsertId() // Check compatibility with other databases
		uParaphrase, err := r.GetById(id)
		if err == nil {
			return uParaphrase, nil
		}
	}
	return paraphrase, err
}

func (r ParaphraseRespository) Update(paraphrase models.Paraphrase) (models.Paraphrase, error) {
	_, err := r.db.NamedExec(`UPDATE paraphrase SET
		user_id = :user_id,
		rating_id = :rating_id,
		timestamp = :timestamp,
		start_time = :start_time,
		end_time = :end_time,
		original_file_uri = :original_file_uri,
		result_file_uri = :result_file_uri
		WHERE id = :id`, paraphrase)
	if err == nil {
		uParaphrase, err := r.GetById(paraphrase.Id)
		if err == nil {
			return uParaphrase, nil
		}
	}
	return paraphrase, err
}

func (r ParaphraseRespository) AddRating(paraphraseId int64, ratingId int64) (models.Paraphrase, error) {
	_, err := r.db.Exec(`UPDATE paraphrase SET
		rating_id = $2 WHERE id = $1`, paraphraseId, ratingId)
	if err == nil {
		uParaphrase, err := r.GetById(paraphraseId)
		if err == nil {
			return uParaphrase, nil
		}
	}
	return models.Paraphrase{}, err

}
