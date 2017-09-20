package configuration

import (
	"errors"
	"testing"
)

func TestSetPath(t *testing.T) {
	var conf Config

	conf.SetPath("this is a testpath")

	if conf.Path != "this is a testpath" {
		err := errors.New("path was not set properly")
		panic(err)
	}
}

func TestParseConfig(t *testing.T) {
	//TODO Implement this and add testfiles
	return
}
