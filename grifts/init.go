package grifts

import (
	"github.com/barrongineer/plugin_test/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
