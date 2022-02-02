package main

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func TestTormHandler(t *testing.T) {

	profile1 := &Profile{
		Name:            "John Wick",
		Position:        "CEO",
		Address:         "New York City, USA",
		Phone:           "+601234567",
		EmailAddress:    "john.wick@mail.com",
		PersonalWebsite: "https://johnwick.com",
		URLToImage:      "",
		About:           "Profile of John Wick, CEO of Heaven and Earth Corp.",
		Skills: []Skills{
			{Skill: "HTML", Rating: 100},
			{Skill: "Go", Rating: 75},
			{Skill: "Angular", Rating: 60},
		},
		Socials: []Socials{
			{Social: "Facebook", Link: "https://facebook.com/johnwick"},
			{Social: "Twitter", Link: "https://twitter.com/johnwick"},
		},
		Experiences: []Experiences{
			{YearFrom: 2000, YearTo: 2010, Company: "NCR Corporation", Role: "CEO", Description: "Performed leadership and management"},
			{YearFrom: 2010, YearTo: 2020, Company: "IBM Corporation", Role: "CEO", Description: "Performed leadership and management"},
		},
		Hobbies: []Hobbies{
			{Name: "Swimming", Details: "Competitive swimming", Awards: "Olympic Medalist"},
			{Name: "Badminton", Details: "Leisure sports", Awards: "Company wide champion"},
		},
		PublishedAt: time.Now(),
	}

	buf := &bytes.Buffer{}
	err := tpl.Execute(buf, profile1)
	if err != nil {
		// handle error
		fmt.Printf("TEST: Error occurred")
	}

}
