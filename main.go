package main

import (
	"fmt"
	"github.com/thecsw/mira"
)

func main() {
	configFilePath := "./rcredentials.txt"
	r, err := mira.Init(mira.ReadCredsFromFile(configFilePath))
	if err != nil {
		panic(fmt.Sprintf("Encountered an error trying to load config file %s, err: %s", configFilePath, err))
	}
	fmt.Println("HELLO")
	// c, _, _ := r.Subreddit("all").StreamSubmissions()
	fmt.Println(r.Me().Info())
	// for {
	// 	post := <-c
	// 	id := post.GetId()
	// 	fmt.Println(r.Submission(id).Info())
	// }
}
