package updater

import (
	"context"
	"fmt"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

const (
	repoURL = "https://github.com/delphinus/homebrew-macvim-kaoriya"
)

// Updater updates cask for macvim-kaoriya
func Updater(ctx context.Context) error {
	_, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: repoURL,
	})
	if err != nil {
		return fmt.Errorf("error in Clone: %v", err)
	}
	return nil
}
