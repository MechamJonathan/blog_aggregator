package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	ctx := context.Background()
	err := s.db.ResetUsers(ctx)
	if err != nil {
		return fmt.Errorf("couldn't reset users table: %w", err)
	}
	fmt.Println("Users table reset successfully!")

	return nil
}
