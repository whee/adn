// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

type Entities struct {
	Mentions []Mention
	Hashtags []Hashtag
	Links    []Link
}

type Mention struct {
	Name string `json:"name"`
	Id   string `json:"id"`
	Pos  int    `json:"pos"`
	Len  int    `json:"len"`
}

type Hashtag struct {
	Name string `json:"name"`
	Pos  int    `json:"pos"`
	Len  int    `json:"len"`
}

type Link struct {
	Text string `json:"text"`
	Url  string `json:"url"`
	Pos  int    `json:"pos"`
	Len  int    `json:"len"`
}
