package pipelines

import (
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
)

func (p *Handler) IncrementVersion(c commit.GitCommit) (commit.GitCommit, error) {
	q := "Need to increment version:"
	vars := []string{"yes", "no", "force"}

	iv, err := p.In.Select(q, vars, &input.Options{
		Default:  "yes",
		Loop:	  true,
		Required: true,
	})

	if err != nil {
		return c, err
	}

	if iv == "force" {
		c, err = p.ForceVersion(c)

		if err != nil {
			return c, err
		}
	}

	c.IncrementVersion = iv == "yes"

	return c, nil
}
