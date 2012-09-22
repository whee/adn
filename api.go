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
	// Users
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

	// Tokens
	"check current token": {httpGet, "/stream/0/token"},

	// Posts
	"create post":                       {httpPost, "/stream/0/posts"},
	"retrieve post":                     {httpGet, "/stream/0/posts/{{.PostId}}"},
	"delete post":                       {httpDelete, "/stream/0/posts/{{.PostId}}"},
	"retrieve post replies":             {httpGet, "/stream/0/posts/{{.PostId}}/replies"},
	"retrieve user posts":               {httpGet, "/stream/0/users/{{.UserId}}/posts"},
	"repost post":                       {httpPost, "/stream/0/posts/{{.PostId}}/repost"},
	"unrepost post":                     {httpDelete, "/stream/0/posts/{{.PostId}}/repost"},
	"star post":                         {httpPost, "/stream/0/posts/{{.PostId}}/star"},
	"unstar post":                       {httpDelete, "/stream/0/posts/{{.PostId}}/star"},
	"retrieve user starred posts":       {httpGet, "/stream/0/users/{{.UserId}}/stars"},
	"retrieve posts mentioning user":    {httpGet, "/stream/0/users/{{.UserId}}/mentions"},
	"retrieve user personalized stream": {httpGet, "/stream/0/posts/stream"},
	"retrieve global stream":            {httpGet, "/stream/0/posts/stream/global"},
	"retrieve tagged posts":             {httpGet, "/stream/0/posts/tag/{{.Hashtag}}"},

	// Streams
	"retrieve realtime user personalized stream":          {httpGet, "/stream/0/streams/user"},
	"retrieve realtime multiple user personalized stream": {httpGet, "/stream/0/streams/app"},
	"retrieve realtime public stream":                     {httpGet, "/stream/0/streams/public"},
	"retrieve stream status":                              {httpGet, "/stream/0/streams/{{.StreamId}}"},
	"control stream":                                      {httpPost, "/stream/0/streams/{{.StreamId}}"},

	// Real-time updates
	"list subscriptions":       {httpGet, "/stream/0/subscriptions"},
	"create subscription":      {httpPost, "/stream/0/subscriptions"},
	"delete subscription":      {httpDelete, "/stream/0/subscriptions/{{.SubscriptionId}}"},
	"delete all subscriptions": {httpDelete, "/stream/0/subscriptions"},

	// Filters
	"retrieve current user filters": {httpGet, "/stream/0/filters"},
	"create filter":                 {httpPost, "/stream/0/filters"},
	"retrieve filter":               {httpGet, "/stream/0/filters/{{.FilterId}}"},
	"delete filter":                 {httpGet, "/stream/0/filters/{{.FilterId}}"},
}

func init() {
	for k, v := range apiEndpoints {
		template.Must(epTemplates.New(k).Parse(v.Path))
	}
}
