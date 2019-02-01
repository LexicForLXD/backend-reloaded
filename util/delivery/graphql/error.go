package graphql

import (
	"context"

	"github.com/lexicforlxd/backend-reloaded/lexicError"
	"github.com/pkg/errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/gqlerror"
)

func CustomErrorHandler(ctx context.Context, e error) *gqlerror.Error {
	if lexicError, ok := e.(*lexicError.LexicError); ok {

		extensions := make(map[string]interface{})

		extensions["code"] = lexicError.AppCode
		extensions["error"] = errors.Cause(lexicError).Error()

		return &gqlerror.Error{
			Message:    lexicError.StatusText,
			Path:       graphql.GetResolverContext(ctx).Path(),
			Extensions: extensions,
		}

	}

	return graphql.DefaultErrorPresenter(ctx, e)
}
