package main

import (
	"flag"
	"fmt"
	"log"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

func flagInit() (int, string, bool) {
	prID := flag.Int("n", -1, "number of pull request you want to fetch for review - mandatory")
	branch := flag.String("b", "review", "name of branch you want PR fetched to")
	overwrite := flag.Bool("o", false, "whether to overwrite exiting branch with same name")

	flag.Parse()
	return *prID, *branch, *overwrite
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//git fetch origin pull/ID/head:BRANCHNAME && git checkout branchname
func main() {
	number, branch, _ := flagInit()
	if number < 0 {
		log.Fatal(fmt.Errorf("wrong PR number"))

	}
	r, err := git.PlainOpen(".")
	checkError(err)

	externalRefs := config.RefSpec(fmt.Sprintf("refs/pull/%d/head:refs/heads/%s", number, branch))
	err = r.Fetch(&git.FetchOptions{RemoteName: "upstream", RefSpecs: []config.RefSpec{externalRefs}})
	checkError(err)

	w, err := r.Worktree()
	checkError(err)

	branchAsPlmbRef := plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch))
	err = w.Checkout(&git.CheckoutOptions{Branch: branchAsPlmbRef})

	checkError(err)
}
