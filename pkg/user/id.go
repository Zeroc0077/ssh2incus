package user

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

var idExe = "" // path to the 'id' program.

func init() {
	if path, err := exec.LookPath("/usr/bin/id"); err == nil {
		idExe = path
	} else {
		log.Fatal(err)
	}
}

func idGroupList(u *User) ([]string, error) {
	data, err := command(idExe, "-G", u.Username)
	if err != nil {
		return nil, err
	}

	data = bytes.TrimSpace(data)
	if len(data) > 0 {
		return strings.Fields(string(data)), nil
	}

	return nil, errNotFound
}

func idGroupNameList(u *User) ([]string, error) {
	data, err := command(idExe, "-Gn", u.Username)
	if err != nil {
		return nil, err
	}

	data = bytes.TrimSpace(data)
	if len(data) > 0 {
		return strings.Fields(string(data)), nil
	}

	return nil, errNotFound
}
