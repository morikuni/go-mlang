package mlang

import (
	"fmt"
)

// Language can be any comparable types.
// Basically, golang.org/x/text/language.Tag is used, but you can use any your original types as well.
type Language = any

type Message interface {
	isMessage()
	Get(lang Language) (string, bool)
	MustGet(lang Language) string
}

// Dict is the set of messages for each language. The type M is usually string.
// Template is used for dynamical messages that is evaluated at a timing
// language specified.
type Dict[M string | Template] map[Language]M

func (d Dict[M]) isMessage() {}

func (d Dict[M]) Get(lang Language) (string, bool) {
	if msg, ok := d[lang]; ok {
		return eval(msg, lang)
	}

	return "", false
}

func (d Dict[M]) MustGet(lang Language) string {
	if msg, ok := d[lang]; ok {
		return mustEval(msg, lang)
	}

	// fallback to random language
	for l, msg := range d {
		return mustEval(msg, l) // use `l`, not `lang` to ensure using the same language.
	}

	panic("empty set")
}

func eval(msg any, lang Language) (string, bool) {
	switch m := msg.(type) {
	case Template:
		return m.eval(lang)
	case string:
		return m, true
	default:
		panic("unreachable")
	}
}

func mustEval(msg any, lang Language) string {
	switch m := msg.(type) {
	case Template:
		return m.mustEval(lang)
	case string:
		return m
	default:
		panic("unreachable")
	}
}

// Template is a message template.
type Template struct {
	format string
	args   []any
}

func NewTemplate(format string, args ...any) Template {
	return Template{
		format: format,
		args:   args,
	}
}

func (tmp Template) eval(lang Language) (string, bool) {
	args := make([]any, len(tmp.args))
	for i, arg := range tmp.args {
		switch arg := arg.(type) {
		case Message:
			var ok bool
			args[i], ok = arg.Get(lang)
			if !ok {
				return "", false
			}
		default:
			args[i] = arg
		}
	}
	return fmt.Sprintf(tmp.format, args...), true
}

func (tmp Template) mustEval(lang Language) string {
	args := make([]any, len(tmp.args))
	for i, arg := range tmp.args {
		switch arg := arg.(type) {
		case Message:
			args[i] = arg.MustGet(lang)
		default:
			args[i] = arg
		}
	}
	return fmt.Sprintf(tmp.format, args...)
}
