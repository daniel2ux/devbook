package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"webapp/src/config"
	"webapp/src/requests"
)

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Nick      string    `json:"nick"`
	CreateAt  time.Time `json:"createAt"`
	Followers []User    `json:"followers"`
	Follows   []User    `json:"follows"`
	Posts     []Post    `json:"posts"`
}

func GetFullUserByID(userID uint64, r *http.Request) (User, error) {
	userChanel := make(chan User)
	followerChanel := make(chan []User)
	followChanel := make(chan []User)
	postChanel := make(chan []Post)

	go getUserData(userChanel, userID, r)
	go getFollowersData(followerChanel, userID, r)
	go getFollowsData(followChanel, userID, r)
	go getPostsData(postChanel, userID, r)

	var (
		user      User
		followers []User
		follows   []User
		posts     []Post
	)

	for i := 0; i < 4; i++ {
		select {
		case userLoad := <-userChanel:
			if userLoad.ID == 0 {
				return User{}, errors.New("User load error")
			}
			user = userLoad

		case followersLoad := <-followerChanel:
			// if followersLoad == nil {
			// 	return User{}, errors.New("Followers load error")
			// }
			followers = followersLoad

		case followsLoad := <-followChanel:
			// if followsLoad == nil {
			// 	return User{}, errors.New("Follows load error")
			// }
			follows = followsLoad

		case postsLoad := <-postChanel:
			// if postsLoad == nil {
			// 	return User{}, errors.New("Posts load error")
			// }
			posts = postsLoad
		}
	}

	user.Followers = followers
	user.Follows = follows
	user.Posts = posts
	return user, nil
}

func getUserData(c chan<- User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		c <- User{}
		return
	}
	defer response.Body.Close()

	var user User
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		c <- User{}
		return
	}

	c <- user
}

func getFollowersData(c chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/followers", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		c <- nil
		return
	}
	defer response.Body.Close()

	var followers []User
	if err := json.NewDecoder(response.Body).Decode(&followers); err != nil {
		c <- nil
		return
	}

	c <- followers
}

func getFollowsData(c chan<- []User, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/follow", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		c <- nil
		return
	}
	defer response.Body.Close()

	var follows []User
	if err := json.NewDecoder(response.Body).Decode(&follows); err != nil {
		c <- nil
		return
	}

	c <- follows
}

func getPostsData(c chan<- []Post, userID uint64, r *http.Request) {
	url := fmt.Sprintf("%s/users/%d/posts", config.APIURL, userID)
	response, err := requests.RequestWithAuth(r, http.MethodGet, url, nil)
	if err != nil {
		c <- nil
		return
	}
	defer response.Body.Close()

	var posts []Post
	if err := json.NewDecoder(response.Body).Decode(&posts); err != nil {
		c <- nil
		return
	}

	c <- posts
}
