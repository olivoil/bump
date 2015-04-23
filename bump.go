package bump

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/blang/semver"
)

var re = regexp.MustCompile(`\d+\.\d+\.\d+`)

// Bump increments the version numbers contained in a file,
// and saves the file
func File(command, fileName string) error {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	d, err := ByteSlice(command, b)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, d, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Bump increments the version numbers contained in a byte slice,
// and returns a new byte slice with updated version numbers
func ByteSlice(command string, b []byte) ([]byte, error) {
	matches := re.FindAll(b, 1)
	if matches == nil {
		return b, nil
	}

	if len(matches) > 1 {
		return b, errors.New("more than 1 version number found")
	}

	v, err := String(command, string(matches[0]))
	if err != nil {
		return b, err
	}

	return re.ReplaceAll(b, []byte(v)), nil
}

// Bump increments the version number represented by a string,
// and returns a new string representing the updated version number
func String(command, v string) (string, error) {
	version, err := semver.Parse(v)
	if err != nil {
		return "", err
	}

	if err := bump(command, &version); err != nil {
		return "", err
	}

	return version.String(), nil
}

func bump(command string, v *semver.Version) error {
	switch command {
	case "major":
		v.Major++
		v.Minor = 0
		v.Patch = 0
	case "minor":
		v.Minor++
		v.Patch = 0
	case "patch":
		v.Patch++
	default:
		return errors.New(fmt.Sprintf("%v is not a supported command (major, minor, patch)", command))
	}

	return nil
}
