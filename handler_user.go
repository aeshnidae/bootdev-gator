package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aeshnidae/bootdev-gator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage : %s <name>", cmd.Name)
	}

	name := cmd.Args[0]
	ctx := context.Background()

	_, err := s.db.GetUser(ctx, name)
	if err != nil {
		return fmt.Errorf("username : %s does not exist", name)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage : %v <name>", cmd.Name)
	}

	ctx := context.Background()

	name := cmd.Args[0]

	_, err := s.db.GetUser(ctx, name)
	if err == nil {
		return fmt.Errorf("username : %v already exists", name)
	}

	insertedUser, err := s.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return fmt.Errorf("error registering user : %v ,%w", name, err)
	}

	err = s.cfg.SetUser(insertedUser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	fmt.Println("User registered successfully!")
	fmt.Println(insertedUser)
	return nil
}
