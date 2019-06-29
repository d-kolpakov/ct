package commit

import "sync"

type GitCommit struct {
	Msg string
	CtType string
	Version string
	IncrementVersion bool
	CommitReady bool
	MergeInto string
	NeedPush bool
}

func (c *GitCommit) Commit(wg *sync.WaitGroup) error {
	defer wg.Done()

	return nil
}

func formCommitMsg() string {
	var msg string

	return msg
}