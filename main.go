package main

import (
	"fmt"
	"log"

	"github.com/aeshnidae/bootdev-gator/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name      string
	arguments []string
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	err = cfg.SetUser("aeshnidae")
	if err != nil {
		log.Fatalf("couldn't set current user: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Printf("Read config again: %+v", cfg)

}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("missing username argument")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %v", err)
	}

}
