package board

import "testing"

type testPositionEqual struct {
	aPosition *Position
	bPosition *Position
	expect bool
	description string
}

func TestPositionEqual(t *testing.T) {
	values := []*testPositionEqual{
		{
			aPosition: &Position{
				x: 0,
				y: 0,
			},
			bPosition: &Position{
				x: 0,
				y: 0,
			},
			expect: true,
			description: "both same case",
		},
		{
			aPosition: &Position{
				x: 0,
				y: 1,
			},
			bPosition: &Position{
				x: 0,
				y: 0,
			},
			expect: false,
			description: "x different case",
		},
		{
			aPosition: &Position{
				x: 1,
				y: 0,
			},
			bPosition: &Position{
				x: 0,
				y: 0,
			},
			expect: false,
			description: "y different case",
		},
		{
			aPosition: &Position{
				x: 1,
				y: 1,
			},
			bPosition: &Position{
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