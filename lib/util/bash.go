package util

import (
	"fmt"
	"github.com/autom8ter/util"
	"os"
	"os/exec"
)

type CmdStat struct {
	Path   string   `json:"path"`
	Dir    string   `json:"dir"`
	Args   []string `json:"args"`
	Output []byte   `json:"output"`
	Err    error    `json:"err"`
}

func (c *CmdStat) Error() string {
	return fmt.Sprintln(util.ToPrettyJsonString(c))
}

type Script string

func NewScript(s string) Script {
	return Script(s)
}

func (s Script) Bash(script string) ([]byte, error) {
	cmd := exec.Command("bash", "-c", script)
	cmd.Dir = "."
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	if err != nil {
		newErr := &CmdStat{
			Path:   cmd.Path,
			Dir:    cmd.Dir,
			Args:   cmd.Args,
			Output: out,
			Err:    err,
		}
		return nil, newErr
	}
	return out, nil
}
