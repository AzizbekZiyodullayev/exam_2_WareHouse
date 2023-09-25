package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"warehouse/api/models"
	"warehouse/pkg/helper"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type branchRepo struct {
	db *pgxpool.Pool
}

func NewBranchRepo(db *pgxpool.Pool) *branchRepo {
	return &branchRepo{
		db: db,
	}
}

func (r *branchRepo) Create(ctx context.Context, req *models.CreateBranch) (string, error) {

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return "", nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO branch(id, name, address, phone_number, updated_at)
		VALUES ($1, $2, $3, $4, NOW())
	`

	_, err = trx.Exec(ctx, query,
		id,
		req.Name,
		req.Address,
		req.PhoneNumber,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (r *branchRepo) GetByID(ctx context.Context, req *models.BranchPrimaryKey) (*models.Branch, error) {

	var (
		query string

		id          sql.NullString
		name        sql.NullString
		address     sql.NullString
		phoneNumber sql.NullString
		createdAt   sql.NullString
		updatedAt   sql.NullString
	)

	query = `
		SELECT
			id,
			name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM branch
		WHERE id=$1
	`

	err := r.db.QueryRow(ctx, query, req.Id).Scan(
		&id,
		&name,
		&address,
		&phoneNumber,
		&createdAt,
		&updatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &models.Branch{
		Id:          id.String,
		Name:        name.String,
		Address:     address.String,
		PhoneNumber: phoneNumber.String,
		CreatedAt:   createdAt.String,
		UpdatedAt:   updatedAt.String,
	}, nil
}

func (r *branchRepo) GetList(ctx context.Context, req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {

	var (
		resp   = &models.BranchGetListResponse{}
		query  string
		where  = " WHERE TRUE"
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id,
			name,
			address,
			phone_number,
			created_at,
			updated_at
		FROM branch
	`

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Search != "" {
		where += ` AND name ILIKE '%' || '` + req.Search + `' || '%'`
	}

	query += where + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var (
			id          sql.NullString
			name        sql.NullString
			address     sql.NullString
			phoneNumber sql.NullString
			createdAt   sql.NullString
			updatedAt   sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&address,
			&phoneNumber,
			&createdAt,
			&updatedAt,
		)

		if err != nil {
			return nil, err
		}

		resp.Branches = append(resp.Branches, &models.Branch{
			Id:          id.String,
			Name:        name.String,
			Address:     address.String,
			PhoneNumber: phoneNumber.String,
			CreatedAt:   createdAt.String,
			UpdatedAt:   updatedAt.String,
		})
	}

	return resp, nil
}

func (r *branchRepo) Update(ctx context.Context, req *models.UpdateBranch) (int64, error) {
	trx, err := r.db.Begin(ctx)
	if err != nil {
		return 0, nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	var (
		query  string
		params map[string]interface{}
	)

	query = `
		UPDATE
			branch
		SET
			name = :name,
			address = :address,
			phone_number = :phone_number,
			updated_at = NOW()
		WHERE id = :id
	`

	params = map[string]interface{}{
		"id":           req.Id,
		"name":         req.Name,
		"address":      req.Address,
		"phone_number": req.PhoneNumber,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	result, err := trx.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected(), nil
}

func (r *branchRepo) Delete(ctx context.Context, req *models.BranchPrimaryKey) error {

	trx, err := r.db.Begin(ctx)
	if err != nil {
		return nil
	}

	defer func() {
		if err != nil {
			trx.Rollback(ctx)
		} else {
			trx.Commit(ctx)
		}
	}()

	_, err = trx.Exec(ctx, "DELETE FROM branch WHERE id = $1", req.Id)
	if err != nil {
		return err
	}

	return nil
}
