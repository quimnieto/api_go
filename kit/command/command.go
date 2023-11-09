package command

import "context"

type CommandBus interface {
	Dispatch(context.Context, Command) error
	Register(Type, CommandHandler)
}

type Type string

type Command interface {
	Type() Type
}

type CommandHandler interface {
	Handle(context.Context, Command) error
}
