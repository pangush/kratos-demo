package paginate

import (
	"fmt"
	"testing"
)

func TestNewPaginate(t *testing.T) {
	conf := Config{
		CurrentPage: 0,
		Limit:       0,
	}

	p := NewPaginate(conf)
	fmt.Printf("%+v", p)
}

func TestPaginate_SetTotal(t *testing.T) {
	conf := Config{
		CurrentPage: 2,
		Limit:       30,
	}

	p := NewPaginate(conf)

	p.SetTotal(0)
	fmt.Printf("%+v \n", p)

	p.SetTotal(19)
	fmt.Printf("%+v \n", p)

	p.SetTotal(20)
	fmt.Printf("%+v \n", p)

	p.SetTotal(21)
	fmt.Printf("%+v \n", p)

	p.SetTotal(30)
	fmt.Printf("%+v \n", p)

	p.SetTotal(40)
	fmt.Printf("%+v \n", p)

	p.SetTotal(41)
	fmt.Printf("%+v \n", p)

}
