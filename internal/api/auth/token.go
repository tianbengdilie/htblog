package auth

import (
	"htblog/internal/config"
	"htblog/internal/models/usermod"
	"htblog/internal/pkg/utils"
)

const (
	prefixAccount2Token = "account-2-token-"
	prefixToken2User    = "token-2-user-"

	headerToken = "token"
)

func GenToken() string {
	return utils.GetRandomStr(32)
}

func SetTokenByAccount(account, token string) bool {
	return config.GetUserCache().Set(prefixAccount2Token+account, token)
}

func GetTokenByAccount(account string) string {
	return config.GetUserCache().Get(prefixAccount2Token + account).(string)
}

func DelTokenByAccount(account string) {
	config.GetUserCache().Del(prefixAccount2Token + account)
}

func SetUserByToken(token string, user *usermod.UserInfo) bool {
	return config.GetUserCache().Set(prefixToken2User+token, user)
}

func GetUserByToken(token string) *usermod.UserInfo {
	return config.GetUserCache().Get(prefixToken2User + token).(*usermod.UserInfo)
}

func DelUserByToken(token string) {
	config.GetUserCache().Del(prefixToken2User + token)
}
