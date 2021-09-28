package main

import (
	"io/ioutil"
	"iochen.com/v2gen/v2"
	"iochen.com/v2gen/v2/infra"
)

func BasicConf(tmplPath, confPath string, defaultTmpl []byte, exConf map[string]string) string {
	tmpl := defaultTmpl
	if tmplPath != "" {
		t, err := ioutil.ReadFile(tmplPath)
		if err == nil {
			tmpl = t
		}
	}
	v2genConf := infra.V2genConfig{}
	confFile, err := ioutil.ReadFile(confPath)
	if err == nil {
		v2genConf = infra.ParseV2genConf(confFile)
	}
	conf := infra.DefaultConf()
	bytes, err := infra.GenV2RayConf(*conf.Append(v2genConf).Append(exConf), tmpl)
	if err != nil {
		panic(err)
	}
	v2rayConf := string(bytes)
	return v2rayConf

}

type LinkConf struct {
	name     string
	link     v2gen.Link
	tmplPath string
	confPath string
}

func (lc *LinkConf) String() string {
	replace := lc.link.Config()
	replace["name"] = lc.name
	v2rayConf := BasicConf(lc.tmplPath, lc.confPath, []byte(DefaultLinkConfTmpl), replace)
	return v2rayConf
}

type LinksConf struct {
	linkConfs []LinkConf
	tmplPath  string
	confPath  string
}

func (lsc *LinksConf) String() string {
	tags := ""
	outbounds := ""
	for _, lc := range lsc.linkConfs {
		tags += ",\"" + lc.name + "\""
		outbounds += "," + lc.String()
	}
	tags = tags[1:]
	outbounds = outbounds[1:]

	replace := map[string]string{"outbounds": outbounds, "tags": tags}
	v2rayConf := BasicConf(lsc.tmplPath, lsc.confPath, []byte(DefaultLinksConfTmpl), replace)
	return v2rayConf
}
