package db

import (
	"context"
	"database/sql"

	"github.com/ynoacamino/alert-go/gen/db"
	"github.com/ynoacamino/alert-go/internal/core/domain"
	"github.com/ynoacamino/alert-go/internal/core/ports"
)

type ResultRepository struct {
	query *db.Queries
}

func NewResultRepository(sqlDB *sql.DB) ports.ResultRepository {
	queries := db.New(sqlDB)

	return &ResultRepository{
		query: queries,
	}
}

func (r *ResultRepository) Save(ctx context.Context, result *domain.Result) (*domain.Result, error) {
	row, err := r.query.CreateResult(ctx, db.CreateResultParams{
		ID:     result.ID,
		Status: result.Status,
	})
	if err != nil {
		return nil, mapSQLErrorResult(err)
	}

	saveResult := &domain.Result{
		ID:        row.ID,
		Status:    row.Status,
		CreatedAt: row.Createdat,
		UpdatedAt: row.Updatedat,
	}

	return saveResult, nil
}

func (r *ResultRepository) FindByID(ctx context.Context, id string) (*domain.Result, error) {
	row, err := r.query.GetResult(ctx, id)
	if err != nil {
		return nil, mapSQLErrorResult(err)
	}

	return &domain.Result{
		ID:     row.ID,
		Status: row.Status,
	}, nil
}

func (r *ResultRepository) List(ctx context.Context, limit, page int64) ([]*domain.Result, error) {
	var rows []db.Result
	var err error

	if limit == 0 {
		rows, err = r.query.ListAllResults(ctx)
		if err != nil {
			return nil, mapSQLErrorResult(err)
		}
	} else {
		rows, err = r.query.ListResults(ctx, db.ListResultsParams{
			Limit:  limit,
			Offset: (page - 1) * limit,
		})
		if err != nil {
			return nil, mapSQLErrorResult(err)
		}
	}

	results := make([]*domain.Result, len(rows))
	for i, row := range rows {
		results[i] = &domain.Result{
			ID:     row.ID,
			Status: row.Status,
		}
	}

	return results, nil
}

func (r *ResultRepository) Update(ctx context.Context, result *domain.Result) error {
	err := r.query.UpdateResult(ctx, db.UpdateResultParams{
		ID:     result.ID,
		Status: result.Status,
	})
	if err != nil {
		return mapSQLErrorResult(err)
	}

	return nil
}

func (r *ResultRepository) Delete(ctx context.Context, id string) error {
	err := r.query.DeleteResult(ctx, id)
	if err != nil {
		return mapSQLErrorResult(err)
	}

	return nil
}
