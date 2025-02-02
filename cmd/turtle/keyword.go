package main

import (
	"fmt"
	"io"

	"github.com/devnoname120/turtle"
	"github.com/spf13/cobra"
)

func newKeywordCmd(w io.Writer) *cobra.Command {
	return &cobra.Command{
		Use:   "keyword [KEYWORD]",
		Short: "Print all emojis with the keyword",
		Long:  "Print all emojis with the keyword",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runKeyword(w, args[0])
		},
	}
}

func runKeyword(w io.Writer, k string) error {

	emojis := turtle.Keyword(k)

	if emojis == nil {
		return fmt.Errorf("cannot find emojis for keyword %q", k)
	}

	j, err := NewJSONWriter(w, WithIndent(prefix, indent))

	if err != nil {
		return fmt.Errorf("error creating JSONWriter: %v", err)
	}

	return j.WriteEmojis(emojis)
}
