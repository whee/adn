package adn

import (
	"text/template"
)

const apiHost = "https://alpha-api.app.net"

type httpMethod int

const (
	httpGet httpMethod = iota
	httpPost
	httpDelete
)

type endpoint struct {
	Method httpMethod
	Path   string
}

var epTemplates = new(template.Template)

var apiEndpoints = map[string]endpoint{
	// User endpoints
	"retrieve user":        {httpGet, "/stream/0/users/{{.UserId}}"},
	"follow user":          {httpPost, "/stream/0/users/{{.UserId}}/follow"},
	"unfollow user":        {httpDelete, "/stream/0/users/{{.UserId}}/follow"},
	"list followed users":  {httpGet, "/stream/0/users/{{.UserId}}/following"},
	"list following users": {httpGet, "/stream/0/users/{{.UserId}}/followers"},
	"mute user":            {httpPost, "/stream/0/users/{{.UserId}}/mute"},
	"unmute user":          {httpDelete, "/stream/0/users/{{.UserId}}/mute"},
	"list muted users":     {httpGet, "/stream/0/users/me/muted"},
	"list reposters":       {httpGet, "/stream/0/posts/{{.PostId}}/reposters"},
	"list starrers":        {httpGet, "/stream/0/posts/{{.PostId}}/star"},
}

func init() {
	for k, v := range apiEndpoints {
		template.Must(epTemplates.New(k).Parse(v.Path))
	}
}
