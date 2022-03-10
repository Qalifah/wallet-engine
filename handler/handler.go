package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"wallet-engine/errors"

	log "github.com/sirupsen/logrus"
	model "wallet-engine"
)

type Handler struct {
	walletRepo model.WalletRepository
}

func New(walletRepo model.WalletRepository) *Handler {
	return &Handler{walletRepo: walletRepo}
}

func (h *Handler) CreateNewWallet(c *gin.Context) {
	req := struct {
		UserID string `json:"user_id"`
	}{}

	err := c.ShouldBind(&req)
	if err != nil {
		log.WithError(err).Error(errors.ErrInvalidRequestBody)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrInvalidRequestBody))
		return
	}

	userID, err := model.GetIdFromStr(req.UserID)
	if err != nil {
		log.WithError(err).Error(errors.ErrInvalidUserID)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrInvalidUserID))
		return
	}

	walletID, err := h.walletRepo.Add(c, model.New(userID))
	if err != nil {
		log.WithError(err).Error(errors.ErrUnExpectedError)
		c.JSON(http.StatusInternalServerError, NewErrorResponse(errors.ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusCreated, NewResponse(StatusSuccess, "wallet successfully created!", gin.H{"wallet_id": walletID}))
}

func (h *Handler) DebitWallet(c *gin.Context) {
	req := struct {
		WalletID string  `json:"wallet_id"`
		Amount   float64 `json:"amount"`
	}{}

	err := c.ShouldBind(&req)
	if err != nil {
		log.WithError(err).Error(errors.ErrInvalidRequestBody)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrInvalidRequestBody))
		return
	}

	wallet, err := h.walletRepo.GetByID(c, req.WalletID)
	if err != nil {
		log.WithError(err).Error(errors.ErrWalletNotFound)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrWalletNotFound))
	}

	newBalance := wallet.Balance - req.Amount
	err = h.walletRepo.UpdateBalance(c, req.WalletID, newBalance)
	if err != nil {
		log.WithError(err).Error(errors.ErrUnExpectedError)
		c.JSON(http.StatusInternalServerError, NewErrorResponse(errors.ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusOK, NewResponse(StatusSuccess, "wallet successfully debited!", nil))
}

func (h *Handler) CreditWallet(c *gin.Context) {
	req := struct {
		WalletID string  `json:"wallet_id"`
		Amount   float64 `json:"amount"`
	}{}

	err := c.ShouldBind(&req)
	if err != nil {
		log.WithError(err).Error(errors.ErrInvalidRequestBody)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrInvalidRequestBody))
		return
	}

	wallet, err := h.walletRepo.GetByID(c, req.WalletID)
	if err != nil {
		log.WithError(err).Error(errors.ErrWalletNotFound)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrWalletNotFound))
	}

	newBalance := wallet.Balance + req.Amount
	err = h.walletRepo.UpdateBalance(c, req.WalletID, newBalance)
	if err != nil {
		log.WithError(err).Error(errors.ErrUnExpectedError)
		c.JSON(http.StatusInternalServerError, NewErrorResponse(errors.ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusOK, NewResponse(StatusSuccess, "wallet successfully credited!", nil))
}

func (h *Handler) ActivateWallet(c *gin.Context) {
	req := struct {
		WalletID string `json:"wallet_id"`
	}{}

	err := c.ShouldBind(&req)
	if err != nil {
		log.WithError(err).Error(errors.ErrInvalidRequestBody)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrInvalidRequestBody))
		return
	}

	_, err = h.walletRepo.GetByID(c, req.WalletID)
	if err != nil {
		log.WithError(err).Error(errors.ErrWalletNotFound)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrWalletNotFound))
	}

	err = h.walletRepo.UpdateIsActive(c, req.WalletID, true)
	if err != nil {
		log.WithError(err).Error(errors.ErrUnExpectedError)
		c.JSON(http.StatusInternalServerError, NewErrorResponse(errors.ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusOK, NewResponse(StatusSuccess, "wallet successfully activated!", nil))
}

func (h *Handler) DeActivateWallet(c *gin.Context) {
	req := struct {
		WalletID string `json:"wallet_id"`
	}{}

	err := c.ShouldBind(&req)
	if err != nil {
		log.WithError(err).Error(errors.ErrInvalidRequestBody)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrInvalidRequestBody))
		return
	}

	_, err = h.walletRepo.GetByID(c, req.WalletID)
	if err != nil {
		log.WithError(err).Error(errors.ErrWalletNotFound)
		c.JSON(http.StatusBadRequest, NewErrorResponse(errors.ErrWalletNotFound))
	}

	err = h.walletRepo.UpdateIsActive(c, req.WalletID, false)
	if err != nil {
		log.WithError(err).Error(errors.ErrUnExpectedError)
		c.JSON(http.StatusInternalServerError, NewErrorResponse(errors.ErrUnExpectedError))
		return
	}

	c.JSON(http.StatusOK, NewResponse(StatusSuccess, "wallet successfully deactivated!", nil))
}
