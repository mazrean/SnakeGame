package board

import "testing"

type testPositionEqual struct {
	aPosition *position
	bPosition *position
	expect bool
	description string
}

func TestPositionEqual(t *testing.T) {
	values := []*testPositionEqual{
		{
			aPosition: &position{
				x: 0,
				y: 0,
			},
			bPosition: &position{
				x: 0,
				y: 0,
			},
			expect: true,
			description: "both same case",
		},
		{
			aPosition: &position{
				x: 0,
				y: 1,
			},
			bPosition: &position{
				x: 0,
				y: 0,
			},
			expect: false,
			description: "x different case",
		},
		{
			aPosition: &position{
				x: 1,
				y: 0,
			},
			bPosition: &position{
				x: 0,
				y: 0,
			},
			expect: false,
			description: "y different case",
		},
		{
			aPosition: &position{
				x: 1,
				y: 1,
			},
			bPosition: &position{
				x: 0,
				y: 0,
			},
			expect: false,
			description: "both different case",
		},
	}

	for _,v := range values {
		res := v.aPosition.equal(v.bPosition)
		if res != v.expect {
			t.Fatalf("Failed: %s", v.description)
		}
	}
}