package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fatih/color"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func flagInit() (int, string) {
	color.Set(color.FgGreen, color.Bold)
	defer color.Unset()
	prID := flag.Int("n", -1, color.HiRedString("number of pull request you want to fetch for review - mandatory"))
	branch := flag.String("b", "review", color.HiBlueString("name of branch you want PR fetched to"))
	flag.Parse()
	return *prID, *branch
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	number, branch := flagInit()
	if number < 0 {
		log.Fatal(fmt.Errorf(color.RedString(`use "-n" flag to set number of pull request you want to fetch`)))

	}
	r, err := git.PlainOpen(".")
	checkError(err)

	externalRefs := config.RefSpec(fmt.Sprintf("refs/pull/%d/head:refs/heads/%s", number, branch))
	err = r.Fetch(&git.FetchOptions{Progress: os.Stdout, RemoteName: "upstream", RefSpecs: []config.RefSpec{externalRefs}})
	checkError(err)

	w, err := r.Worktree()
	checkError(err)

	branchAsPlmbRef := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch))
	err = w.Checkout(&git.CheckoutOptions{Branch: branchAsPlmbRef})

	checkError(err)
	color.HiGreen("Done!")
}
