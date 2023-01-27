package usecase

import (
	"bytes"
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

type UserReq struct {
	Src         *string  `json:",omitempty"`
	Dest        *string  `json:",omitempty"`
	Sum         *float64 `json:",omitempty"`
	WalletLogin *string  `json:"wallet_login,omitempty"`
	Login       *string  `json:",omitempty"`
	Password    *string  `json:",omitempty"`
}

type webApiHandler struct {
	u Usecase
}

func NewWebApiHandler(u Usecase) *webApiHandler {
	return &webApiHandler{u}
}

type WebApiHandler interface {
	CheckWallet(*gin.Context)
	Charge(*gin.Context)
	GetWalletHistory(*gin.Context)
	GetWalletBalance(*gin.Context)
	Login(*gin.Context)
}

func (w *webApiHandler) CheckWallet(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	reqBody := buf.Bytes()
	req := UserReq{}

	if err := json.Unmarshal(reqBody, &req); err != nil {
		fmt.Println(string(reqBody))
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if _, exists := w.u.WalletExists(*req.WalletLogin); !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("wallet %s not found", *req.WalletLogin)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("wallet %s  exists", *req.WalletLogin)})
}

func (w *webApiHandler) Charge(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	reqBody := buf.Bytes()
	req := UserReq{}

	if err := json.Unmarshal(reqBody, &req); err != nil {
		fmt.Println(string(reqBody))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	tranID, err := w.u.Charge(*req.Src, *req.Dest, *req.Sum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment proccesed", "transactionID": tranID})
}

func (w *webApiHandler) GetWalletHistory(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	reqBody := buf.Bytes()
	req := UserReq{}
	if err := json.Unmarshal(reqBody, &req); err != nil {
		fmt.Println(string(reqBody))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	walletHistory, err := w.u.WalletHistory(*req.WalletLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": walletHistory})
}

func (w *webApiHandler) GetWalletBalance(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	reqBody := buf.Bytes()
	req := UserReq{}

	if err := json.Unmarshal(reqBody, &req); err != nil {
		fmt.Println(string(reqBody))
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	wallet, exists := w.u.WalletExists(*req.WalletLogin)

	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("wallet %s not found", *req.WalletLogin)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("wallet %s  exists", *req.WalletLogin), "balance": wallet.Sum})
}

func (w *webApiHandler) Login(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	reqBody := buf.Bytes()
	req := UserReq{}

	if err := json.Unmarshal(reqBody, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := w.u.GetLogin(*req.Login)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": err.Error(),
		})
		return
	}

	hasher := sha256.New()
	userPass := fmt.Sprintf("%s:%s", *req.Password, user.Salt)
	hasher.Write([]byte(userPass))
	encodedPass := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	if encodedPass != user.Password {
		c.JSON(http.StatusForbidden, gin.H{"message": "invalid credentianls"})
		return
	}
	w.u.UserValidate(ctx, user.ID)
	c.JSON(http.StatusOK, gin.H{"messsage": "fucking good"})
}

func (w *webApiHandler) HeaderCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		// middleware
		var bodyBytes []byte
		if c.Request.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		digest := c.GetHeader("X-Digest")
		userID := c.GetHeader("X-UserId")
		b := sha1.New()
		b.Write(bodyBytes)
		sha := base64.URLEncoding.EncodeToString(b.Sum(nil))
		if sha != digest {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		id, _ := strconv.Atoi(userID)
		if valid := w.u.CheckUserByID(ctx, id); !valid {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
