package infrastructure

import "server/core/models"

type ParaphrasingApi interface {
	RequestParaphrase(id string, originalText string) (models.Paraphrase, error)
}
