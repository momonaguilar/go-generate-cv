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
	Year        string `json:"year"`
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
	Name            string `json:"name"`
	Position        string `json:"position"`
	Address         string `json:"address"`
	Phone           string `json:"phone"`
	EmailAddress    string `json:"emailAddress"`
	PersonalWebsite string `json:"personalWebsite"`
	URLToImage      string `json:"urlToImage"`
	About           string `json:"about"`

	//skills
	Skills []Skills `json:"skills"`

	//socials
	Socials []Socials `json:"socials"`

	//experience
	Experiences []Experiences `json:"experiences"`

	//education
	Education []Education `json:"education"`

	//hobby
	Hobby []Hobbies `json:"hobbies"`

	PublishedAt time.Time `json:"publishedAt"`
}
