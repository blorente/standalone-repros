package lib

import (
	"testing"
	"github.com/golang/mock/gomock"	
	"github.com/blorente/standalone-repros/go/lib/mocks"
)

func TestMyFunc(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mock := mock_lib.NewMockMyInterface(ctrl)

	want := "Hello, World!"
	mock.
		EXPECT().
		World(gomock.Eq("Hello")).
		Return(want)
	
	msg := MyFun(mock)

	if msg != want {
		t.Fatalf(`MyFunc("Hello") = %q, want match for %q`, msg, want)
	}
}
