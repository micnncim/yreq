package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/micnncim/yreq"
)

func main() {
	var (
		f   = flag.String("f", "", "request defined yaml file")
		url = flag.String("url", "", "url getting request")
	)
	flag.Parse()

	buf, err := ioutil.ReadFile(*f)
	if err != nil {
		log.Fatal(err)
	}

	reqs, err := yreq.FromYaml(buf)
	if err != nil {
		log.Fatal(err)
	}

	resps := make([]interface{}, 0, len(reqs))
	for _, req := range reqs {
		resp, err := yreq.SendReq(*url, req)
		if err != nil {
			log.Fatal(err)
		}
		resps = append(resps, resp)
	}

	bytes, err := yreq.ToYamls(resps)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
