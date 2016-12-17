package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"os/exec"
	"gopkg.in/dixonwille/wmenu.v2"
)

type Story struct {
	By    string `json:"by"`
	Title string `json:"title"`
	URL   string `json:"url"`
}

func getStory(id string) (*Story, error) {
	storyReq, err := http.Get("https://hacker-news.firebaseio.com/v0/item/" + id + ".json")
	defer storyReq.Body.Close()
	if err != nil {
		return nil, err
	}
	story := new(Story)

	json.NewDecoder(storyReq.Body).Decode(&story)

	return story, nil
}

func main() {
	menu := wmenu.NewMenu("Which story would you like to read?")

	fmt.Println("💻  Hacker News 💻 ")
	topReq, _ := http.Get("https://hacker-news.firebaseio.com/v0/topstories.json")
	defer topReq.Body.Close()
	topReqBody, err := ioutil.ReadAll(topReq.Body)
	if err != nil {
		fmt.Println(err)
	}

	idString := string(topReqBody)

	ids := strings.Split(idString[1:len(idString)-1], ",")

	for i := 0; i < 3; i++ {
		topStory, topStoryErr := getStory(ids[i])
		if topStoryErr != nil {
			fmt.Println(topStoryErr)
		}
		menu.Option("\x1b[35m"+ topStory.Title + "\x1b[0m", nil, false, func() error {
			openCmd := exec.Command("open",topStory.URL)
			openCmd.Output()
			return nil
		})
	}

	menuErr := menu.Run()
	if menuErr != nil {
		fmt.Println(menuErr)
	}
}
