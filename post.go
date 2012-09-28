// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

import (
	"time"
)

type Post struct {
	Id        string    `json:"id"`
	User      User      `json:"user"`       // Embedded version of the User object. May not exist if the account has since been deleted.
	CreatedAt time.Time `json:"created_at"` // The time at which the post was created.
	Text      string    `json:"text"`       // User supplied text of the post.
	HTML      string    `json:"html"`       // Server-generated annotated HTML rendering of the post text.
	Source    Source    `json:"source"`
	ReplyTo   string    `json:"reply_to"`  // Id of the post this post is replying to (or "" if not a reply).
	ThreadId  string    `json:"thread_id"` // Id of the post at the root of the thread this post is a part of.

	NumReplies int `json:"num_replies"` // The number of posts created in reply to this post.
	NumStars   int `json:"num_stars"`   // The number of users who have starred this post.
	NumReposts int `json:"num_reposts"` // The number of users who have reposted this post.

	Annotations []interface{} `json:"annotations"` // TODO implement

	Entities Entities `json:"entities"`

	IsDeleted   bool `json:"is_deleted"`   // Has this post been deleted?
	MachineOnly bool `json:"machine_only"` // Is this post meant for machines or humans?

	YouStarred bool     `json:"you_starred"` // Has the current user starred this post?
	StarredBy  []string `json:"starred_by"`  // A partial list of users who have starred this post. (TODO verify)

	YouReposted bool     `json:"you_reposted"` // Has the current user reposted this post?
	Reposters   []string `json:"reposters"`    // A partial list of users who have reposted this post. (TODO verify)
	RepostOf    *Post    `json:"repost_of"`    // If this post is a repost, the original post. (TODO verify self-reference)
}

type Source struct {
	Name string `json:"name"` // Description of the API consumer that created this post.
	Link string `json:"link"` // Link provided by the API consumer that created this post.
}

// Retrieve the post specified by id using token as authentication.
func (c *Application) GetPost(token string, id string) (p *Post, err error) {
	p = &Post{}
	err = c.Do(&Request{Token: token}, "retrieve post", EpArgs{Post: id}, p)
	return
}

// Calls GetPost on the DefaultApplication.
func GetPost(token string, id string) (*Post, error) {
	return DefaultApplication.GetPost(token, id)
}
