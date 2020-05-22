package distance

import (
	"fmt"
	"testing"
)

func TestDistance_Input(t *testing.T) {
	conf := Config{
		Latitude:  1,
		Longitude: 1,
	}
	d := NewDistance(conf)
	d.Input(&Destination{
		Latitude:  60,
		Longitude: 0,
		Distance:  0,
	})


	d.Input(&Destination{
		Latitude:  66,
		Longitude: 0,
		Distance:  0,
	})

	d.Input(&Destination{
		Latitude:  2,
		Longitude: 0,
		Distance:  0,
	})

	d.Input(&Destination{
		Latitude:  7,
		Longitude: 0,
		Distance:  0,
	})

	fmt.Println(d.Elements)
	for _,v := range d.Elements {
		t.Error(v)
	}
}
