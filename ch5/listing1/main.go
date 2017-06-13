package main

import (
	"fmt"
)

type user struct {
	name       string
	email      string
	ext        int
	privileged bool
}

type admin struct {
	person user
	level  string
}

// Duration is new type
type Duration int64

func typeD() {
	var bill user

	lisa := user{
		name:       "Lisa",
		email:      "lisa@email.com",
		ext:        123,
		privileged: true,
	}

	john := user{"John", "john@email.com", 333, true}

	fmt.Printf("%s's email address is '%s'.\n", lisa.name, lisa.email)
	fmt.Printf("%s's email address is '%s'.\n", john.name, john.email)
	fmt.Printf("%s's email address is '%s'.\n", bill.name, bill.email)

	fred := admin{
		person: user{
			name:       "Scott",
			email:      "scott@email.com",
			ext:        888,
			privileged: true,
		},
		level: "super",
	}

	fmt.Printf("%v's level is %s.\n", fred.person, fred.level)
	fmt.Printf("%v\n", fred)
}

func (u user) notify() {
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	typeD()

	var dur Duration
	dur = Duration(1000)

	fmt.Printf("dur is: %d\n", dur)

	george := user{"George", "george@email.com", 999, true}
	george.notify()

	lisa := user{"Lisa", "lisa@email.com", 666, true}
	lisa.notify()

	george.changeEmail("george@newton.com")
	george.notify()

	lisa.changeEmail("lisa@163.com")
	lisa.notify()

	anna := &user{"Anna", "anna@email.com", 567, true}
	anna.notify()
}
