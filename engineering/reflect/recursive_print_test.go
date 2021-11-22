package reflect

import (
	"io"
	"os"
	"testing"
)

func Test_RecursivePrint(t *testing.T) {
	type age struct {
		age int
	}
	type address struct {
		city string
	}
	type person struct {
		Name   string
		gender string
		Age    *age     /* exported pointer */
		addr   address  /* unexported struct */
		addr2  *address /* unexported pointer */
	}
	f := func(a int) int {
		return a
	}
	var (
		r io.Reader
		w io.Writer
	)
	w = os.Stdout
	inputs := []interface{}{
		nil,
		123,
		"test",
		f,
		r,
		w,
		make(chan int),
		[3]int{1, 2, 3},
		[]person{{Name: "Jack"}, {Name: "Mary"}},
		&person{"Tom", "man", &age{10}, address{"Shanghai"}, &address{"Beijing"}},
	}

	for _, input := range inputs {
		t.Run("recursive print", func(t *testing.T) {
			t.Logf("%+v", input)
			if got := RecursivePrint(input); got == "not supported" {
				t.Errorf("not supported")
			} else {
				t.Logf("%s\n", got)
			}
		})
	}

}

func Test_printPtr(t *testing.T) {
	tests := []struct {
		arg  uintptr
		want string
	}{
		{0xffff, "0x0000ffff"},
		{0, "0x00000000"},
		{0x10000000, "0x10000000"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := printPtr(tt.arg); got != tt.want {
				t.Errorf("printPtr() = %v, want %v", got, tt.want)
			}
		})
	}
}
