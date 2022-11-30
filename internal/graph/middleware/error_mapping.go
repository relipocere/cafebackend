package middleware

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	graphmodel "github.com/relipocere/cafebackend/internal/graph/graph-model"
	"github.com/relipocere/cafebackend/internal/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

type logFn func(format string, args ...interface{})

// ErrorHandlerMw логгирует и прообразует ошибки в graphql-ные.
func ErrorHandlerMw() graphql.ErrorPresenterFunc {
	return func(ctx context.Context, err error) *gqlerror.Error {

		// gql в генерированном коде оборачивает ошибку
		gqlErr, ok := err.(*gqlerror.Error)
		if ok {
			err = gqlErr.Unwrap()
		}

		logFunc, mappedErr := mapError(ctx,  err)

		path := graphql.GetPath(ctx)
		logFunc("'%s': %v", path.String(), err)

		return mappedErr
	}
}

func mapError(ctx context.Context,  err error) (logFn, *gqlerror.Error) {
	if businessError := new(model.Error); errors.As(err, businessError) {
		var code graphmodel.ErrorCode
		switch businessError.Code {
		case model.ErrorCodeUnauthenticated:
			code = graphmodel.ErrorCodeUnauthenticated
		case model.ErrorCodeUnauthorized:
			code = graphmodel.ErrorCodeUnauthorized
		case model.ErrorCodeBadRequest:
			code = graphmodel.ErrorCodeBadRequest
		case model.ErrorCodeNotFound:
			code = graphmodel.ErrorCodeNotFound
		case model.ErrorCodeFailedPrecondition:
			code = graphmodel.ErrorCodeFailedPrecondition
		}

		return zap.S().Infof, graphqlError(ctx, businessError.Error(), code)
	}

	return zap.S().Errorf, graphqlError(ctx, "Internal server error", graphmodel.ErrorCodeInternal)
}

func graphqlError(ctx context.Context, message string, code graphmodel.ErrorCode) *gqlerror.Error {
	return &gqlerror.Error{
		Message: message,
		Extensions: map[string]interface{}{
			"code": code,
		},
		Path: graphql.GetPath(ctx),
	}
}
