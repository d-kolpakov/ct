package commit

import (
	"fmt"
	"os/exec"
	"regexp"
)

func (c *GitCommit) ParseCurrentState() bool{
	out, err := exec.Command("git", "status").Output()

	if err != nil {
		return false
	}
	status := string(out)

	re := regexp.MustCompile(`(?s)(Changes to be committed:\n)(.*?)\n\n(.*?)\n\n`)
	match := re.FindStringSubmatch(status)
	if len(match) <= 0 {
		return false
	}
	cs := match[0]

	re = regexp.MustCompile(`(?m)(deleted|new file|modified|renamed).*`)
	match = re.FindStringSubmatch(cs)
	fmt.Println(match)

	return false
}