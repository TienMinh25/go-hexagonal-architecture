package http

import (
	"strconv"

	domainauth "github.com/TienMinh25/go-hexagonal-architecture/internal/application/domain/auth"
	modelv1 "github.com/TienMinh25/go-hexagonal-architecture/pkg/model/v1"
	"github.com/gin-gonic/gin"
)

// stringToUint64 is a helper function to convert a string to uint64
func stringToUint64(str string) (uint64, error) {
	num, err := strconv.ParseUint(str, 10, 64)

	return num, err
}

// getAuthPayload is a helper function to get the auth payload from the context
func getAuthPayload(ctx *gin.Context, key string) *domainauth.TokenPayload {
	return ctx.MustGet(key).(*domainauth.TokenPayload)
}

// toMap is a helper function to add meta and data to a map
func toMap(m modelv1.Meta, data any, key string) map[string]any {
	return map[string]any{
		"meta": m,
		key:    data,
	}
}
