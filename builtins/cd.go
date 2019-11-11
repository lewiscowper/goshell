package builtins

import (
	"errors"
	"fmt"
	"github.com/lewiscowper/shell/helpers"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// ErrTooManyArgs is returned when 'cd' was called with too many arguments
var ErrTooManyArgs = errors.New("Too many arguments passed to cd")

func getUser() (*user.User, error) {
	usr, err := user.Current()
	if err != nil {
		return nil, err
	}

	return usr, nil
}

// Cd changes the user directory, handling special cases with ~ replacement to user's home directory
func Cd(ctx *helpers.ShellContext, args []string) error {
	if len(args) > 2 {
		return ErrTooManyArgs
	}

	targetPath := args[0]

	if targetPath == "..." || targetPath == "...." {
		fmt.Println("cd using more than 2 dots as parent directories is not implemented yet")
	}

	pathArgs := strings.Split(targetPath, string(os.PathSeparator))

	switch pathArgs[0] {
	case "~":
		usr, err := getUser()
		if err != nil {
			return err
		}

		pathArgs[0] = usr.HomeDir
	case "-":
		if len(ctx.DirStack) == 1 {
			usr, err := getUser()
			if err != nil {
				return err
			}

			pathArgs[0] = usr.HomeDir
		}
		pathArgs[0] = ctx.DirStack[len(ctx.DirStack)-1]
	}

	path := strings.Join(pathArgs, string(os.PathSeparator))

	finalPath, err := filepath.Abs(path)
	if err != nil {
		return err
	}

	err = os.Chdir(finalPath)
	if err != nil {
		return err
	}

	var ds []string

	ctx.DirStack = helpers.AppendIfMissing(ds, ctx.DirStack[len(ctx.DirStack)-1])
	ctx.DirStack = helpers.AppendIfMissing(ctx.DirStack, finalPath)

	fmt.Println(ctx.DirStack)

	return nil
}
