package main

import (
	"fmt"
	"net/http"
)

type RequestBuilder struct {
	method  string
	url     string
	headers map[string]string
	body    string
}

func NewRequestBuilder() *RequestBuilder {
	return &RequestBuilder{
		headers: make(map[string]string),
	}
}

func (b *RequestBuilder) SetMethod(m string) *RequestBuilder {
	b.method = m
	return b
}

func (b *RequestBuilder) SetURL(url string) *RequestBuilder {
	b.url = url
	return b
}

func (b *RequestBuilder) AddHeader(key, val string) *RequestBuilder {
	b.headers[key] = val
	return b
}

func (b *RequestBuilder) SetBody(body string) *RequestBuilder {
	b.body = body
	return b
}

func (b *RequestBuilder) Build() (*http.Request, error) {
	req, err := http.NewRequest(b.method, b.url, nil)
	if err != nil {
		return nil, fmt.Errorf("NewRequest failed: %v", err)
	}

	for key, val := range b.headers {
		req.Header.Add(key, val)
	}

	return req, nil
}
