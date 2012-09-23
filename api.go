package adn

import (
	"bytes"
	"io"
	"net/http"
	"text/template"
)

const apiHost = "https://alpha-api.app.net"

type httpMethod string

const (
	httpGet    httpMethod = "GET"
	httpPost              = "POST"
	httpDelete            = "DELETE"
)

type endpoint struct {
	Method httpMethod
	Path   string
}

var epTemplates = new(template.Template)

var apiEndpoints = map[string]endpoint{
	// Users
	"retrieve user":        {httpGet, "/stream/0/users/{{.User}}"},
	"follow user":          {httpPost, "/stream/0/users/{{.User}}/follow"},
	"unfollow user":        {httpDelete, "/stream/0/users/{{.User}}/follow"},
	"list followed users":  {httpGet, "/stream/0/users/{{.User}}/following"},
	"list following users": {httpGet, "/stream/0/users/{{.User}}/followers"},
	"mute user":            {httpPost, "/stream/0/users/{{.User}}/mute"},
	"unmute user":          {httpDelete, "/stream/0/users/{{.User}}/mute"},
	"list muted users":     {httpGet, "/stream/0/users/me/muted"},
	"list reposters":       {httpGet, "/stream/0/posts/{{.Post}}/reposters"},
	"list starrers":        {httpGet, "/stream/0/posts/{{.Post}}/star"},

	// Tokens
	"check current token": {httpGet, "/stream/0/token"},

	// Posts
	"create post":                       {httpPost, "/stream/0/posts"},
	"retrieve post":                     {httpGet, "/stream/0/posts/{{.Post}}"},
	"delete post":                       {httpDelete, "/stream/0/posts/{{.Post}}"},
	"retrieve post replies":             {httpGet, "/stream/0/posts/{{.Post}}/replies"},
	"retrieve user posts":               {httpGet, "/stream/0/users/{{.User}}/posts"},
	"repost post":                       {httpPost, "/stream/0/posts/{{.Post}}/repost"},
	"unrepost post":                     {httpDelete, "/stream/0/posts/{{.Post}}/repost"},
	"star post":                         {httpPost, "/stream/0/posts/{{.Post}}/star"},
	"unstar post":                       {httpDelete, "/stream/0/posts/{{.Post}}/star"},
	"retrieve user starred posts":       {httpGet, "/stream/0/users/{{.User}}/stars"},
	"retrieve posts mentioning user":    {httpGet, "/stream/0/users/{{.User}}/mentions"},
	"retrieve user personalized stream": {httpGet, "/stream/0/posts/stream"},
	"retrieve global stream":            {httpGet, "/stream/0/posts/stream/global"},
	"retrieve tagged posts":             {httpGet, "/stream/0/posts/tag/{{.Hashtag}}"},

	// Streams
	"retrieve realtime user personalized stream":          {httpGet, "/stream/0/streams/user"},
	"retrieve realtime multiple user personalized stream": {httpGet, "/stream/0/streams/app"},
	"retrieve realtime public stream":                     {httpGet, "/stream/0/streams/public"},
	"retrieve stream status":                              {httpGet, "/stream/0/streams/{{.Stream}}"},
	"control stream":                                      {httpPost, "/stream/0/streams/{{.Stream}}"},

	// Real-time updates
	"list subscriptions":       {httpGet, "/stream/0/subscriptions"},
	"create subscription":      {httpPost, "/stream/0/subscriptions"},
	"delete subscription":      {httpDelete, "/stream/0/subscriptions/{{.Subscription}}"},
	"delete all subscriptions": {httpDelete, "/stream/0/subscriptions"},

	// Filters
	"retrieve current user filters": {httpGet, "/stream/0/filters"},
	"create filter":                 {httpPost, "/stream/0/filters"},
	"retrieve filter":               {httpGet, "/stream/0/filters/{{.Filter}}"},
	"delete filter":                 {httpGet, "/stream/0/filters/{{.Filter}}"},
}

type epArgs struct {
	User, Post, Hashtag, Stream, Subscription, Filter string
}

func init() {
	for k, v := range apiEndpoints {
		template.Must(epTemplates.New(k).Parse(v.Path))
	}
}

var apiClient = &http.Client{}

func epExecute(name string, args epArgs) (body io.ReadCloser, err error) {
	var path bytes.Buffer

	err = epTemplates.ExecuteTemplate(&path, name, args)
	if err != nil {
		return
	}

	method := apiEndpoints[name].Method
	url := apiHost + path.String()
	req, err := http.NewRequest(string(method), url, nil)
	if err != nil {
		return
	}

	resp, err := apiClient.Do(req)
	if err != nil {
		return
	}
	body = resp.Body

	return
}
