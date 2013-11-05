/* Simulation of the ball clock */
package clock

import (
	"github.com/bgmerrell/goballclock/ball"
	"github.com/bgmerrell/goballclock/ballholders"
	"math"
)

// static ballholder capacities
const HOUR_RAIL_CAP = 11
const FIVE_MIN_RAIL_CAP = 11
const ONE_MIN_RAIL_CAP = 4

var queue ballholders.Queue
var hourRail ballholders.Rail
var fiveMinRail ballholders.Rail
var oneMinRail ballholders.Rail
var nClockRefreshes uint64

func init() {
	hourRail = ballholders.NewRail(HOUR_RAIL_CAP)
	fiveMinRail = ballholders.NewRail(FIVE_MIN_RAIL_CAP)
	oneMinRail = ballholders.NewRail(ONE_MIN_RAIL_CAP)
}

// Update the clock state by adding ball
func updateClockState(b ball.Ball) {
	var spilledBalls []ball.Ball

	spilledBalls = oneMinRail.Push(b)
	if len(spilledBalls) == 0 {
		return
	}
	queue.Push(spilledBalls)

	spilledBalls = fiveMinRail.Push(b)
	if len(spilledBalls) == 0 {
		return
	}
	queue.Push(spilledBalls)

	spilledBalls = hourRail.Push(b)
	if len(spilledBalls) == 0 {
		return
	}
	queue.Push(append(spilledBalls, b))
}

// Detect a cycle occurrence in a ball clock and track time for that cycle to
// occur
func findCycle(queueCapacity uint8) {
	// break when the balls are all back in their original positions in the
	// queue
	for {
		ball := queue.Pop()
		updateClockState(ball)
		if queue.IsFull() {
			nClockRefreshes++
			if queue.DoCycleCheck() {
				break
			}
		}
	}
}

func GetDaysUntilCycle(queueCapacity uint8) uint64 {
	// Number of times the clock refreshes, i.e., the number of 12-hour
	// periods
	nClockRefreshes = 0
	queue = ballholders.NewQueue(queueCapacity)
	findCycle(queueCapacity)
	// There 2 clock refreshes in a day
	return uint64(math.Ceil(float64(nClockRefreshes) / 2.0))
}
