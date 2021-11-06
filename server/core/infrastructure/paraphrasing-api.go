package infrastructure

import "server/core/models"

type ParaphrasingApi interface {
	RequestParaphrase(originalText string) (models.ParaphraseResponse, error)
}
