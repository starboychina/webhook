package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type githubPayload struct {
	// ref
	Ref string `json:"ref"` // "refs/heads/master"
	// repository
	Repo struct {
		URL string `json:"url"` // "https://github.com/starboychina/webhook"
	} `json:"repository"`
}

func init() {

	http.HandleFunc("/github", githubHandle)

}

func githubHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "")
	defer req.Body.Close()
	decoder := json.NewDecoder(req.Body)

	var payload githubPayload
	err := decoder.Decode(&payload)
	if err != nil {
		log.Printf("payload json decode failed: %s\n", err)
		return
	}

	for _, hook := range config.Hooks {
		if strings.TrimRight(hook.Repo, "/") == payload.Repo.URL && strings.Contains(payload.Ref, hook.Branch) {
			fmt.Println(hook.Repo)
			executeShell(hook.Shell)
		}
	}

}
