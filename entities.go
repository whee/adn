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
	Name string `json:"name"` // The username being mentioned (not including '@').
	Id   string `json:"id"`   // The user id of the mentioned user.
	Pos  int    `json:"pos"`  // The 0-based index where the entity includes Text (includes '@').
	Len  int    `json:"len"`  // The length of the substring in Text.
}

type Hashtag struct {
	Name string `json:"name"` // The name of the hashtag (not including '@').
	Pos  int    `json:"pos"`  // The 0-based index where the entity includes Text (includes '#').
	Len  int    `json:"len"`  // The length of the substring in Text.
}

type Link struct {
	Text string `json:"text"` // The anchor text to be linked.
	Url  string `json:"url"`  // The destination url.
	Pos  int    `json:"pos"`  // The 0-based index where the entity begins in Text.
	Len  int    `json:"len"`  // The length of the substring in Text that represents this link.
}
