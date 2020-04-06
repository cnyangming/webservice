package test

import (
	"github.com/pkg/errors"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	err := errors.New("Custome error")
	log.Fatalf("%s", err)
	panic(err)
}
