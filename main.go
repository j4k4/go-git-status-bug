package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/go-git/go-git/v5"
)

func main() {
	var err error
	var repo *git.Repository
	var worktree *git.Worktree
	var status git.Status

	if repo, err = git.PlainOpen("."); err != nil {
		log.Fatal(err)
	}

	if worktree, err = repo.Worktree(); err != nil {
		log.Fatal(err)
	}

	err = os.Remove("this-should-be-ignored/ignored-too.md")
	if err = os.Remove("this-should-be-ignored/ignored-too.md"); err != nil && !errors.Is(err, syscall.ENOENT) {
		log.Fatal(err)
	}

	if status, err = worktree.Status(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("working tree is clean: %v\n", status.IsClean())
	fmt.Printf("creating file %q which should be ignored\n", "this-should-be-ignored/ignored-too.md")

	if err = os.WriteFile("this-should-be-ignored/ignored-too.md", []byte("test"), os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if status, err = worktree.Status(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("working tree is clean: %v\n", status.IsClean())
}
