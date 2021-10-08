package auth

import (
	ahelper "htblog/internal/api/helper"
	mhelper "htblog/internal/models/helper"
	"htblog/internal/models/usermod"
	"htblog/internal/terror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResp struct {
	Token string `json:"token"`
}

func Login(ctx *gin.Context) {
	logger := ahelper.GetLog(ctx)

	var req LoginReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, terror.ErrReqParam)
		return
	}

	user, err := usermod.GetUser(req.Account, req.Password)
	if err != nil {
		logger.Error("usermod.GetUser err : ", err)
		ctx.JSON(http.StatusOK, terror.ErrUnexpected)
		return
	}
	if user == nil {
		ctx.JSON(http.StatusOK, terror.ErrDataNotFound)
		return
	}

	// create session
	DelTokenByAccount(user.Account) // 防止多点登录
	token := GenToken()
	ok := SetTokenByAccount(user.Account, token)
	if !ok {
		logger.Error("SetTokenByAccount ok : ", ok)
		ctx.JSON(http.StatusOK, terror.ErrUnexpected)
		return
	}
	ok = SetUserByToken(token, user)
	if !ok {
		logger.Error("SetTokenByAccount ok : ", ok)
		ctx.JSON(http.StatusOK, terror.ErrUnexpected)
		return
	}

	logger.Infof("user %s login succeed token : %s", user.NickName, token)
	ctx.JSON(http.StatusOK, terror.Msg{
		MsgBase: terror.Succeed,
		Data: LoginResp{
			Token: token,
		},
	})
}

type RegisterReq struct {
	Account  string `json:"account" binding:"required"`
	Password string `json:"password" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
}

func Register(ctx *gin.Context) {
	logger := ahelper.GetLog(ctx)

	var req RegisterReq
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusOK, terror.ErrReqParam)
		return
	}

	err = usermod.CreateUser(req.Account, req.Password, req.NickName)
	if mhelper.IsUniqueConstraintError(err) {
		logger.Info("usermod.CreateUser UniqueConstraint : ", err)
		ctx.JSON(http.StatusOK, terror.ErrDuplicateUser)
		return
	}
	if err != nil {
		logger.Error("usermod.CreateUser err : ", err)
		ctx.JSON(http.StatusOK, terror.ErrUnexpected)
		return
	}

	logger.Infof("register user success [%+v]", req)
	ctx.JSON(http.StatusOK, terror.Succeed)
}

func Logout(ctx *gin.Context) {
	logger := ahelper.GetLog(ctx)

	token := ctx.GetHeader("x-token")
	if token == "" {
		ctx.JSON(http.StatusOK, terror.ErrReqParam)
		return
	}

	user := GetUserByToken(token)
	if user == nil {
		logger.Warn("invalid token")
		ctx.JSON(http.StatusOK, terror.Succeed)
		return
	}

	DelTokenByAccount(user.Account)
	DelUserByToken(token)

	logger.Infof("user %s logout", user.NickName)
	ctx.JSON(http.StatusOK, terror.Succeed)
}
