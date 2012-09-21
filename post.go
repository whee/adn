package adn

type Post struct {
	Id        string
	User      User
	CreatedAt string
	Text      string
	Html      string
	Source    interface{} // TODO implement
	ReplyTo   string
	ThreadId  string

	NumReplies int
	NumStars   int
	NumReposts int

	Annotations []interface{} // TODO implement

	Entities interface{} // TODO implement

	IsDeleted   bool
	MachineOnly bool

	YouStarred bool
	StarredBy  []string // TODO verify

	YouReposted bool
	Reposters   []string // TODO verify
	RepostOf    *Post    // TODO verify self-reference
}

type Source struct {
	Name string
	Link string
}
