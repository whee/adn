// Copyright (c) 2012 Brian Hetro <whee@smaertness.net>
// Use of this source code is governed by the ISC
// license that can be found in the LICENSE file.

package adn

type Filter struct {
	Type           string   `json:"type"`             // Either "show" or "block" for whether this filter should exclude everything except for what's shown or show everything except for what's blocked.
	Name           string   `json:"name"`             // A User assigned name for this filter.
	UserIds        []string `json:"user_ids"`         // A slice of user ids a Post must or must not be created by.
	Hashtags       []string `json:"hashtags"`         // A slice of hashtags a Post must or must not have.
	LinkDomains    []string `json:"link_domains"`     // A slice of domains a Post must or must not have a link to.
	MentionUserIds []string `json:"mention_user_ids"` // A slice of user ids a Post must or must not mention.
}
