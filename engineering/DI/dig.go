package DI

import (
	"fmt"
	"github.com/yfgogogo/fundamental/engineering/reflect"
	"go.uber.org/dig"
)

func Build() error {
	c := dig.New()
	c.Provide(func() string {
		return "base"
	})
	c.Provide(newLevel0)
	c.Provide(newLevel1)
	c.Provide(newLevel2)

	f := func(l2 *level2) {
		str := reflect.RecursivePrint(l2)
		fmt.Printf("%s", str)
	}
	return c.Invoke(f)
}
