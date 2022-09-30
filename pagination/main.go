package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type HNResponse struct {
	Hits []HNHit `json:"hits"`
}

type HNHit struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func foo() (string, string) {
	a := "1"
	b := "2"

	return a, b
}

func main() {
	// ctx := context.Background()
	//
	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// })

	client := &http.Client{}

	res, err := client.Get("https://hn.algolia.com/api/v1/search?query=foo&tags=story")
	if err != nil {
		// alguma coisa deu merda
		log.Fatal(err)
	}

	text, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var hackerNewsResponse HNResponse
	err = json.Unmarshal(text, &hackerNewsResponse)
	if err != nil {
		log.Fatal(err)
	}

	// criar um UUID
	// colocar isso em uma lista no Redis

	id := uuid.NewString()

	fmt.Printf("ID: %s\nNumero de hits: %d\n", id, len(hackerNewsResponse.Hits))

	fmt.Printf("%+v\n", hackerNewsResponse)

	// err = rdb.LPush(ctx, id).Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}



