package pipelines

import (
	"errors"
	"fmt"
	"github.com/d-kolpakov/ct/commit"
	"github.com/tcnksm/go-input"
	"regexp"
)

var reg = regexp.MustCompile(`^\d{1,4}\.\d{1,4}\.\d{1,4}$`)

func (p *Handler) ForceVersion(c commit.GitCommit) (commit.GitCommit, error) {
	q := "Version (ex: 1.0.3):"

	version, err := p.In.Ask(q, &input.Options{
		Loop:	  true,
		Required: true,
		ValidateFunc: func(v string) error {
			if !reg.MatchString(v) {
				return errors.New(fmt.Sprintf("Wrong version %s. Example: 1.0.0", v))
			}

			return nil
		},
	})

	if err != nil {
		return c, err
	}

	c.Version = version

	return c, nil
}
