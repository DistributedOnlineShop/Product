package gapi

import (
	db "Product/db/sqlc"
	"Product/token"
	"Product/util"
	"fmt"

	"github.com/go-redis/redis"
)

type Server struct {
	config util.Config
	token  token.Maker
	store  db.Store
	redis  *redis.Client
}

func ServerSetup(config util.Config, store db.Store, redis *redis.Client) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.KeySeed)
	if err != nil {
		return nil, fmt.Errorf("fail to create key pair: %v", err)
	}

	return &Server{
		config: config,
		store:  store,
		token:  tokenMaker,
		redis:  redis,
	}, nil
}
