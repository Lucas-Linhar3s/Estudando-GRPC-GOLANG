package main

import (
	"context"
	"encoding/json"
	"fmt"
	user "grpc/proto/gen"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
}

type User struct {
	Username   string `json:"login"`
	ID         int64  `json:"id"`
	Node_Id    string `json:"node_Id"`
	AvatarURL  string `json:"avatar_url"`
	URL        string `json:"url"`
	Followers  int64  `json:"followers"`
	Following  int64  `json:"following"`
	Gists      int64  `json:"public_gists"`
	Name       string `json:"name"`
	Location   string `json:"location"`
	Repos      int64  `json:"public_repos"`
	Bio        string `json:"bio"`
	StarredURL string `json:"starred_url"`
	Create_in  int64  `json:"create_In"`
	ReposURL   string `json:"repos_url"`
}

// felipeagger

//GetUser is get user on github
func (s *Server) GetUser(ctx context.Context, in *user.UserRequest) (*user.UserResponse, error) {
	log.Printf("Receive message from client: %s", in.Username)

	res, err := http.Get(fmt.Sprintf("https://api.github.com/users/%v", in.Username))
	if err != nil {
		log.Fatal(err)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	usr := User{}
	jsonErr := json.Unmarshal(body, &usr)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return &user.UserResponse{
		Username:  usr.Username,
		Id:        usr.ID,
		NodeId:    usr.Node_Id,
		Avatarurl: usr.AvatarURL,
		Statistics: &user.Statistics{
			Followers: usr.Followers,
			Following: usr.Following,
			Repos:     usr.Repos,
			Gists:     usr.Gists,
		},
		Name:      usr.Name,
		Location:  usr.Location,
		Bio:       usr.Bio,
		CreatedAt: usr.Create_in,
		ListURLs:  []string{usr.URL, usr.StarredURL, usr.ReposURL},
	}, nil
}
