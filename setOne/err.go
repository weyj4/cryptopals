package main

import (
	"errors"
	"log"
)

func Check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func CheckCustom(cond bool, err string) {
	if cond {
		log.Fatal(errors.New(err))
	}
}
