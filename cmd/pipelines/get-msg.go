package pipelines

import (
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
)

func (p *Handler) GetMsg(c commit.GitCommit) (commit.GitCommit, error) {
	q := "Commit msg:"

	version, err := p.In.Ask(q, &input.Options{})

	if err != nil {
		return c, err
	}

	c.Msg = version

	return c, nil
}
