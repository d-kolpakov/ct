package pipelines

import (
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
)

func (p *Handler) IsReady(c commit.GitCommit) (commit.GitCommit, error) {
	q := "Is ready:"
	vars := []string{"yes", "no"}

	ir, err := p.In.Select(q, vars, &input.Options{
		Default:  "yes",
		Loop:     true,
		Required: true,
	})

	if err != nil {
		return c, err
	}

	c.CommitReady = ir == "yes"

	return c, nil
}
