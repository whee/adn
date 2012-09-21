package adn

type Mention struct {
	Name string
	Id   string
	Pos  int
	Len  int
}

type Hashtag struct {
	Name string
	Pos  int
	Len  int
}

type Link struct {
	Text string
	Url  string
	Pos  int
	Len  int
}
