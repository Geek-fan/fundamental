package DI

type level0 struct {
	message string
}

type level1 struct {
	Base *level0
}

type level2 struct {
	L1 *level1
	L0 *level0
}

func newLevel0(msg string) *level0 {
	return &level0{msg}
}

func newLevel1(base *level0) *level1 {
	return &level1{base}
}

func newLevel2(base *level1, lowerBase *level0) *level2 {
	return &level2{base, lowerBase}
}
