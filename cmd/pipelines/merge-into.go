package pipelines

import (
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
)

func (p *Handler) MergeInto(c commit.GitCommit) (commit.GitCommit, error) {
	q := "Merge into:"
	vars := []string{"develop", "release", "none"}

	branch, err := p.In.Select(q, vars, &input.Options{
		Default:  "develop",
		Loop:	  true,
		Required: true,
	})

	if err != nil {
		return c, err
	}

	c.MergeInto = branch

	return c, nil
}
