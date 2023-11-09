package query

import "context"

type QueryBus interface {
	Query(context.Context, Query) (Response, error)
	Register(Type, QueryHandler)
}

type Type string

type Query interface {
	Type() Type
}

type QueryHandler interface {
	Handle(context.Context, Query) (Response, error)
}

type Response interface {
}
