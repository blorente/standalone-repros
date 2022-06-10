package lib_test

import (
	"github.com/blorente/standalone-repros/go/lib"
	"testing"
)

func TestLib(t *testing.T) {
	res := lib.MyFun(1)
	if res != 1 {
		t.Fail()
	}
}