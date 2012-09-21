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
