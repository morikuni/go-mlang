package mlangfailure

import (
	"github.com/morikuni/failure/v2"

	"github.com/morikuni/go-mlang"
	"github.com/morikuni/go-mlang/internal"
)

const (
	Key = internal.FailureKey
)

func MessageOf(err error) mlang.Message {
	m, _ := failure.ValueAs[mlang.Message](err, Key)
	return m
}
