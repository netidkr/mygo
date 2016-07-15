package user

import (
	"os/user"
)

func Username() string {
	r, err := user.Current()
	if err != nil {
		return ""
	}
	
	return r.Username
}
