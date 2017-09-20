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
	var conf Config

	conf.SetPath("testfiles/config.yml")

	err := conf.Parse()
	if err != nil {
		t.Errorf("parsing of config file failed %s", conf.Path)
	}
	if conf.Botconfig.RunAsDaemon != false ||
		conf.Botconfig.Token != "testtoken" ||
		conf.Databaseconfig.Path != "testpath" {
		t.Errorf("parsing of testfile failed [%+v]", conf)
	}
	return
}
