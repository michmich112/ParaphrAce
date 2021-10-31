package infrastructure

import "server/core/models"

type UserRepository interface {
	GetBySessionToken(token string) (models.User, error)
	Create(models.User) (models.User, error)
}

type ParaphraseRespository interface {
	GetById(id int64) (models.Paraphrase, error)
	Create(paraphrase models.Paraphrase) (models.Paraphrase, error)
	Update(paraphrase models.Paraphrase) (models.Paraphrase, error)
	AddRating(ratingId int64) (models.Paraphrase, error)
}

type RatingRepository interface {
	Create(rating models.Rating) (models.Rating, error)
}
