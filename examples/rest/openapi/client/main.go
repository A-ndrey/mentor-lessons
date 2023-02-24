package main

import (
	"context"
	"log"
	"openapi_example/api"
)

func main() {
	client, err := api.NewClient("http://localhost:3000")
	if err != nil {
		log.Fatalln(err)
	}

	users, err := client.UserGet(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(users)

	for _, u := range users {
		uId, _ := u.GetID().Get()
		user, err := client.UserIDGet(context.Background(), api.UserIDGetParams{ID: int(uId)})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(user)
	}
}
