package sqlite

import (
	"context"

	"github.com/brittonhayes/staffing"
	"github.com/uptrace/bun"
)

type recommendationRepository struct {
	db *bun.DB
}

func NewRecommendationRepository(connection string, inmem bool) staffing.RecommendationRepository {
	// sqldb, err := sql.Open(sqliteshim.ShimName, connection)
	// if err != nil {
	// 	panic(err)
	// }
	// if inmem {
	// 	sqldb.SetMaxOpenConns(1)
	// 	sqldb.SetMaxIdleConns(1000)
	// 	sqldb.SetConnMaxLifetime(0)
	// }
	//
	// db := bun.NewDB(sqldb, sqlitedialect.New())
	//
	// _, err = db.NewCreateTable().Model((*staffing.Recommendation)(nil)).Exec(context.Background())
	return &recommendationRepository{
		db: nil,
	}
}
func (r *recommendationRepository) Close() error {
	return r.db.Close()
}

func (r *recommendationRepository) CreateUser(ctx context.Context, name string) error {
	return nil
}
