package repositories

import (
	"ai_project/internal/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LendersRepository interface {
	UpdateByName(ctx context.Context, lenderName string, updateData models.Lender) (*mongo.UpdateResult, error)
}

type lendersRepo struct {
	col *mongo.Collection
}

func NewLenderRepo(db *mongo.Database) LendersRepository {
	return &lendersRepo{col: db.Collection("lenders")}
}

func (r *lendersRepo) UpdateByName(ctx context.Context, lenderName string, updateData models.Lender) (*mongo.UpdateResult, error) {
	// Implement the logic to update a document by name in the MongoDB collection
	update := bson.M{
		"$set": updateData,
	}
	opts := options.Update().SetUpsert(false) // avoid inserting new document if not found
	filter := bson.M{
		"lender_name": bson.M{
			"$regex":   lenderName,
			"$options": "i", // Case-insensitive option
		},
	}
	result, err := r.col.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return nil, err
	}
	return result, nil
}
