package grifts

import (
	"movie-suggestor-fanbase/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
