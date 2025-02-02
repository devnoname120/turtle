# turtle

CLI for the [turtle library][library] 😄🐢💻

## Installation

``go get github.com/devnoname120/turtle/cmd/turtle``

## Usage

```text
Print the emoji with the specified name identifier

Usage:
  turtle [NAME] [flags]
  turtle [command]

Available Commands:
  category    Print all emojis of the category
  char        Print the emoji for the emoji character
  help        Help about any command
  keyword     Print all emojis with the keyword
  list        Print a list of values from the turtle library
  random      Print a random emoji
  search      Print emojis with a name that contains the search string

Flags:
  -h, --help            help for turtle
  -i, --indent string   indent for JSON output
  -p, --prefix string   prefix for JSON output
  -v, --version         version for turtle

Use "turtle [command] --help" for more information about a command.
```

### Emoji lookup

```bash
turtle -i "  " rocket
```

```json
{
  "name": "rocket",
  "category": "travel_and_places",
  "char": "🚀",
  "keywords": [
    "launch",
    "ship",
    "staffmode",
    "NASA",
    "outer space",
    "outer_space",
    "fly"
  ]
}
```

### Search

```bash
turtle -i "  " search computer
```

```json
[
  {
    "name": "computer",
    "category": "objects",
    "char": "💻",
    "keywords": [
      "technology",
      "laptop",
      "screen",
      "display",
      "monitor"
    ]
  },
  {
    "name": "computer_mouse",
    "category": "objects",
    "char": "🖱",
    "keywords": [
      "click"
    ]
  },
  {
    "name": "desktop_computer",
    "category": "objects",
    "char": "🖥",
    "keywords": [
      "technology",
      "computing",
      "screen"
    ]
  }
]
```

### Char

```bash
turtle -i "  " char 🐝
```

```json
{
  "name": "honeybee",
  "category": "animals_and_nature",
  "char": "🐝",
  "keywords": [
    "animal",
    "insect",
    "nature",
    "bug",
    "spring",
    "honey"
  ]
}
```

### Category

```bash
turtle -i "  " category travel_and_places
```

```json
[
  {
    "name": "aerial_tramway",
    "category": "travel_and_places",
    "char": "🚡",
    "keywords": [
      "transportation",
      "vehicle",
      "ski"
    ]
  },
  {
    "name": "airplane",
    "category": "travel_and_places",
    "char": "✈️",
    "keywords": [
      "vehicle",
      "transportation",
      "flight",
      "fly"
    ]
  },
  {
    "name": "ambulance",
    "category": "travel_and_places",
    "char": "🚑",
    "keywords": [
      "health",
      "911",
      "hospital"
    ]
  }
]
```

### Keyword

```bash
turtle -i "  " keyword happy
```

```json
[
  {
    "name": "blush",
    "category": "people",
    "char": "😊",
    "keywords": [
      "face",
      "smile",
      "happy",
      "flushed",
      "crush",
      "embarrassed",
      "shy",
      "joy"
    ]
  },
  {
    "name": "grin",
    "category": "people",
    "char": "😁",
    "keywords": [
      "face",
      "happy",
      "smile",
      "joy",
      "kawaii"
    ]
  },
  {
    "name": "grinning",
    "category": "people",
    "char": "😀",
    "keywords": [
      "face",
      "smile",
      "happy",
      "joy",
      ":D",
      "grin"
    ]
  }
]
```

### Random

```bash
turtle -i "  " random
```

```json
{
  "name": "woman_technologist",
  "category": "people",
  "char": "👩‍💻",
  "keywords": [
    "coder",
    "developer",
    "engineer",
    "programmer",
    "software",
    "woman",
    "human"
  ]
}
```

### List names, categories, keywords

Use the ``list`` subcommand to get information about emoji names, categories,
and keywords.

```bash
turtle -i "  " list categories
```

```json
[
    "activity",
    "animals_and_nature",
    "flags",
    "food_and_drink",
    "objects",
    "people",
    "symbols",
    "travel_and_places"
]
```

[library]: ../../README.md
