package mlang

import (
	"fmt"
)

// Set is the set of messages for each language. The type M is usually string.
// Template is used for dynamical messages that is evaluated at a timing
// language specified.
type Set[M string | Template] map[any]M

func (s Set[M]) Get(lang any) (string, bool) {
	if msg, ok := s[lang]; ok {
		return eval(msg, lang)
	}

	return "", false
}

func (s Set[M]) MustGet(lang any) string {
	if msg, ok := s[lang]; ok {
		return mustEval(msg, lang)
	}

	// fallback to random language
	for l, msg := range s {
		return mustEval(msg, l) // use `l`, not `lang` to ensure using the same language.
	}

	panic("empty set")
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

func (tmp Template) eval(lang any) (string, bool) {
	args := make([]any, len(tmp.args))
	for i, arg := range tmp.args {
		switch arg := arg.(type) {
		case Set[string]:
			var ok bool
			args[i], ok = arg.Get(lang)
			if !ok {
				return "", false
			}
		case Set[Template]:
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

func (tmp Template) mustEval(lang any) string {
	args := make([]any, len(tmp.args))
	for i, arg := range tmp.args {
		switch arg := arg.(type) {
		case Set[string]:
			args[i] = arg.MustGet(lang)
		case Set[Template]:
			args[i] = arg.MustGet(lang)
		default:
			args[i] = arg
		}
	}
	return fmt.Sprintf(tmp.format, args...)
}

func eval(msg any, language any) (string, bool) {
	switch m := msg.(type) {
	case Template:
		return m.eval(language)
	case string:
		return m, true
	default:
		panic("unreachable")
	}
}

func mustEval(msg any, language any) string {
	switch m := msg.(type) {
	case Template:
		return m.mustEval(language)
	case string:
		return m
	default:
		panic("unreachable")
	}
}
