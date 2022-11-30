package middleware

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/relipocere/cafebackend/internal/auth"
	"github.com/relipocere/cafebackend/internal/database"
	userdb "github.com/relipocere/cafebackend/internal/database/user"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
	"go.uber.org/zap"
)

type userGetter interface {
	Get(ctx context.Context, q database.Queryable, filter userdb.GetFilter) (*model.User, error)
}

func AuthenticationMW(q database.Queryable, userGetter userGetter) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader(auth.Header)

		user, err := userGetter.Get(c, q, userdb.GetBySessionID(sessionID))
		if err != nil {
			resp := mapToGQLErrorResponse("Internal error", graphmodel.ErrorCodeInternal)
			c.AbortWithStatusJSON(http.StatusInternalServerError, resp)
			return
		}

		if user != nil {
			ctx := context.WithValue(c.Request.Context(), auth.User, *user)
			c.Request = c.Request.WithContext(ctx)
		}
		zap.S().Debugw("auth mw", "headerLen", len(sessionID), "retrievedUser", user != nil)
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
