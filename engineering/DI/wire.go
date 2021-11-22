//go:build wireinject
// +build wireinject

package DI

import "github.com/google/wire"

func InitLevel2(msg string) *level2 {
	wire.Build(newLevel0, newLevel1, newLevel2)
	return &level2{}
}
