package prov

import (
	"io/ioutil"
	"os"
	"os/exec"
)

func init() {
	RegisterRunner("hostname", Hostname)
}

func Hostname(dir string, vars Vars, args Args, run bool) (Status, error) {
	hostname, ok := args.String("hostname")
	if !ok {
		return OK, ErrInvalidArg("hostname")
	}
	current, err := os.Hostname()
	if err != nil {
		return OK, err
	}
	if current == hostname {
		return OK, nil
	}
	if run {
		output, err := exec.Command("hostname", hostname).CombinedOutput()
		if err != nil {
			return OK, ErrCommandFailed{err, output}
		}
		err = ioutil.WriteFile("/etc/hostname", []byte(hostname), 0644)
		if err != nil {
			return OK, err
		}
	}
	return Changed, nil
}
