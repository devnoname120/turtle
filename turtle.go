package turtle

// Version of the turtle library
const Version = "v0.2.0"

// Emojis maps a name to an Emoji
var Emojis = make(map[string]*Emoji)

// EmojisByChar maps a character to an Emoji
var EmojisByChar = make(map[string]*Emoji)

func init() {
	for _, e := range emojis {
		Emojis[e.Name] = e
		EmojisByChar[e.Char] = e
	}
}

// Search emojis by a name
func Search(s string) []*Emoji {
	return search(emojis, s)
}

// Slug filters the emojis by a slug
func Slug(s string) []*Emoji {
	return slug(emojis, s)
}

// Keyword filters the emojis by a keyword
func Keyword(k string) []*Emoji {
	return keyword(emojis, k)
}

// Category filters the emojis by a category
func Category(c string) []*Emoji {
	return category(emojis, c)
}

// Filter the emojis based on the given comparison function
func Filter(f func(e *Emoji) bool) []*Emoji {
	return filter(emojis, f)
}
