package db

import (
	"context"
	"database/sql"

	"github.com/ynoacamino/alert-go/gen/db"
	"github.com/ynoacamino/alert-go/internal/core/domain"
	"github.com/ynoacamino/alert-go/internal/core/ports"
)

type MailAddressRepository struct {
	queries *db.Queries
}

func NewMailAddressRepository(sqlDB *sql.DB) ports.MailAddressRepository {
	return &MailAddressRepository{
		queries: db.New(sqlDB),
	}
}

func (r *MailAddressRepository) Save(ctx context.Context, mail *domain.MailAddress) (*domain.MailAddress, error) {
	row, err := r.queries.CreateMailAddress(ctx, db.CreateMailAddressParams{
		ID:      mail.ID,
		Address: mail.Address,
	})
	if err != nil {
		return nil, mapSQLErrorMailAddress(err)
	}

	saveMailAddress := &domain.MailAddress{
		ID:        row.ID,
		Address:   row.Address,
		Active:    row.Active,
		CreatedAt: row.Createdat,
		UpdatedAt: row.Updatedat,
	}

	return saveMailAddress, nil
}

func (r *MailAddressRepository) FindByID(ctx context.Context, id string) (*domain.MailAddress, error) {
	row, err := r.queries.GetMailAddress(ctx, id)
	if err != nil {
		return nil, mapSQLErrorMailAddress(err)
	}

	return &domain.MailAddress{
		ID:        row.ID,
		Address:   row.Address,
		Active:    row.Active,
		CreatedAt: row.Createdat,
		UpdatedAt: row.Updatedat,
	}, nil
}

func (r *MailAddressRepository) List(ctx context.Context, limit, page int64) ([]*domain.MailAddress, error) {
	rows, err := r.queries.ListMailAddresses(ctx, db.ListMailAddressesParams{
		Limit:  limit,
		Offset: (page - 1) * limit,
	})
	if err != nil {
		return nil, mapSQLErrorMailAddress(err)
	}

	mailAddresses := make([]*domain.MailAddress, len(rows))
	for i, row := range rows {
		mailAddresses[i] = &domain.MailAddress{
			ID:        row.ID,
			Address:   row.Address,
			Active:    row.Active,
			CreatedAt: row.Createdat,
			UpdatedAt: row.Updatedat,
		}
	}

	return mailAddresses, nil
}

func (r *MailAddressRepository) Delete(ctx context.Context, id string) error {
	err := r.queries.DeleteMailAddress(ctx, id)
	if err != nil {
		return mapSQLErrorMailAddress(err)
	}

	return nil
}

func (r *MailAddressRepository) Update(ctx context.Context, mail *domain.MailAddress) error {
	err := r.queries.UpdateMailAddress(ctx, db.UpdateMailAddressParams{
		ID:      mail.ID,
		Address: mail.Address,
		Active:  mail.Active,
	})
	if err != nil {
		return mapSQLErrorMailAddress(err)
	}

	return nil
}
