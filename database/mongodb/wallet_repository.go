package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	model "wallet-engine"
	"wallet-engine/errors"
)

type walletRepository struct {
	collection *mongo.Collection
}

func NewWalletRepository(db *mongo.Database) *walletRepository {
	c := db.Collection("wallets")
	return &walletRepository{collection: c}
}

func (r *walletRepository) Add(ctx context.Context, wallet *model.Wallet) (string, error) {
	result, err := r.collection.InsertOne(ctx, wallet)
	if err != nil {
		return "", err
	}
	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.ErrInvalidWalletID
	}
	return id.Hex(), nil
}

func (r *walletRepository) GetByID(ctx context.Context, id string) (*model.Wallet, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var wallet *model.Wallet
	err = r.collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&wallet)
	if err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *walletRepository) UpdateBalance(ctx context.Context, id string, newBalance float64) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"balance": newBalance}}
	if _, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update); err != nil {
		return err
	}
	return nil
}

func (r *walletRepository) UpdateIsActive(ctx context.Context, id string, status bool) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{"$set": bson.M{"is_active": status}}
	if _, err := r.collection.UpdateOne(ctx, bson.M{"_id": objectID}, update); err != nil {
		return err
	}
	return nil
}
