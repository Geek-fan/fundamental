package DI

import (
	"github.com/yfgogogo/fundamental/engineering/reflect"
	"testing"
)

func Test_Wire(t *testing.T) {
	res := InitLevel2("base")
	t.Logf("%s", reflect.RecursivePrint(res))
}

func Test_Dig(t *testing.T) {
	err := Build()
	if err != nil {
		t.Errorf("%v", err)
	}
}
