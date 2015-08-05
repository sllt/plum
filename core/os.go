package core

import (
	"errors"
	"fmt"
	. "github.com/sllt/plum/types"
	"os"
	"os/exec"
	"strings"
)

func Exec(a []PlumType) (PlumType, error) {
	cmd := a[0]
	switch tcmd := cmd.(type) {
	case string:
		any := strings.Split(tcmd, " ")
		o, _ := exec.Command(any[0], any[1:]...).Output()
		fmt.Println(string(o))
		return nil, nil
	default:
		return nil, errors.New("doesn't support type ")
	}
}

func Exit(a []PlumType) (PlumType, error) {
	os.Exit(0)
	return nil, nil
}
