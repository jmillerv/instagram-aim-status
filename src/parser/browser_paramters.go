package parser

import (
	"net/http"
	"time"
)

type parameters struct {
	Header        http.Header
	Method        string
	Url           string
	UsePhantomJS  bool
	PostBody      string
	RedirectTimes int
	DialTimeout   time.Duration
	ConnTimeout   time.Duration
	RetryPause    time.Duration
	TryTimes      int
}

func (p parameters) GetMethod() string {
	return p.Method
}

func (p parameters) GetUrl() string {
	return p.Method
}

func (p parameters) GetHeader() http.Header {
	return p.Header
}

func (p parameters) GetPostBody() string {
	return p.PostBody
}

func (p parameters) GetRedirectTimes() int {
	return p.RedirectTimes
}

func (p parameters) GetDialTimeout() time.Duration {
	return p.DialTimeout
}

func (p parameters) GetConnTimeout() time.Duration {
	return p.ConnTimeout
}

func (p parameters) GetRetryPause() time.Duration {
	return p.RetryPause
}

func (p parameters) GetTryTimes() int {
	return p.TryTimes
}

func (p parameters) GetusePhantomJS() bool {
	return p.UsePhantomJS
}
