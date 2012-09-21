package adn

type Filter struct {
	Type           string
	Name           string
	UserIds        []string
	Hashtags       []string
	LinkDomains    []string
	MentionUserIds []string
}
