package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"net/http"
"os"
)

type ReleaseInfo struct {
	Id      uint   `json:"id"`
	Tagname string `json:"tag_name"`
}

type GithubReleaseInfoer struct{}

// the following program will query github repo for latest release based on some tag info.

func main() {
	gh:=GithubReleaseInfoer{}
	msg, err := getReleaseTagMessag(gh,"docker/machine")
	if err != nil {

		fmt.Println("error occured", os.Stderr, nil)
	}
	fmt.Sprintf(" the message is ", msg)
}

// Method to actually query Github API for release info.
// this method returns the tag of the github repo and error if any

func (gh GithubReleaseInfoer) GetLatestreleaseTag(repo string) (string, error) {
	apiUrl := fmt.Sprintf("https://api.github.com/repos/%s/releases", repo)
	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	release := []ReleaseInfo{}
	if err := json.Unmarshal(body, &release); err != nil {
		return "", err
	}
	tag := release[0].Tagname
	return tag, nil // here we got the tag name .
}

//ReleaseInformer this interface implements - GetLatestreleaseTag Methodwhich returns tag
type ReleaseInformer interface {
	GetLatestreleaseTag(string) (string, error)
}

// we will pass the - ReleaseInformer- Interface to - getReleaseTagMessage - function
func getReleaseTagMessag(ri ReleaseInformer, repo string) (string, error) {
	tag, err := ri.GetLatestreleaseTag(repo)
	if err != nil {
		return "", fmt.Errorf("error querrying github API", err)
	}
	return tag, nil // otherwise send the tag
}


