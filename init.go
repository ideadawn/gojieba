package gojieba

import (
	"github.com/ideadawn/gojieba/deps/cppjieba"
	"github.com/ideadawn/gojieba/deps/limonp"
	"github.com/ideadawn/gojieba/dict"
)

func init() {
	dict.Init()
	limonp.Init()
	cppjieba.Init()
}
