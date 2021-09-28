package main

import (
	"io"
	"io/ioutil"
	"iochen.com/v2gen/v2"
	"iochen.com/v2gen/v2/common/base64"
	"iochen.com/v2gen/v2/vmess"
	"net/http"
)

func Subscribe(addr string) ([]v2gen.Link, error) {
	resp, err := http.Get(addr)
	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	str, err := base64.Decode(string(bytes))
	if err != nil {
		return nil, err
	}
	linkList, err := vmess.Parse(str)
	if err != nil {
		return nil, err
	}
	links := make([]v2gen.Link, len(linkList))
	for i := range linkList {
		links[i] = linkList[i]
	}
	return links, nil
}
