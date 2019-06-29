package commit

import (
	"sync"
)

type GitCommit struct {
	currentState state
	Msg string
	CtType string
	Version string
	IncrementVersion bool
	CommitReady bool
	MergeInto string
	NeedPush bool
}

type state struct {
	branch string
}

func (c *GitCommit) NeedToCommit() bool{
	c.ParseCurrentState()

	return false
}

func (c *GitCommit) Commit(wg *sync.WaitGroup) error {
	defer wg.Done()

	formCommitMsg()

	return nil
}

func formCommitMsg() string {
	var msg string

	return msg
}