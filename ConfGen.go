package main

import (
	"io/ioutil"
	"iochen.com/v2gen/v2"
	"iochen.com/v2gen/v2/infra"
	"strings"
)

type LinkConf struct {
	name     string
	link     v2gen.Link
	tmplPath string
	confPath string
}

func (lc *LinkConf) String() string {
	tmpl := []byte(DefaultLinkConfTmpl)
	if lc.tmplPath != "" {
		t, err := ioutil.ReadFile(lc.tmplPath)
		if err == nil {
			tmpl = t
		}
	}
	v2genConf := infra.V2genConfig{}
	confFile, err := ioutil.ReadFile(lc.confPath)
	if err == nil {
		v2genConf = infra.ParseV2genConf(confFile)
	}
	conf := infra.DefaultConf()
	bytes, err := infra.GenV2RayConf(*conf.Append(v2genConf).Append(lc.link.Config()), tmpl)
	if err != nil {
		panic(err)
	}
	v2rayConf := string(bytes)
	v2rayConf = strings.ReplaceAll(v2rayConf, "{{name}}", lc.name)
	return v2rayConf
}
