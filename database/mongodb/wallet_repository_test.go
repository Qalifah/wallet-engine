package mongodb

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
	model "wallet-engine"
)

func TestWalletRepository_Add(t *testing.T) {
	ctx := context.Background()
	db, err := CreateTestDB(ctx)
	require.NoError(t, err)
	defer db.Client().Disconnect(ctx)

	walletRepo := NewWalletRepository(db)
	wallet := &model.Wallet{
		ID:        model.GenerateNewID(),
		UserID:    model.GenerateNewID(),
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	id, err := walletRepo.Add(ctx, wallet)
	require.NoError(t, err)
	require.Equal(t, id, wallet.ID.Hex())
}

func TestWalletRepository_GetByID(t *testing.T) {
	ctx := context.Background()
	db, err := CreateTestDB(ctx)
	require.NoError(t, err)
	defer db.Client().Disconnect(ctx)

	walletRepo := NewWalletRepository(db)
	wallet := &model.Wallet{
		ID:        model.GenerateNewID(),
		UserID:    model.GenerateNewID(),
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	id, err := walletRepo.Add(ctx, wallet)
	require.NoError(t, err)

	result, err := walletRepo.GetByID(ctx, id)
	require.NoError(t, err)
	require.Equal(t, wallet.UserID, result.UserID)
}

func TestWalletRepository_UpdateBalance(t *testing.T) {
	ctx := context.Background()
	db, err := CreateTestDB(ctx)
	require.NoError(t, err)
	defer db.Client().Disconnect(ctx)

	walletRepo := NewWalletRepository(db)
	wallet := &model.Wallet{
		ID:        model.GenerateNewID(),
		UserID:    model.GenerateNewID(),
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	id, err := walletRepo.Add(ctx, wallet)
	require.NoError(t, err)

	var newBalance float64 = 150
	require.NoError(t, walletRepo.UpdateBalance(ctx, id, newBalance))
	result, err := walletRepo.GetByID(ctx, id)
	require.NoError(t, err)
	require.Equal(t, newBalance, result.Balance)
}

func TestWalletRepository_UpdateIsActive(t *testing.T) {
	ctx := context.Background()
	db, err := CreateTestDB(ctx)
	require.NoError(t, err)
	defer db.Client().Disconnect(ctx)

	walletRepo := NewWalletRepository(db)
	wallet := &model.Wallet{
		ID:        model.GenerateNewID(),
		UserID:    model.GenerateNewID(),
		IsActive:  true,
		CreatedAt: time.Now(),
	}
	id, err := walletRepo.Add(ctx, wallet)
	require.NoError(t, err)

	var status bool
	require.NoError(t, walletRepo.UpdateIsActive(ctx, id, status))
	result, err := walletRepo.GetByID(ctx, id)
	require.NoError(t, err)
	require.Equal(t, status, result.IsActive)
}
