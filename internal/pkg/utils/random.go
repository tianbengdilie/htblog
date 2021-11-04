package utils

import (
	"encoding/base64"

	"github.com/google/uuid"
)

//https://github.com/taskcluster/slugid-go/blob/master/slugid/slugid.go google compress uuid
func GetUUID() string {
	uid := []byte(uuid.New().String())
	// unset most significant bit of first byte to ensure slug starts with [A-Za-f]
	uid[0] &= 0x7f
	return base64.URLEncoding.EncodeToString(uid)[:22] // Drop '==' padding
}
