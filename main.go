package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var versionRe = regexp.MustCompile(`\d+\.\d+\.\d+`)

func main() {
	if err := Main(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func Main(args []string) error {
	if len(args) != 3 {
		return fmt.Errorf("Usage: vup <major|minor|patch> <fname>")
	}
	target := args[1]
	fname := args[2]
	b, err := ioutil.ReadFile(fname)
	if err != nil {
		return err
	}
	src := string(b)
	match := versionRe.FindString(src)
	if match == "" {
		return fmt.Errorf("%s does not include a version", fname)
	}
	v, err := versionUp(match, target)
	if err != nil {
		return err
	}
	newContent := versionRe.ReplaceAllString(src, v)
	return ioutil.WriteFile(fname, []byte(newContent), 0644)
}

func versionUp(version string, target string) (string, error) {
	v := strings.Split(version, ".")
	major, _ := strconv.Atoi(v[0])
	minor, _ := strconv.Atoi(v[1])
	patch, _ := strconv.Atoi(v[2])

	switch target {
	case "major":
		return fmt.Sprintf("%d.0.0", major+1), nil
	case "minor":
		return fmt.Sprintf("%d.%d.0", major, minor+1), nil
	case "patch":
		return fmt.Sprintf("%d.%d.%d", major, minor, patch+1), nil
	default:
		return "", fmt.Errorf("%s is bad target", target)
	}
}
