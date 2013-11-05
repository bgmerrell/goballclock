/*
things that hold balls

For example, the clock's ball queue is a queue.  The clock's time rails
are Rails, and both are BallHolders.

*/
package ballholders

import (
	"container/ring"
	"github.com/bgmerrell/goballclock/ball"
)

// A BallHolder is a thing that holds Balls
type BallHolder struct {
	// how much the ball holder can hold
	capacity uint8
	// how much the baller holder is holding
	nBalls uint8
}

// Create a new BallHolder
func NewBallHolder(capacity uint8, nBalls uint8) BallHolder {
	return BallHolder{capacity, nBalls}
}

func (bh BallHolder) IsFull() bool {
	return bh.capacity == bh.nBalls
}

// The clock's queue
//
// Queue uses a ring buffer to store the balls; instead of the balls moving
// within the Queue, the ring buffer is updated to point to the appropriate
// ball.
//
// Because balls are only ever appended, nBalls is used to determine which
// of the balls are valid, and the rest of the queue is considered empty.
type Queue struct {
	BallHolder
	ring *ring.Ring
}

// Create a new, full, BallHolder
func NewQueue(capacity uint8) Queue {
	bh := NewBallHolder(capacity, capacity)
	r := ring.New(int(capacity))
	for i := uint8(0); i < capacity; i++ {
		r.Value = ball.New(i)
		r = r.Next()
	}
	return Queue{bh, r}
}

// Get a ball from the beginning of the queue
func (q *Queue) Pop() ball.Ball {
	q.nBalls--
	ball := q.ring.Value.(ball.Ball)
	q.ring = q.ring.Next()
	return ball
}

// Return true if the balls are in their original position in the queue
func (q *Queue) DoCycleCheck() bool {
	if !q.IsFull() {
		return false
	}
	tmp := q.ring
	for i := uint8(0); i < q.capacity; i++ {
		ball := q.ring.Value.(ball.Ball)
		q.ring = q.ring.Next()
		if ball.Id != i {
			// restore ring
			q.ring = tmp
			return false
		}
	}
	return true
}

// Return a representation of the queue for testing
//
// -1 means empty
func (q *Queue) GetTestRepr() []int {
	repr := make([]int, q.capacity)
	for i := uint8(0); i < q.capacity; i++ {
		ball := q.ring.Value.(ball.Ball)
		q.ring = q.ring.Next()
		if i >= q.nBalls {
			repr[i] = -1 // empty
		} else {
			repr[i] = int(ball.Id)
		}
	}
	return repr
}

// Put an array of balls back to the end of the queue
func (q *Queue) Push(balls []ball.Ball) {
	tmp := q.ring
	q.ring = q.ring.Move(int(q.nBalls))
	for i := range balls {
		q.nBalls++
		q.ring.Value = balls[i]
		q.ring = q.ring.Next()
	}
	q.ring = tmp
}

// The clock's time rails
//
// A Rail holds Balls, but can spill them (down to another ball holder)
//
// Rail uses an array to store the balls.
//
// Unlike a Queue, balls are never "popped" one at a time; instead, when the
// rail is full, all balls are spilled in reverse order.
//
// Because balls are only ever appended, nBalls is used to determine where to
// put the balls in the array, and to determine which of the balls are valid,
// while the rest of the array is considered empty.
type Rail struct {
	BallHolder
	Balls []ball.Ball
}

// Create a new, empty, Rail
func NewRail(capacity uint8) Rail {
	bh := NewBallHolder(capacity, 0)
	balls := make([]ball.Ball, capacity)
	return Rail{bh, balls}
}

// Empty the ball holder and return a reversed list of the spilt Balls
func (r *Rail) spill() []ball.Ball {
	// Seriously, golang, no reverse abstraction? :\
	spilledBalls := make([]ball.Ball, r.capacity)
	for i := range r.Balls {
		spilledBalls[r.capacity-1-uint8(i)] = r.Balls[i]
	}
	return spilledBalls
}

// Add a ball to the rail.  If the rail is full, it will spill.
// A slice of spilled balls is returned.
func (r *Rail) Push(b ball.Ball) []ball.Ball {
	if r.IsFull() {
		// Reset state and spill
		r.nBalls = 0
		return r.spill()
	}

	r.Balls[r.nBalls] = b
	r.nBalls++
	return []ball.Ball{}
}

// Return a representation of the rail for testing
//
// -1 means empty
func (r *Rail) GetTestRepr() []int {
	repr := make([]int, r.capacity)
	for i := uint8(0); i < r.capacity; i++ {
		if i >= r.nBalls {
			repr[i] = -1 // empty
		} else {
			repr[i] = int(r.Balls[i].Id)
		}
	}
	return repr
}
