package services

import (
	"github.com/mercadolibre/CircuitBreaker - Federico Salas/src/api/domain"
	"github.com/mercadolibre/CircuitBreaker - Federico Salas/src/api/utils"
)

func GetSitesS() (*domain.Sites, *utils.ApiError) {
	sites := &domain.Sites{}

	if err := sites.Get(); err != nil {
		return nil, err
	}

	return sites, nil
}
