package basic

import (
	"fmt"
	"github.com/chai2010/errors"
	"io/ioutil"
)

func ParseJson(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err =fmt.Errorf("Json internal error : %v", p)
		}
	}()

	return err
}

func loadConfig() error {
	_, err := ioutil.ReadFile("/path/to/file")
	if err != nil {
		return errors.Wrap(err, "read failed")
	}

	return nil
}

func setup() error {
	err := loadConfig()
	if err != nil {
		return errors.Wrap(err, "invalid config")
	}
	return nil
}

func Run2() {
	if err := setup(); err != nil {
		for i, e := range err.(errors.Error).Wraped() {
			fmt.Printf("wraped(%d): %v\n", i, e)
		}
		fmt.Println("\n")
		for i, x := range err.(errors.Error).Caller() {
			fmt.Printf("caller:%d: %s\n", i, x.FuncName)
		}
	}
}
