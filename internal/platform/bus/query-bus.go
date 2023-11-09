package bus

import (
	"api_go/kit/query"
	"context"
	"fmt"
)

type QueryBus struct {
	handlers map[query.Type]query.QueryHandler
}

func NewQueryBus() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.QueryHandler),
	}
}

func (b *QueryBus) Query(ctx context.Context, query query.Query) (query.Response, error) {
	handler, ok := b.handlers[query.Type()]

	if !ok {
		return nil, fmt.Errorf("no handler found for query type %s", query.Type())
	}

	response, err := handler.Handle(ctx, query)

	if err != nil {
		return nil, fmt.Errorf("error while handling %s - %s", query.Type(), err)
	}

	return response, nil
}

func (b *QueryBus) Register(queryType query.Type, handler query.QueryHandler) {
	b.handlers[queryType] = handler
}
