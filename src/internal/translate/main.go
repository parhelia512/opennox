package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"golang.org/x/tools/go/analysis/multichecker"

	"github.com/noxworld-dev/opennox/v1/internal/noxfactor/refactor"
	"github.com/noxworld-dev/opennox/v1/internal/translate/noxanalyz"
)

func main() {
	checkRoot()
	if !wasConverted() {
		if err := moveFilesAndConvert(); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
	log.Printf("running noxfactor")
	if err := refactor.Run("./"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	log.Printf("running analyzer (may take a few minutes)")
	os.Setenv("CGO_CFLAGS_ALLOW", "(-fshort-wchar)|(-fno-strict-aliasing)|(-fno-strict-overflow)")
	os.Setenv("CGO_ENABLED", "1")
	os.Setenv("GOARCH", "386")
	multichecker.Main(noxanalyz.Analyzer)
}

func checkRoot() {
	_, err := os.Stat("go.mod")
	if err != nil {
		fmt.Fprintln(os.Stderr, "this tool must be run from the project's 'src' directory")
		os.Exit(1)
	}
}

func wasConverted() bool {
	list, _ := filepath.Glob("zzz_*.go")
	return len(list) > 0
}

func moveFilesAndConvert() error {
	log.Printf("using cxgo to convert C files to Go")
	cmd := exec.Command("cxgo")
	cmd.Dir, _ = filepath.Abs("..")
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	log.Printf("moving new Go files to the root directory")
	if err := move("./", "../gonox/*.go"); err != nil {
		return err
	}
	log.Printf("moving C files to a separate directory")
	if err := moveCFiles(); err != nil {
		return err
	}
	return nil
}

func moveCFiles() error {
	if err := os.MkdirAll("../csrc", 0755); err != nil {
		return err
	}
	if err := move("../csrc", "./*.c"); err != nil {
		return err
	}
	return nil
}

func move(dst, src string) error {
	list, err := filepath.Glob(src)
	if err != nil {
		return err
	}
	for _, name := range list {
		dname := filepath.Join(dst, filepath.Base(name))
		if err = os.Rename(name, dname); err != nil {
			return err
		}
		if err = exec.Command("git", "add", dname).Run(); err != nil {
			return err
		}
	}
	return nil
}
