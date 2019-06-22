package main

import (
	"flag"
	"log"
	"time"
	"net/http"
	"net/url"

	"github.com/atotto/clipboard"
)

var (
	pushURL string
)

func init() {
	flag.StringVar(&pushURL, "u", "", "set url")
	flag.Parse()
}

func main() {
	c := time.Tick(time.Second * 2)
	
	last := ""

	pu, err := url.Parse(pushURL)

	if err != nil {
		log.Fatalf("Can not set URL with error: %s", err.Error())
	}

	for {
		<- c

		content, err := clipboard.ReadAll()
		if err != nil || content == last {
			continue
		}
		last = content
		log.Println(content)
		go request(*pu, content)
	}

}

func request(u url.URL, text string) error {
	q := u.Query()
	q.Set("copy", text)
	u.RawQuery = q.Encode()
	_, err := http.Get(u.String())
	return err
}