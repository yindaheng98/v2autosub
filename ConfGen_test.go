package main

import (
	"fmt"
	"testing"
)

func TestLinkConf_String(t *testing.T) {
	links, err := Subscribe(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(links)
	lc := LinkConf{link: links[0], name: "test0"}
	t.Log(lc.String())
}

func TestLinksConf_String(t *testing.T) {
	links, err := Subscribe(addr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(links)
	lcs := LinksConf{linkConfs: []LinkConf{}}
	for i, link := range links {
		lcs.linkConfs = append(lcs.linkConfs, LinkConf{link: link, name: fmt.Sprintf("test%d", i)})
	}
	t.Log(lcs.String())
}
