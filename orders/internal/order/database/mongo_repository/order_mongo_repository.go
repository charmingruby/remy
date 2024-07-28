package mongo_repository

import "context"

func NewOrderMongoRepository() *OrderMongoRepository {
	return &OrderMongoRepository{}
}

type OrderMongoRepository struct{}

func (r *OrderMongoRepository) Create(context.Context) error {
	return nil
}
