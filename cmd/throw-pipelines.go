package cmd

import (
	"fmt"
	"github.com/d-kolpakov/ct/cmd/pipelines"
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
	"os"
	"sync"
)


func GoThrow(wg *sync.WaitGroup) (*commit.GitCommit,error) {
	defer wg.Done()
	defer fmt.Println("Config done")

	fmt.Println("Config:")
	c := commit.GitCommit{}
	in := input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
	pipe := &pipelines.Handler{
		In: in,
	}

	c, err := pipe.GetMsg(c)

	if err != nil {
		return nil, err
	}

	c, err = pipe.CtType(c)

	if err != nil {
		return nil, err
	}

	c, err = pipe.IncrementVersion(c)

	if err != nil {
		return nil, err
	}

	c, err = pipe.IsReady(c)

	if err != nil {
		return nil, err
	}

	c, err = pipe.MergeInto(c)

	if err != nil {
		return nil, err
	}

	c, err = pipe.NeedPush(c)

	if err != nil {
		return nil, err
	}

	return &c, nil
}
