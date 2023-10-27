package natsService

import (
	"WB_Intern_L0/internal/usecase"
	"WB_Intern_L0/pkg/cache"
	"github.com/nats-io/nats.go"
)

type Nats struct {
	NATSServer *nats.Conn
	Services   *usecase.UseCase
	Cache      *cache.CacheService
}

func NewNats(services *usecase.UseCase, NATSServer *nats.Conn, cache *cache.CacheService) *Nats {
	return &Nats{NATSServer: NATSServer, Services: services, Cache: cache}
}
