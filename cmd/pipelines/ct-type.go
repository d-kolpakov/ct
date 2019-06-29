package pipelines

import (
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
)

func (p *Handler) CtType(c commit.GitCommit) (commit.GitCommit, error) {
	q := "Choose a commit type:"
	vars := []string{"mistake", "fix", "feature"}

	cType, err := p.In.Select(q, vars, &input.Options{
		Default:  "mistake",
		Loop:	  true,
		Required: true,
	})

	if err != nil {
		return c, err
	}

	c.CtType = cType

	return c, nil
}
