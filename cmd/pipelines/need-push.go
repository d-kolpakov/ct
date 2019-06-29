package pipelines

import (
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
)

func (p *Handler) NeedPush(c commit.GitCommit) (commit.GitCommit, error) {
	q := "Need push:"
	vars := []string{"yes", "no"}

	np, err := p.In.Select(q, vars, &input.Options{
		Default:  "yes",
		Loop:     true,
		Required: true,
	})

	if err != nil {
		return c, err
	}

	c.NeedPush = np == "yes"

	return c, nil
}
