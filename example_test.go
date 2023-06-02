package mlang_test

import (
	"fmt"

	"golang.org/x/text/language"

	"github.com/morikuni/go-mlang"
)

var InvalidUserName = mlang.Set[string]{
	language.English:  "Invalid user name",
	language.Japanese: "ユーザ名が不正です",
}

func InvalidPasswordLength(min, max int) mlang.Set[string] {
	return mlang.Set[string]{
		language.English:  fmt.Sprintf("Password must be between %d and %d characters long", min, max),
		language.Japanese: fmt.Sprintf("パスワードは%d文字以上%d文字以下である必要があります", min, max),
	}
}

func PenCount(count int) mlang.Set[string] {
	s := mlang.Set[string]{
		language.Japanese: fmt.Sprintf("%d本のペン", count),
	}

	if count == 1 {
		s[language.English] = "a pen"
	} else {
		s[language.English] = fmt.Sprintf("%d pens", count)
	}

	return s
}

func IHavePen(count int) mlang.Set[mlang.Template] {
	return mlang.Set[mlang.Template]{
		language.English:  mlang.NewTemplate("I have %s", PenCount(count)),
		language.Japanese: mlang.NewTemplate("私は%sを持っています", PenCount(count)),
	}
}

func ExampleSet() {
	fmt.Println(InvalidUserName.MustGet(language.English))
	fmt.Println(InvalidUserName.MustGet(language.Japanese))

	fmt.Println("-----")

	fmt.Println(InvalidPasswordLength(1, 2).MustGet(language.English))
	fmt.Println(InvalidPasswordLength(1, 2).MustGet(language.Japanese))

	fmt.Println("-----")

	fmt.Println(IHavePen(1).MustGet(language.English))
	fmt.Println(IHavePen(1).MustGet(language.Japanese))
	fmt.Println(IHavePen(2).MustGet(language.English))
	fmt.Println(IHavePen(2).MustGet(language.Japanese))

	// Output:
	// Invalid user name
	// ユーザ名が不正です
	// -----
	// Password must be between 1 and 2 characters long
	// パスワードは1文字以上2文字以下である必要があります
	// -----
	// I have a pen
	// 私は1本のペンを持っています
	// I have 2 pens
	// 私は2本のペンを持っています
}
