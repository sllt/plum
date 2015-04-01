package core

import (
	"errors"
	"fmt"
	"os/exec"
	. "plum/types"
)

func Exec(a []PlumType) (PlumType, error) {
	cmd := a[0]
	switch tcmd := cmd.(type) {
	case string:
		o, _ := exec.Command(tcmd).Output()
		fmt.Println(string(o))
		return nil, nil
	default:
		return nil, errors.New("doesn't support type ")
	}
}
