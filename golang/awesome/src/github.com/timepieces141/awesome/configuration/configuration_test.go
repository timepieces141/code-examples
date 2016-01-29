package configuration

import (
    "go/build"
    "testing"
)

var badTestConfigData string = `
{
    "SomethingNotRequired": "foobar"
}`

var goodTestConfigData string = `
{
    "SomethingRequired": "foobar",
    "SomethingNotRequired": "foobar"
}`

var testConfigDir = build.Default.GOPATH + "/src/github.com/timepieces141/awesome/test/config/"

func TestBadConfiguration(t *testing.T) {
    err := construct([]byte(badTestConfigData))
    if err == nil {
        t.Error("Expected an error when sending bad test config data")
    }
}

func TestGoodConfiguration(t *testing.T) {
    err := construct([]byte(goodTestConfigData))
    if err != nil {
        t.Error("Not expecting an error, received:", err.Error())
    }

    if Config.SomethingRequired != "foobar" {
        t.Errorf("Expected '%s', received '%s'", "foobar", Config.SomethingRequired)
    }

    if Config.SomethingNotRequired != "foobar" {
        t.Errorf("Expected '%s', received '%s'", "foobar", Config.SomethingNotRequired)
    }
}

func TestBadConfigurationFile(t *testing.T) {
    err := ReadConfig(testConfigDir + "bad_config.json")
    if err == nil {
        t.Error("Expected an error when sending bad test config data")
    }
}

func TestGoodConfigurationFile(t *testing.T) {
    err := ReadConfig(testConfigDir + "good_config.json")
    if err != nil {
        t.Error("Not expecting an error, received:", err.Error())
    }

    if Config.SomethingRequired != "foobar" {
        t.Errorf("Expected '%s', received '%s'", "foobar", Config.SomethingRequired)
    }

    if Config.SomethingNotRequired != "foobar" {
        t.Errorf("Expected '%s', received '%s'", "foobar", Config.SomethingNotRequired)
    }
}