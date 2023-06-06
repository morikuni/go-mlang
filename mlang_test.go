package mlang_test

import (
	"fmt"
	"testing"

	"golang.org/x/text/language"

	"github.com/morikuni/go-mlang"
)

func TestDict_String_MustGet(t *testing.T) {
	hello := func(name string) mlang.Message {
		return mlang.Dict[string]{
			language.English: fmt.Sprintf("Hello, %s!", name),
			language.French:  fmt.Sprintf("Bonjour, %s!", name),
		}
	}

	equal(t, hello("Alice").MustGet(language.English), "Hello, Alice!")
	equal(t, hello("Alice").MustGet(language.French), "Bonjour, Alice!")
	oneOf(t, hello("Alice").MustGet(language.Japanese), "Hello, Alice!", "Bonjour, Alice!")
}

func TestDict_DynamicMessage_MustGet(t *testing.T) {
	apple := func() mlang.Message {
		return mlang.Dict[string]{
			language.English: "apple",
			language.French:  "pomme",
		}
	}
	hello := func() mlang.Message {
		return mlang.Dict[mlang.Template]{
			language.English: mlang.NewTemplate("Hello, %s!", apple()),
			language.French:  mlang.NewTemplate("Bonjour, %s!", apple()),
		}
	}

	equal(t, hello().MustGet(language.English), "Hello, apple!")
	equal(t, hello().MustGet(language.French), "Bonjour, pomme!")
	oneOf(t, hello().MustGet(language.Japanese), "Hello, apple!", "Bonjour, pomme!")
}

func TestSetDefaultLanguage(t *testing.T) {
	hello := func(name string) mlang.Message {
		return mlang.Dict[string]{
			language.English: fmt.Sprintf("Hello, %s!", name),
			language.French:  fmt.Sprintf("Bonjour, %s!", name),
		}
	}

	equal(t, hello("Alice").MustGet(language.English), "Hello, Alice!")
	equal(t, hello("Alice").MustGet(language.French), "Bonjour, Alice!")

	mlang.SetDefaultLanguage(language.French)
	equal(t, hello("Alice").MustGet(language.Japanese), "Bonjour, Alice!")
	mlang.SetDefaultLanguage(language.English)
	equal(t, hello("Alice").MustGet(language.Japanese), "Hello, Alice!")
}

func equal(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func oneOf(t *testing.T, got any, want ...any) {
	t.Helper()
	for _, w := range want {
		if got == w {
			return
		}
	}
	t.Errorf("got %v, want one of %v", got, want)
}
