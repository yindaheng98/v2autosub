package main

import (
	"flag"
	"fmt"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"iochen.com/v2gen/v2"
	"log"
	"os"
	"path/filepath"
)

var (
	FlagAddr        = flag.String("u", "", "subscription address(URL)")
	FlagOut         = flag.String("o", "/etc/v2ray/config.json", "output path")
	FlagConf        = flag.String("config", "/etc/v2ray/v2gen.ini", "v2gen config path")
	FlagLinkTmpl    = flag.String("node_template", "", "V2Ray template path for single nodes")
	FlagLinksTmpl   = flag.String("nodes_template", "", "V2Ray template path for all nodes")
	FlagLinkNameFmt = flag.String("node_tag_fmt", "v2gen_%d", "format of the tag for nodes")
	FlagMax         = flag.Int("max", 8, "max number of nodes")
	FlagPing        = flag.Bool("ping", true, "ping nodes")
	FlagDest        = flag.String("dst", "https://cloudflare.com/cdn-cgi/trace", "test destination url (vmess ping only)")
	FlagCount       = flag.Int("c", 3, "ping count for each node")
	FlagThreads     = flag.Int("thread", 3, "threads used when pinging")
)

func main() {
	flag.Parse()
	var links []v2gen.Link // combine links from different sources

	// read from subscribe address(net)
	if *FlagAddr != "" {
		log.Printf("Reading from %s...\n", *FlagAddr)
		var err error
		links, err = Subscribe(*FlagAddr)
		if err != nil {
			log.Fatal(err)
		}
	}

	// if no Link, then exit
	if len(links) == 0 {
		log.Println("no available links, nothing to do")
		os.Exit(0)
	}

	// ping the nodes
	if *FlagPing { // if ping
		// make ping info list
		piList := PingAndSort(links, *FlagCount, *FlagDest, *FlagThreads)
		links = []v2gen.Link{}
		for i, pi := range piList {
			if i >= *FlagMax || pi.Err != nil {
				break
			}
			links = append(links, pi.Link)
		}
	}

	// get the nodes
	if len(links) > *FlagMax {
		links = links[0:*FlagMax]
	}

	// generate config
	lsc := LinksConf{linkConfs: []LinkConf{}, tmplPath: *FlagLinksTmpl, confPath: *FlagConf}
	for i, link := range links {
		lc := LinkConf{link: link, name: fmt.Sprintf(*FlagLinkNameFmt, i), tmplPath: *FlagLinkTmpl, confPath: *FlagConf}
		lsc.linkConfs = append(lsc.linkConfs, lc)
	}

	// write config
	if *FlagOut == "-" || *FlagOut == "" {
		fmt.Println(lsc.String())
		return
	} else {
		err := ioutil.WriteFile(*FlagOut, []byte(lsc.String()), 0644)
		if err != nil {
			logrus.Fatal(err)
		} else {
			log.Printf("config has been written to %s\n", filepath.Clean(*FlagOut))
		}
	}
}
