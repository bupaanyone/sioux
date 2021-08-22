package service

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	"time"

	"github.com/gin-gonic/gin"

	"github.com/bupaanyone/sioux/server/config"
	"github.com/bupaanyone/sioux/server/model"
	"github.com/bupaanyone/sioux/server/utils"
)

const ctxKeyAccountId = "ctx_account_id"

func CtxSetAccountIdGin(c *gin.Context, id int64) error {
	var count int64
	if err := db.Model(&model.Account{}).Where("Id=?", id).Count(&count).Error; err != nil {

	} else if count == 0 {

	}
	c.Set(ctxKeyAccountId, id)
	return nil
}

func CtxGetAccountId(ctx context.Context) int64 {
	res, _ := ctx.Value(ctxKeyAccountId).(int64)
	return res
}

func hashPassword(password string) string {
	for i := 0; i < config.C.Service.PasswordIteration; i++ {
		raw := sha1.Sum([]byte(password + config.C.Service.PasswordSalt))
		password = hex.EncodeToString(raw[:])
	}
	return password
}

func AddAccount(ctx context.Context, account model.Account) error {
	if cur := CtxGetAccountId(ctx); cur != model.RootId {

	}

	account.Password = utils.StringToPtr(hashPassword(*account.Password))
	account.LastLogin = utils.TimeToPtr(time.Now())

	err := db.Create(&account)
	if err != nil {

	}

	return nil
}

func DeleteAccounts(ctx context.Context, ids []int64) error {

}

func Login(ctx context.Context, username, password string) error {

}

func UpdatePassword(ctx context.Context, id *int64, oldPassword *string, newPassword string) error {

}

func FindAccounts(ctx context.Context, ids []int64) ([]model.Account, int64, error) {

}
