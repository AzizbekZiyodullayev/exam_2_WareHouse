package postgres

import (
	"context"
	"fmt"
	"warehouse/config"
	"warehouse/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db             *pgxpool.Pool
	product        *productRepo
	category       *categoryRepo
	branch         *branchRepo
	coming         *comingRepo
	remainder      *remainderRepo
	comingProducts *comingProductsRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {
	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}

	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}

func (s *store) Product() storage.ProductI {

	if s.product == nil {
		s.product = NewProductRepo(s.db)
	}

	return s.product
}

func (s *store) Category() storage.CategoryI {

	if s.category == nil {
		s.category = NewCategoryRepo(s.db)
	}

	return s.category
}

func (s *store) Branch() storage.BranchI {

	if s.branch == nil {
		s.branch = NewBranchRepo(s.db)
	}

	return s.branch
}

func (s *store) Coming() storage.ComingI {

	if s.coming == nil {
		s.coming = NewComingRepo(s.db)
	}

	return s.coming
}

func (s *store) Remainder() storage.RemainderI {

	if s.remainder == nil {
		s.remainder = NewRemainderRepo(s.db)
	}

	return s.remainder
}

func (s *store) ComingProducts() storage.ComingProductsI {

	if s.comingProducts == nil {
		s.comingProducts = NewComingProductsRepo(s.db)
	}

	return s.comingProducts
}
