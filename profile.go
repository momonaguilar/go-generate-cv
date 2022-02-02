package main

import "time"

type Skills struct {
	Skill  string `json:"skill"`
	Rating int    `json:"rating"`
}

type Socials struct {
	Social string `json:"social"`
	Link   string `json:"link"`
}

type Experiences struct {
	YearFrom    int    `json:"yearFrom"`
	YearTo      int    `json:"yearTo"`
	Company     string `json:"company"`
	Role        string `json:"role"`
	Description string `json:"description"`
}

type Education struct {
	Year      string `json:"year"`
	Instution string `json:"institution"`
	Major     string `json:"major"`
	Awards    string `json:"awards"`
}

type Hobbies struct {
	Name    string `json:"name"`
	Details string `json:"details"`
	Awards  string `json:"awards"`
}

type Profile struct {
	Name            string        `json:"name"`
	Position        string        `json:"position"`
	Address         string        `json:"address"`
	Phone           string        `json:"phone"`
	EmailAddress    string        `json:"emailAddress"`
	PersonalWebsite string        `json:"personalWebsite"`
	URLToImage      string        `json:"urlToImage"`
	About           string        `json:"about"`
	Skills          []Skills      `json:"skills"`
	Socials         []Socials     `json:"socials"`
	Experiences     []Experiences `json:"experiences"`
	Education       []Education   `json:"education"`
	Hobbies         []Hobbies     `json:"hobbies"`
	PublishedAt     time.Time     `json:"publishedAt"`
}
