// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

import (
	"fmt"
	"text/template"
)

const (
	apiHost     = "https://alpha-api.app.net"
	apiAuthHost = "https://alpha.app.net"
)

type httpMethod string

const (
	httpGet    httpMethod = "GET"
	httpPost              = "POST"
	httpDelete            = "DELETE"
)

type endpoint struct {
	Method  httpMethod
	Path    string
	Options *epOptions
}

var apiEndpoints = map[string]endpoint{
	// Users
	"retrieve user":        {httpGet, apiHost + "/stream/0/users/{{.User}}", nil},
	"follow user":          {httpPost, apiHost + "/stream/0/users/{{.User}}/follow", nil},
	"unfollow user":        {httpDelete, apiHost + "/stream/0/users/{{.User}}/follow", nil},
	"list followed users":  {httpGet, apiHost + "/stream/0/users/{{.User}}/following", nil},
	"list following users": {httpGet, apiHost + "/stream/0/users/{{.User}}/followers", nil},
	"mute user":            {httpPost, apiHost + "/stream/0/users/{{.User}}/mute", nil},
	"unmute user":          {httpDelete, apiHost + "/stream/0/users/{{.User}}/mute", nil},
	"list muted users":     {httpGet, apiHost + "/stream/0/users/me/muted", nil},
	"search for users":     {httpGet, apiHost + "/stream/0/users/search", nil},
	"list reposters":       {httpGet, apiHost + "/stream/0/posts/{{.Post}}/reposters", nil},
	"list starrers":        {httpGet, apiHost + "/stream/0/posts/{{.Post}}/star", nil},

	// Tokens
	"check current token": {httpGet, apiHost + "/stream/0/token", nil},

	// Posts
	"create post":                       {httpPost, apiHost + "/stream/0/posts", nil},
	"retrieve post":                     {httpGet, apiHost + "/stream/0/posts/{{.Post}}", nil},
	"delete post":                       {httpDelete, apiHost + "/stream/0/posts/{{.Post}}", nil},
	"retrieve post replies":             {httpGet, apiHost + "/stream/0/posts/{{.Post}}/replies", nil},
	"retrieve user posts":               {httpGet, apiHost + "/stream/0/users/{{.User}}/posts", nil},
	"repost post":                       {httpPost, apiHost + "/stream/0/posts/{{.Post}}/repost", nil},
	"unrepost post":                     {httpDelete, apiHost + "/stream/0/posts/{{.Post}}/repost", nil},
	"star post":                         {httpPost, apiHost + "/stream/0/posts/{{.Post}}/star", nil},
	"unstar post":                       {httpDelete, apiHost + "/stream/0/posts/{{.Post}}/star", nil},
	"retrieve user starred posts":       {httpGet, apiHost + "/stream/0/users/{{.User}}/stars", nil},
	"retrieve posts mentioning user":    {httpGet, apiHost + "/stream/0/users/{{.User}}/mentions", nil},
	"retrieve user personalized stream": {httpGet, apiHost + "/stream/0/posts/stream", nil},
	"retrieve global stream":            {httpGet, apiHost + "/stream/0/posts/stream/global", nil},
	"retrieve tagged posts":             {httpGet, apiHost + "/stream/0/posts/tag/{{.Hashtag}}", nil},

	// Streams
	"retrieve realtime user personalized stream":          {httpGet, apiHost + "/stream/0/streams/user", nil},
	"retrieve realtime multiple user personalized stream": {httpGet, apiHost + "/stream/0/streams/app", nil},
	"retrieve realtime public stream":                     {httpGet, apiHost + "/stream/0/streams/public", nil},
	"retrieve stream status":                              {httpGet, apiHost + "/stream/0/streams/{{.Stream}}", nil},
	"control stream":                                      {httpPost, apiHost + "/stream/0/streams/{{.Stream}}", nil},

	// Real-time updates
	"list subscriptions":       {httpGet, apiHost + "/stream/0/subscriptions", nil},
	"create subscription":      {httpPost, apiHost + "/stream/0/subscriptions", nil},
	"delete subscription":      {httpDelete, apiHost + "/stream/0/subscriptions/{{.Subscription}}", nil},
	"delete all subscriptions": {httpDelete, apiHost + "/stream/0/subscriptions", nil},

	// Filters
	"retrieve current user filters": {httpGet, apiHost + "/stream/0/filters", nil},
	"create filter":                 {httpPost, apiHost + "/stream/0/filters", nil},
	"retrieve filter":               {httpGet, apiHost + "/stream/0/filters/{{.Filter}}", nil},
	"delete filter":                 {httpGet, apiHost + "/stream/0/filters/{{.Filter}}", nil},

	// Authentication (Server-side flow)
	"authentication url": {httpGet, apiAuthHost + "/oauth/authenticate?client_id={{urlquery .Id}}&response_type=code&redirect_uri={{urlquery .RedirectURI}}&scope={{urlquery .Scopes.Spaced}}{{if .State}}&state={{urlquery .State}}{{end}}", nil},
	"get access token":   {httpPost, apiAuthHost + "/oauth/access_token", &epOptions{ResponseEnvelope: false}},
}

type EpArgs struct {
	User, Post, Hashtag, Stream, Subscription, Filter string
}

type epOptions struct {
	ResponseEnvelope bool // Do we expect a response envelope?
}

var epTemplates = new(template.Template)

func init() {
	for k, v := range apiEndpoints {
		template.Must(epTemplates.New(k).Parse(v.Path))
	}
}

type responseEnvelope struct {
	Data interface{}  `json:"data"`
	Meta responseMeta `json:"meta"`
}

type responseMeta struct {
	Code         int    `json:"code"`
	ErrorId      string `json:"error_id"`
	ErrorMessage string `json:"error_message"`
}

type APIError responseMeta

func (e APIError) Error() string {
	return fmt.Sprintf("%d: %s (%s)", e.Code, e.ErrorMessage, e.ErrorId)
}
