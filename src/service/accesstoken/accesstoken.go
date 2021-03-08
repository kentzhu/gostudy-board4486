package accesstoken

import (
	"math/rand"
	"time"
)

var tokens = make(map[string]interface{})

func init() {
	rand.Seed(time.Now().UnixNano())
}

// IsAccessTokenValid 检查Token是否有效
func IsAccessTokenValid(token string) bool {
	if _, ok := tokens[token]; ok {
		return true
	} else {
		return false
	}
}

// Gen 生成Token
func Gen(bindingData interface{}) string {
	token := randStringRunes(16)
	save(token, bindingData)
	return token
}

// Destroy 销毁Token
func Destroy(token string) {
	remove(token)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func save(token string, bindingContent interface{}) {
	tokens[token] = bindingContent
}

func remove(token string) {
	delete(tokens, token)
}
