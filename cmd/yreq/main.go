package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"golang.org/x/sync/errgroup"

	"github.com/micnncim/yreq"
)

func main() {
	var (
		req = flag.String("req", "", "request defined yaml file")
		url = flag.String("url", "", "url getting request")
	)
	flag.Parse()

	buf, err := ioutil.ReadFile(*req)
	if err != nil {
		log.Fatal(err)
	}

	reqs, err := yreq.FromYaml(buf)
	if err != nil {
		log.Fatal(err)
	}

	eg := errgroup.Group{}
	for _, req := range reqs {
		req := req
		eg.Go(func() error {
			respBody, err := yreq.SendReq(*url, req)
			if err != nil {
				return err
			}
			fmt.Println(string(respBody))
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
