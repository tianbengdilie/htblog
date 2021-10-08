package utils

import "math/rand"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func GetRandomStr(n uint) string {
	ret := make([]rune, n)
	for i := range ret {
		ret[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(ret)
}
