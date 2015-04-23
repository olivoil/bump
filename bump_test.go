package bump_test

import (
	"bytes"
	"testing"

	"github.com/olivoil/bump"
)

var stringVersions = map[string]map[string]string{
	"0.0.0": {"major": "1.0.0", "minor": "0.1.0", "patch": "0.0.1"},
	"0.0.1": {"major": "1.0.0", "minor": "0.1.0", "patch": "0.0.2"},
	"1.0.9": {"major": "2.0.0", "minor": "1.1.0", "patch": "1.0.10"},
	"2.2.2": {"major": "3.0.0", "minor": "2.3.0", "patch": "2.2.3"},
	"1.2.3": {"major": "2.0.0", "minor": "1.3.0", "patch": "1.2.4"},
}

var actual = `
VERSION=0.0.1
ENV=test
V=0.0.1
`

var patch = `
VERSION=0.0.2
ENV=test
V=0.0.2
`

var minor = `
VERSION=0.1.0
ENV=test
V=0.1.0
`

var major = `
VERSION=1.0.0
ENV=test
V=1.0.0
`
var bytesVersions = map[string]map[string]string{
	actual: {
		"patch": patch,
		"minor": minor,
		"major": major,
	},
}

func TestString(t *testing.T) {
	for version, commands := range stringVersions {
		for command, expected := range commands {
			actual, err := bump.String(command, version)
			if err != nil {
				t.Errorf(`Error bumping %s to %s version: %+v`, version, command, err)
			}
			if actual != expected {
				t.Errorf(`Expected %s of "%s" to be "%s", got "%s".`, command, version, expected, actual)
			}
		}
	}
}

func TestByteSlice(t *testing.T) {
	for version, commands := range bytesVersions {
		for command, expected := range commands {
			actual, err := bump.ByteSlice(command, []byte(version))
			if err != nil {
				t.Errorf(`Error bumping %s to %s version: %+v`, version, command, err)
			}
			if bytes.Compare(actual, []byte(expected)) != 0 {
				t.Errorf("Expected %s version of:\n%s\nto be\n%s\nbut got\n%s\n", command, version, expected, string(actual))
			}
		}
	}
}
