package server

import (
	"context"

	"github.com/google/uuid"
	genResult "github.com/ynoacamino/alert-go/gen/result"
	"github.com/ynoacamino/alert-go/internal/core/services"
)

type ResultEndPointService struct {
	results []*genResult.Result
	repo    *services.ResultService
}

func NewResultEndPointService(repo *services.ResultService) genResult.Service {
	return &ResultEndPointService{
		repo: repo,
	}
}

func (r *ResultEndPointService) ListResults(ctx context.Context, payload *genResult.ListResultsPayload) (*genResult.Resultlist, error) {
	results, err := r.repo.ListResult(ctx, int64(payload.Limit), int64(payload.Page))
	if err != nil {
		return nil, err
	}

	var res []*genResult.Result = make([]*genResult.Result, len(results))
	for i, result := range results {
		res[i] = &genResult.Result{
			ID:        result.ID,
			Status:    result.Status,
			CreatedAt: result.CreatedAt.String(),
			UpdateAt:  result.UpdatedAt.String(),
		}
	}

	return &genResult.Resultlist{
		Limit: payload.Limit,
		Page:  payload.Page,
		Total: len(res),
		Data:  res,
	}, nil
}

func (r *ResultEndPointService) CreateResult(ctx context.Context, payload *genResult.ResultPayload) (*genResult.Result, error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	result, err := r.repo.CreateResult(ctx, uuid.String(), *payload.Status)
	if err != nil {
		return nil, err
	}

	return &genResult.Result{
		ID:        result.ID,
		Status:    result.Status,
		CreatedAt: result.CreatedAt.String(),
		UpdateAt:  result.UpdatedAt.String(),
	}, nil
}
