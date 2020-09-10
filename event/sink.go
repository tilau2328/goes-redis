package event

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/tilau2328/goes/core/command"
)

type Sink struct {
	client redis.Client
	ctx    context.Context
}

func (s *Sink) Dispatch(command command.ICommand) {

}
