package adn

type User struct {
	Id          string
	Username    string
	Name        string // The user-supplied Name may be a pseudonym.
	Description Description
	Timezone    string // The Timezone is in tzinfo format.
	Locale      string
	AvatarImage Image  // The URL and original size of the user's avatar.
	CoverImage  Image  // The URL and original size of the user's cover image.
	Type        string // An account can be human, bot, corporate, or feed.
	CreatedAt   string // The time at which the User was created in ISO 8601 format.
	Counts      Counts
	AppData     interface{} // TODO: Opaque information an application may store.
	FollowsYou  bool        // Does this user follow the user making the request? May be omitted if this is not an authenticated request.
	YouFollow   bool        // Does this user making the request follow this user? May be omitted if this is not an authenticated request.
	YouMuted    bool        // Has the user making the request blocked this user? May be omitted if this is not an authenticated request.
}

type Description struct {
	Text     string      // Biographical information
	Html     string      // Server-generated annotated HTML version of Text.
	Entities interface{} // TODO
}

type Image struct {
	Height, Width int
	Url           string
}

type Counts struct {
	Following int
	Followers int
	Posts     int
	Stars     int
}
