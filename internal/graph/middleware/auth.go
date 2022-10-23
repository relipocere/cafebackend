package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/relipocere/cafebackend/internal/database"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
)

type userGetter interface {
	GetBySessionID(ctx context.Context, q database.Queryable, sessionID string) (*model.User, error)
}

func AuthenticationMW(q database.Queryable, userGetter userGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("Authorization")

		user, err := userGetter.GetBySessionID(c, q, sessionID)
		if err != nil {
			resp := mapToGQLErrorResponse("Internal error", graphmodel.ErrorCodeInternal)
			c.JSON(http.StatusInternalServerError, resp)
			return
		}

		if user != nil {
			ctx := context.WithValue(c.Request.Context(), "user", user)
			c.Request = c.Request.WithContext(ctx)
		}

		c.Next()
	}
}

func mapToGQLErrorResponse(message string, code graphmodel.ErrorCode) map[string]interface{} {
	return map[string]interface{}{
		"message": message,
		"extensions": map[string]interface{}{
			"code": code,
		},
	}
}
