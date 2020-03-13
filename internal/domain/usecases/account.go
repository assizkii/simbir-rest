package usecases

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"time"
)

func GetRandNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}

func HashPassword(password string) string {
	hash := md5.Sum([]byte(password + "secret"))
	return hex.EncodeToString(hash[:])
}

func CheckPassword(password string, currentPassword string) bool {
	return HashPassword(password) == currentPassword
}
