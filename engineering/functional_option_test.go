package engineering

import "testing"

func TestNewPerson(t *testing.T) {
	p := NewPerson("张三", SetAge(18), SetAddress("Shanghai"))
	t.Log(p)
}