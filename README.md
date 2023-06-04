# go-mlang

[![Go Reference](https://pkg.go.dev/badge/github.com/morikuni/go-mlang.svg)](https://pkg.go.dev/github.com/morikuni/go-mlang)

Go library for i18n (multi language message).

## Usage

```go
package main

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/morikuni/go-mlang"
)

var InvalidUserName = mlang.Dict[string]{
	language.English:  "Invalid user name",
	language.Japanese: "ユーザ名が不正です",
}

func InvalidPasswordLength(min, max int) mlang.Message {
	return mlang.Dict[string]{
		language.English:  fmt.Sprintf("Password must be between %d and %d characters long", min, max),
		language.Japanese: fmt.Sprintf("パスワードは%d文字以上%d文字以下である必要があります", min, max),
	}
}

func PenCount(count int) mlang.Message {
	s := mlang.Dict[string]{
		language.Japanese: fmt.Sprintf("%d本のペン", count),
	}

	if count == 1 {
		s[language.English] = "a pen"
	} else {
		s[language.English] = fmt.Sprintf("%d pens", count)
	}

	return s
}

func IHavePen(count int) mlang.Message {
	return mlang.Dict[mlang.Template]{
		language.English:  mlang.NewTemplate("I have %s", PenCount(count)),
		language.Japanese: mlang.NewTemplate("私は%sを持っています", PenCount(count)),
	}
}

func ExampleSet() {
	fmt.Println(InvalidUserName.MustGet(language.English))
	// Invalid user name
	fmt.Println(InvalidUserName.MustGet(language.Japanese))
	// ユーザ名が不正です

	fmt.Println(InvalidPasswordLength(1, 2).MustGet(language.English))
	// Password must be between 1 and 2 characters long
	fmt.Println(InvalidPasswordLength(1, 2).MustGet(language.Japanese))
	// パスワードは1文字以上2文字以下である必要があります

	fmt.Println(IHavePen(1).MustGet(language.English))
	// I have a pen
	fmt.Println(IHavePen(1).MustGet(language.Japanese))
	// 私は1本のペンを持っています
	fmt.Println(IHavePen(2).MustGet(language.English))
	// I have 2 pens
	fmt.Println(IHavePen(2).MustGet(language.Japanese))
	// 私は2本のペンを持っています
}
```
