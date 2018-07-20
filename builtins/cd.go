package builtins

import (
	"errors"
	"fmt"
	"os"
	"os/user"
	"strings"
)

// ErrTooManyArgs is returned when 'cd' was called with too many arguments
var ErrTooManyArgs = errors.New("Too many arguments passed to cd")

func goHome(usr *user.User) error {
	err := os.Chdir(usr.HomeDir)

	if err != nil {
		return err
	}

	return nil
}

func getUser() (*user.User, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	return usr, nil
}

func goUserHome() error {
	usr, err := getUser()
	if err != nil {
		return err
	}

	errr := goHome(usr)
	if errr != nil {
		return err
	}

	return nil
}

// Cd changes the user directory, handling special cases with ~ replacement to user's home directory
func Cd(args []string) error {
	if len(args) > 2 {
		return ErrTooManyArgs
	}

	// `cd` to home dir with empty path
	if len(args) < 2 || args[1] == "~" {
		err := goUserHome()
		if err != nil {
			return err
		}

		return nil
	}

	if args[1] == "..." || args[1] == "...." {
		fmt.Println("cd using more than 2 dots as parent directories is not implemented yet")
	}

	// replace ~ with user home directory
	pathArgs := strings.Split(args[1], string(os.PathSeparator))

	if pathArgs[0] == "~" {
		usr, err := getUser()
		if err != nil {
			return err
		}

		pathArgs[0] = usr.HomeDir
	}

	path := strings.Join(pathArgs, string(os.PathSeparator))

	err := os.Chdir(path)
	if err != nil {
		return err
	}

	return nil
}
