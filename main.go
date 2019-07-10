package main

import (
	"flag"
	"log"
	"time"
	"net/http"
	"net/url"
	"io/ioutil"
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/kardianos/service"
)

var (
	token string
	asService string
)

type Program struct{}

func init() {
	flag.StringVar(&token, "t", "", "set token")
	flag.StringVar(&asService, "s", "run", "service commands [install|uninstall|start|run]")
	flag.Parse()
}

func main() {

	if asService != "" {
		p := &Program{}
	  serviceConfig := &service.Config{
			Name:        "bark-helper-go",
			DisplayName: "Bark Helper",
			Description: "Brak Helper in Go",
			Arguments: []string{"-t", token},
		}
		s, err := service.New(p, serviceConfig)
		if err != nil {
			log.Fatal(err)
		}

		switch asService {
		case "install":
			err = s.Install()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("service install successed")
		case "uninstall":
			err = s.Uninstall()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("service uninstall successed")
		case "start":
			err = s.Start()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("service start successed")
		case "stop":
			err = s.Stop()
			if err != nil {
				log.Fatal(err)
			}
			log.Println("service stop successed")
		case "run":
			s.Run()
		default:
			log.Fatalf("command not legal")
		}

	}

}

func request(u url.URL, text string) error {

	if (len([]byte(text)) > 3900) {
		log.Printf("Cliptext too large to push with length: %d", len([]byte(text)))
		return fmt.Errorf("Cliptext too large to push with length: %d", len([]byte(text)))
	}

	resp, err := http.PostForm(u.String(), url.Values{"copy": {text}})
	if err != nil {
		log.Printf("Request(%s) with error: %s", u.String(),err.Error())
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	log.Printf("Response body: %s", body)
	
	return nil
}


func (p *Program) Start(s service.Service) error {
	go p.run()
	return nil
}

func (p *Program) Stop(s service.Service) error {
	return nil
}

func (p *Program) run() {
	c := time.Tick(time.Second * 2)
	
	last := ""

	pu, err := url.Parse(fmt.Sprintf("https://api.day.app/%s/bark-helper-go?automaticallyCopy=1&copy=optional", token))

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
		go request(*pu, content)
	}
}