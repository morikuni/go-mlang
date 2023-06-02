package mlang

import (
	"fmt"
)

type Message struct {
	format string
	args   []any
}

func Messagef(format string, args ...any) Message {
	return Message{
		format: format,
		args:   args,
	}
}

func (m Message) eval(lang any) (string, bool) {
	args := make([]any, len(m.args))
	for i, arg := range m.args {
		switch arg := arg.(type) {
		case Set[string]:
			var ok bool
			args[i], ok = arg.Get(lang)
			if !ok {
				return "", false
			}
		case Set[Message]:
			var ok bool
			args[i], ok = arg.Get(lang)
			if !ok {
				return "", false
			}
		default:
			args[i] = arg
		}
	}
	return fmt.Sprintf(m.format, args...), true
}

func (m Message) mustEval(lang any) string {
	args := make([]any, len(m.args))
	for i, arg := range m.args {
		switch arg := arg.(type) {
		case Set[string]:
			args[i] = arg.MustGet(lang)
		case Set[Message]:
			args[i] = arg.MustGet(lang)
		default:
			args[i] = arg
		}
	}
	return fmt.Sprintf(m.format, args...)
}

type Set[M string | Message] map[any]M

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

func eval(msg any, language any) (string, bool) {
	switch m := msg.(type) {
	case Message:
		return m.eval(language)
	case string:
		return m, true
	default:
		panic("unreachable")
	}
}

func mustEval(msg any, language any) string {
	switch m := msg.(type) {
	case Message:
		return m.mustEval(language)
	case string:
		return m
	default:
		panic("unreachable")
	}
}
