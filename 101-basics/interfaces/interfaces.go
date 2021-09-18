package main

import (
	"fmt"
	"time"
)

type User interface {
	Like()
	LikeCount() int
}

type user struct {
	Name    string `json:"name"`
	Surname string `json:""`
	Likes   []Like
}

type Like struct {
	ID   string
	Date time.Time
}

func (u *user) Like() {
	if u.Likes == nil {
		u.Likes = make([]Like, 0)
	}

	l := Like{
		ID:   "id",
		Date: time.Time{},
	}

	u.Likes = append(u.Likes, l)
}

func (u user) LikeCount() int {
	return len(u.Likes)
}

func main() {
	u1 := &user{
		Name:    "go",
		Surname: "turkiye",
	}

	u2 := &user{
		Name:    "go",
		Surname: "turkey",
	}

	u1.Like()
	u1.Like()
	u1.Like()
	u1.Like()

	u2.Like()
	u2.Like()
	u2.Like()
	u2.Like()
	u2.Like()

	if IsGreaterThan(u1, u2) {
		fmt.Printf("%s %s like count greater than %s %s", u1.Name, u1.Surname, u2.Name, u2.Surname)
	} else {
		fmt.Printf("%s %s like count greater than %s %s", u2.Name, u2.Surname, u1.Name, u1.Surname)
	}
}

func IsGreaterThan(u1, u2 User) bool {
	return u1.LikeCount() > u2.LikeCount()
}
