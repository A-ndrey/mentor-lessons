package main

import (
	"context"
	"github.com/go-faster/errors"
	"log"
	"net/http"
	"openapi_example/api"
)

type service struct {
	storage map[int]api.User
}

func (s *service) UserGet(ctx context.Context) ([]api.User, error) {
	var result []api.User
	for _, u := range s.storage {
		result = append(result, u)
	}

	return result, nil
}

func (s *service) UserIDGet(ctx context.Context, params api.UserIDGetParams) (*api.User, error) {
	u, ok := s.storage[params.ID]
	if !ok {
		return nil, errors.New("user not found")
	}

	return &u, nil
}

func main() {

	s := service{storage: map[int]api.User{
		1: {
			ID:   api.NewOptInt64(1),
			Name: api.NewOptString("User1"),
		},
		2: {
			ID:   api.NewOptInt64(2),
			Name: api.NewOptString("User2"),
		},
	}}

	server, err := api.NewServer(&s)
	if err != nil {
		log.Fatalln(err)
	}

	err = http.ListenAndServe(":8080", server)
	if err != nil {
		log.Fatalln(err)
	}
}
