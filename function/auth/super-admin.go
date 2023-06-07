package auth

import (
	"BackendCRM/utility/token"
	"strconv"
	"strings"
)

func CheckSuperAdmin(tokenParam TokenParam) bool {
	if tokenParam.Data == "" {
		return false
	}
	data := strings.TrimPrefix(tokenParam.Data, "Bearer ")
	resJWT, err := token.VerfiyJWT(data, "koentji")
	if err != nil {
		return false
	}
	roleId, err := strconv.Atoi(resJWT)
	if err != nil {
		return false
	}
	if roleId != 1 {
		return false
	}
	return true
}
