package main

import (
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis"
)

type Podcast struct {
	Title  string   `bson:"title,omitempty"`
	Author string   `bson:"author,omitempty"`
	Tags   []string `bson:"tags,omitempty"`
}

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-18970.c280.us-central1-2.gce.cloud.redislabs.com:18970",
		Password: "XcmC8twt3LWsfZECWI4tUyu1IlvdvFiH", // no password set
		DB:       0,                                  // use default DB
	})

	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)

	podcast := Podcast{
		Title:  "The Polyglot Developer",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}

	json, err := json.Marshal(podcast)
	if err != nil {
		fmt.Println(err)
	}

	err = rdb.Set("ID1001", json, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Data has been successfully published")
}
