package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	ChallengeDifficulty = 4
	Challenge           = "guytyu24jklj"
)

func POWMiddleware(c *gin.Context) {
	hashCash := c.GetHeader("HashCash")
	if hashCash == "" {
		c.Header("Challenge", Challenge)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Proof of Work challenge required"})
		return
	}

	valid := validateClientHashCash(hashCash, ChallengeDifficulty)
	if !valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Proof of Work"})
		return
	}

	c.Next()
}

func validateClientHashCash(hashCash string, difficulty int) bool {
	result := strings.Split(hashCash, ":")
	if result[0] != Challenge {
		return false
	}
	hash := sha256.Sum256([]byte(result[1]))
	stringHash := hex.EncodeToString(hash[:])

	return strings.HasPrefix(stringHash, strings.Repeat("0", difficulty))
}
