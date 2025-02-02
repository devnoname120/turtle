# turtle

Emojis for Go 😄🐢🚀

## Reference

Follow this link to view the reference documentation: [GoDoc Reference][godoc] 📝

## Installation

### Library

To install the **turtle** library run:

``go get github.com/devnoname120/turtle``

### CLI app

If you would also like to use the **turtle** CLI app run:

``go get github.com/devnoname120/turtle/cmd/turtle``

See the [turtle CLI][cli] README for more information.

## Usage

### Emoji lookup

``turtle.Emojis`` is a map which contains all emojis available in **turtle**.
You can use it to look up emoji by their name.

```go
package main

import (
	"fmt"
	"os"

	"github.com/devnoname120/turtle"
)

func main() {
	name := "turtle"
	emoji, ok := turtle.Emojis[name]

	if !ok {
		fmt.Fprintf(os.Stderr, "no emoji found for name: %v\n", name)
		os.Exit(1)
	}

	fmt.Printf("Name: %q\n", emoji.Name)
	fmt.Printf("Char: %s\n", emoji.Char)
	fmt.Printf("Category: %q\n", emoji.Category)
	fmt.Printf("Keywords: %q\n", emoji.Keywords)
}
```

```text
Name: "turtle"
Char: 🐢
Category: "animals_and_nature"
Keywords: ["animal" "slow" "nature" "tortoise"]
```

### Search

Use ``Search()`` to find all emojis with a name that contains the search string.

```go
package main

import (
	"fmt"
	"os"

	"github.com/devnoname120/turtle"
)

func main() {
	s := "computer"
	emojis := turtle.Search(s)

	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emojis found for search: %v\n", s)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", s, emojis)
}
```

```text
computer: [💻 🖱 🖥 ]
```

### Category

Use ``Category()`` to find all emojis of the specified category.

```go
package main

import (
	"fmt"
	"os"

	"github.com/devnoname120/turtle"
)

func main() {
	c := "travel_and_places"
	emojis := turtle.Category(c)

	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emojis found for category: %v\n", c)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", c, emojis)
}
```

```text
travel_and_places: [🚡 ✈️ 🚑 ]
```

### Keyword

Use ``Keyword()`` to find all emojis by a keyword.

```go
package main

import (
	"fmt"
	"os"

	"github.com/devnoname120/turtle"
)

func main() {
	k := "happy"
	emojis := turtle.Keyword(k)

	if emojis == nil {
		fmt.Fprintf(os.Stderr, "no emoji found for keyword: %v\n", k)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", k, emojis)
}
```

```text
happy: [😊 😁 😀 😂 ]
```

## Emojis

Emoji names, categories and keywords are based on the fantastic
[muan/emojilib][emojilib] keyword library 📖

At this point, the **turtle** project supports the emojis that are also
available on GitHub. See the [GitHub REST API documentation][github-api] for
more information.

## Issues

If you encounter any problems, please [file an issue][new-issue] along with a
detailed description.

## Contributing

Contributions are welcome, and they are greatly appreciated! Every little bit
helps, and credit will always be given.

## License

Distributed under the terms of the [MIT license][mit], turtle is free and
open source software.

[cli]: /cmd/turtle/README.md
[emojilib]: https://github.com/muan/emojilib
[github-api]: https://developer.github.com/v3/emojis/
[godoc]: https://godoc.org/github.com/devnoname120/turtle (See GoDoc Reference)
[mit]: /LICENSE
[new-issue]: https://github.com/devnoname120/turtle/issues/new
