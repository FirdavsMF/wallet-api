package usecase

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/FirdavsMF/wallet-api/internal/entity"
	"github.com/FirdavsMF/wallet-api/internal/service"
	"github.com/FirdavsMF/wallet-api/pkg/logger"
)

type usecase struct {
	db          *sqlx.DB
	userRepo    service.UserSer
	walletRepo  service.WalletSer
	clientRepo  service.ClientSer
	paymentRepo service.PaymentSer
	auth        service.Authentificator
	logger      logger.Interface
}

type Usecase interface {
	Charge(src string, dest string, sum float64) (string, error)
	WalletExists(login string) (entity.Wallet, bool)
	WalletHistory(login string) (service.WalletHistory, error)
	GetLogin(login string) (entity.User, error)
	UserValidate(ctx context.Context, id int) error
	CheckUserByID(ctx context.Context, id int) bool
}

func NewUsecase(db *sqlx.DB, l logger.Interface, r *redis.Client) *usecase {
	userRepo := service.NewUserSer(db)
	walletRepo := service.NewWalletSer(db)
	clientRepo := service.NewClientSer(db)
	paymentRepo := service.NewPaymentSer(db)
	authRepo := service.NewAuthentificator(r)
	return &usecase{
		db:          db,
		userRepo:    userRepo,
		walletRepo:  walletRepo,
		clientRepo:  clientRepo,
		paymentRepo: paymentRepo,
		auth:        authRepo,
		logger:      l,
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func (u *usecase) Charge(src string, dest string, sum float64) (string, error) {
	tx := u.db.MustBegin()
	extID := RandStringRunes(20)

	srcWallet, err := u.walletRepo.GetByLogin(src, tx)

	if err != nil {
		u.logger.Error(err)
		return "", err
	}

	if srcWallet.Sum < sum {
		errStr := fmt.Sprintf("Insuffient funds in wallet %s", srcWallet.Login)
		u.logger.Error(errStr)
		return "", errors.New(errStr)
	}

	destWallet, err := u.walletRepo.GetByLogin(dest, tx)

	if err != nil {
		u.logger.Error(err)
		return "", err
	}

	client, err := u.clientRepo.GetByWalletID(destWallet.ID, tx)
	if err != nil {
		u.logger.Error(err)
		return "", err
	}

	if !client.Identified && destWallet.Sum+sum > 10000 {
		errStr := fmt.Sprintf("too big sum for unidentificated destination %s", destWallet.Login)
		return "", errors.New(errStr)
	}

	if client.Identified && destWallet.Sum+sum > 100000 {
		errStr := fmt.Sprintf("too big sum for unidentificated destination %s", destWallet.Login)
		return "", errors.New(errStr)
	}

	if err := u.walletRepo.ChargeWallet(tx, srcWallet.Login, -sum); err != nil {
		return "", err
	}
	if err := u.walletRepo.ChargeWallet(tx, destWallet.Login, sum); err != nil {
		return "", err
	}

	if err := u.paymentRepo.CreatePayment(tx, src, dest, sum, extID); err != nil {
		return "", err
	}
	err = tx.Commit()
	return extID, err
}

func (u *usecase) WalletExists(login string) (entity.Wallet, bool) {

	wallet, err := u.walletRepo.GetByLogin(login, nil)
	if err != nil {
		return wallet, false
	}

	return wallet, true
}

func (u *usecase) WalletHistory(login string) (service.WalletHistory, error) {
	return u.walletRepo.GetWalletHistory(login)
}

func (u *usecase) GetLogin(login string) (entity.User, error) {
	return u.userRepo.GetByLogin(login)
}

func (u *usecase) UserValidate(ctx context.Context, id int) error {
	return u.auth.UserValidate(ctx, id)
}

func (u *usecase) CheckUserByID(ctx context.Context, id int) bool {
	return u.auth.CheckUserByID(ctx, id)
}
