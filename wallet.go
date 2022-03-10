package wallet_engine

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Wallet struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	UserID    primitive.ObjectID `json:"user_id" bson:"user_id"`
	Balance   float64            `json:"balance" bson:"balance"`
	IsActive  bool               `json:"is_active" bson:"is_active"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

func New(userID primitive.ObjectID) *Wallet {
	return &Wallet{
		ID:        GenerateNewID(),
		UserID:    userID,
		Balance:   0,
		IsActive:  true,
		CreatedAt: time.Now(),
	}
}

func GenerateNewID() primitive.ObjectID {
	return primitive.NewObjectID()
}

func GetIdFromStr(id string) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id)
}

type WalletRepository interface {
	Add(ctx context.Context, w *Wallet) (string, error)
	GetByID(ctx context.Context, id string) (*Wallet, error)
	UpdateBalance(ctx context.Context, id string, newBalance float64) error
	UpdateIsActive(ctx context.Context, id string, status bool) error
}
