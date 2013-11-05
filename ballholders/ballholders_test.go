package ballholders

import (
	"fmt"
	"github.com/bgmerrell/goballclock/ball"
	"testing"
)

func TestNewBallHolder(t *testing.T) {
	const EXPECTED_CAPACITY = 4
	const EXPECTED_NBALLS = EXPECTED_CAPACITY
	bh := NewBallHolder(EXPECTED_CAPACITY, EXPECTED_CAPACITY)
	if bh.capacity != EXPECTED_CAPACITY {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			bh.capacity, EXPECTED_CAPACITY)
	}
	if bh.nBalls != EXPECTED_NBALLS {
		t.Errorf("Unexpected nBalls (actual %d, expected %d)",
			bh.nBalls, EXPECTED_NBALLS)
	}
	if !bh.IsFull() {
		t.Errorf("Expected ballHolder to be full")
	}
	bh.nBalls--
	if bh.IsFull() {
		t.Errorf("Expected ballHolder not to be full")
	}
}

func TestNewQueue(t *testing.T) {
	const EXPECTED_CAPACITY = 11
	const EXPECTED_NBALLS = EXPECTED_CAPACITY
	q := NewQueue(EXPECTED_CAPACITY)
	if q.capacity != EXPECTED_CAPACITY {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			q.capacity, EXPECTED_CAPACITY)
	}
	if q.nBalls != EXPECTED_NBALLS {
		t.Errorf("Unexpected nBalls (actual %d, expected %d)",
			q.nBalls, EXPECTED_NBALLS)
	}
	if q.ring.Len() != EXPECTED_CAPACITY {
		t.Errorf("Unexpected length of ring (actual %d, expected %d)",
			q.ring.Len(), EXPECTED_CAPACITY)
	}
	// Go through ring twice to test it
	for i := 0; i < 2; i++ {
		for i := 0; i < q.ring.Len(); i++ {
			if q.ring.Value.(ball.Ball).Id != uint8(i) {
				t.Errorf("Ball out of order (actual %d, expected %d)",
					q.ring.Value.(ball.Ball).Id,
					i)
			}
			q.ring = q.ring.Next()
		}
	}
}

func TestQueue(t *testing.T) {
	const CAPACITY = 11
	q := NewQueue(CAPACITY)
	b := q.Pop()
	// First popped ball should have an ID of 0
	if b.Id != 0 {
		t.Errorf("Unexpected ball ID (actual %d, expected %d)", b.Id, 0)
	}
	b = q.Pop()
	// Second popped ball should have an ID of 1
	if b.Id != 1 {
		t.Errorf("Unexpected ball ID (actual %d, expected %d)", b.Id, 1)
	}
	q.Push([]ball.Ball{ball.New(200)})
	b = q.Pop()
	// popped ball should have an ID of 2
	if b.Id != 2 {
		t.Errorf("Unexpected ball ID (actual %d, expected %d)", b.Id, 2)
	}
	// after we've gone through the original 11 (8 more), the next ball we
	//  see should have id 200
	for i := 0; i < 8; i++ {
		b = q.Pop()
	}
	b = q.Pop()
	// popped ball should have an ID of 200
	if b.Id != 200 {
		t.Errorf("Unexpected ball ID (actual %d, expected %d)", b.Id, 200)
	}
}

func TestNewRail(t *testing.T) {
	const EXPECTED_CAPACITY = 11
	const EXPECTED_NBALLS = 0
	r := NewRail(EXPECTED_CAPACITY)
	if r.capacity != EXPECTED_CAPACITY {
		t.Errorf("Unexpected capacity (actual %d, expected %d)",
			r.capacity, EXPECTED_CAPACITY)
	}
	if r.nBalls != EXPECTED_NBALLS {
		t.Errorf("Unexpected nBalls (actual %d, expected %d)",
			r.nBalls, EXPECTED_NBALLS)
	}
}

func TestRailPushAndSpill(t *testing.T) {
	const EXPECTED_CAPACITY = 4
	const EXPECTED_NBALLS = 0
	var spilledBalls []ball.Ball
	r := NewRail(EXPECTED_CAPACITY)

	// Push a ball with an ID of 1, and check the Balls slice
	spilledBalls = r.Push(ball.New(1))
	if r.Balls[0].Id != 1 {
		t.Errorf("Unexpected ball ID after rail push (actual %d, expected %d)",
			r.Balls[0].Id, 1)
	}
	// Nothing should have spilled
	if len(spilledBalls) != 0 {
		t.Errorf("Unexpected spilled balls (%v), expected no spillage", spilledBalls)
	}

	// Push a three more balls
	spilledBalls = r.Push(ball.New(2))
	if len(spilledBalls) != 0 {
		t.Errorf("Unexpected spilled balls (%v), expected no spillage", spilledBalls)
	}
	spilledBalls = r.Push(ball.New(3))
	if len(spilledBalls) != 0 {
		t.Errorf("Unexpected spilled balls (%v), expected no spillage", spilledBalls)
	}
	spilledBalls = r.Push(ball.New(4))
	if len(spilledBalls) != 0 {
		t.Errorf("Unexpected spilled balls (%v), expected no spillage", spilledBalls)
	}
	for i := range r.Balls {
		// i + 1, because we started at 1 to distinguish between test
		// the zero-value of the array
		if r.Balls[i].Id != uint8(i+1) {
			t.Errorf("Unexpected ball ID after rail push (actual %d, expected %d)",
				r.Balls[i].Id, i+1)
		}
	}

	if !r.IsFull() {
		t.Fatalf("Expected rail to be full")
	}

	// OK, rail is full and thus spill on the next push
	spilledBalls = r.Push(ball.New(5))
	if len(spilledBalls) == 0 {
		t.Errorf("Expected spilled balls, but got no spillage")
	}
	expected := []ball.Ball{ball.New(4), ball.New(3), ball.New(2), ball.New(1)}
	if fmt.Sprintf("%v", spilledBalls) != fmt.Sprintf("%v", expected) {
		t.Errorf("Unexpected spilled balls:\n"+
			"Actual: %v\n"+
			"Expected: %v",
			spilledBalls,
			expected)
	}
}
