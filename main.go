package main

import (
	"fmt"
	"github.com/GnawNom/riposte/store"
	"github.com/GnawNom/riposte/hash"
	"github.com/thecsw/mira"
	"strings"
)

var supportedImageTypes map[string]bool = map[string]bool{"jpg": true, "jpeg": true, "png": true}

func main() {
	configFilePath := "./rcredentials.txt"
	r, err := mira.Init(mira.ReadCredsFromFile(configFilePath))
	if err != nil {
		panic(fmt.Sprintf("Encountered an error trying to load config file %s, err: %s", configFilePath, err))
	}
	fmt.Println("HELLO")
	subredditName := "all"
	c, _, _ := r.Subreddit(subredditName).StreamSubmissions()
	// c, _, _ := r.Subreddit("BikiniBottomTwitter").StreamSubmissions()
	info, err := r.Me().Info()
	fmt.Printf("%+v \n", info)
	hashStore := store.NewVPTreeHashStore(subredditName)

	for {
		fmt.Println("-------------------------")
		post := <-c
		id := post.GetId()
		info, err := r.Submission(id).Info()
		if err != nil {
			fmt.Println("encountered err when attempting to fetch submission: ", err)
		}
		postUrl := info.GetUrl()
		fmt.Printf("%+v \n", postUrl)
		tokens := strings.Split(info.GetUrl(), ".")

		if len(tokens) > 0 {
			fileType := tokens[len(tokens)-1]
			if _, ok := supportedImageTypes[fileType]; ok {
				hash, _ := hash.GeneratePHash(postUrl, fileType)
				fmt.Println("hash: ", hash)
				hashStore.FindDuplicate(store.PerceptionHash{hash}, 21)
			}
		}

	}
}
