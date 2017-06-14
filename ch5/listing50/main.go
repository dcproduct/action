package main

import (
	"fmt"
)

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.email)
}

type test struct {
	name  string
	email string
}

func (t *test) notify() {
	fmt.Printf("Sending test email to %s<%s>\n", t.name, t.email)
}

type admin struct {
	user
	//test
	level string
}

func main() {
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@email.com",
		},
		level: "super",
	}

	ad.user.notify()
	ad.notify()
}
