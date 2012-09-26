// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

type Filter struct {
	Type           string   `json:"type"`
	Name           string   `json:"name"`
	UserIds        []string `json:"user_ids"`
	Hashtags       []string `json:"hashtags"`
	LinkDomains    []string `json:"link_domains"`
	MentionUserIds []string `json:"mention_user_ids"`
}
