package cmd

import (
	"github.com/therehello/cf-tool/client"
)

// Pull command
func Pull() (err error) {
	cln := client.Instance
	info := Args.Info
	ac := Args.Accepted
	contestPath := info.Path()
	if err != nil {
		return
	}
	if err = cln.Pull(info, contestPath, ac); err != nil {
		if err = loginAgain(cln, err); err == nil {
			err = cln.Pull(info, contestPath, ac)
		}
	}
	return
}
