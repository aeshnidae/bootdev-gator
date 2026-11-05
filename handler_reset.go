package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.Reset(ctx)
	if err != nil {
		return fmt.Errorf("could not reset the database : %w", err)
	}
	fmt.Println("Reset complete")
	return nil
}
