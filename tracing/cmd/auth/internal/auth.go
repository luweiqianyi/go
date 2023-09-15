package internal

import (
	"fmt"
	"sync"
)

var gMemoryStore sync.Map

func GenToken(accountID string) string {
	value, ok := gMemoryStore.Load(accountID)
	if ok {
		return value.(string)
	}

	token := fmt.Sprintf("token-%v", accountID)
	gMemoryStore.Store(accountID, token)

	return token
}

func DeleteToken(accountID string) {
	gMemoryStore.Delete(accountID)
}

func TokenVerification(accountID string, token string) bool {
	value, ok := gMemoryStore.Load(accountID)
	if !ok {
		return false
	}

	storeToken := value.(string)
	if storeToken == token {
		return true
	}
	return false
}
