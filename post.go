package adn

import (
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	Text      string    `json:"text"`
	HTML      string    `json:"html"`
	Source    Source    `json:"source"`
	ReplyTo   string    `json:"reply_to"`
	ThreadId  string    `json:"thread_id"`

	NumReplies int `json:"num_replies"`
	NumStars   int `json:"num_stars"`
	NumReposts int `json:"num_reposts"`

	Annotations []interface{} `json:"annotations"` // TODO implement

	Entities Entities `json:"entities"`

	IsDeleted   bool `json:"is_deleted"`
	MachineOnly bool `json:"machine_only"`

	YouStarred bool     `json:"you_starred"`
	StarredBy  []string `json:"starred_by"` // TODO verify

	YouReposted bool     `json:"you_reposted"`
	Reposters   []string `json:"reposters"` // TODO verify
	RepostOf    *Post    `json:"repost_of"` // TODO verify self-reference
}

type Source struct {
	Name string `json:"name"`
	Link string `json:"link"`
}

func (p *Post) Get(id string) error {
	return epExecuteGet("retrieve post", epArgs{Post: id}, p)
}
