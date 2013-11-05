package clock

import (
	"github.com/bgmerrell/goballclock/ballholders"
	"fmt"
	"testing"
)

func TestUpdateClockState(t *testing.T) {
	const QUEUE_CAP = 27
	queue = ballholders.NewQueue(QUEUE_CAP)

	// Run ball 0 through the clock
	b := queue.Pop()
	updateClockState(b)

	// Check queue state
	actual := queue.GetTestRepr()
	expected := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, -1}

	if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
		t.Fatalf("Unexpected queue state\n"+
			"Actual: %v\n"+
			"Expected: %v",
			actual,
			expected)
	}

	// Check rail states
	actual = oneMinRail.GetTestRepr()
	expected = []int{0, -1, -1, -1}
	if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
		t.Fatalf("Unexpected queue state\n"+
			"Actual: %v\n"+
			"Expected: %v",
			actual,
			expected)
	}
	// And the other rails should be empty
	actual = fiveMinRail.GetTestRepr()
	expected = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
		t.Fatalf("Unexpected queue state\n"+
			"Actual: %v\n"+
			"Expected: %v",
			actual,
			expected)
	}
	// And the other rails should be empty
	actual = hourRail.GetTestRepr()
	expected = []int{-1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
		t.Fatalf("Unexpected queue state\n"+
			"Actual: %v\n"+
			"Expected: %v",
			actual,
			expected)
	}

	// Running 4 more balls through the clock shoul should result in the
	// first four balls going back on the queue in reverse order...
	for i := 0; i < 4; i++ {
		b = queue.Pop()
		updateClockState(b)
	}

	actual = queue.GetTestRepr()
	expected = []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 3, 2, 1, 0, -1}

	if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
			t.Fatalf("Unexpected queue state\n"+
				"Actual: %v\n"+
				"Expected: %v",
				actual,
				expected)
	}
	// ...And we should see the 4 ball show up on the next rail down
	actual = fiveMinRail.GetTestRepr()
	expected = []int{4, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}
	if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
		t.Fatalf("Unexpected queue state\n"+
			"Actual: %v\n"+
			"Expected: %v",
			actual,
			expected)
	}

	// The clock capacity is 11 hours, 59 minutes (719 minutes).  We've
	// already run 5.
	for i := 0; i < (719 - 5); i++ {
		b = queue.Pop()
		updateClockState(b)
	}

	actual = queue.GetTestRepr()
	for i := 0; i < QUEUE_CAP; i++ {
		// The clock should be full and there should only be one ball
		// in the queue
		if (i == 0 && actual[i] == -1) || (i != 0 && actual[i] != -1) {
			if fmt.Sprintf("%v", actual) != fmt.Sprintf("%v", expected) {
				t.Fatalf("Unexpected queue state\n"+
					"Actual: %v\n",
					actual)
			}
		}
	}

	// After one more ball run the queue should be full and the rails
	// should be empty
	b = queue.Pop()
	updateClockState(b)
	if !queue.IsFull() {
		t.Fatalf("Expected queue to be full")
	}
	// Check rail states
	for n, rail := range([]ballholders.Rail{oneMinRail, fiveMinRail, hourRail}) {
		actual = rail.GetTestRepr()
		for i := 0; i < len(actual); i++ {
			if actual[i] != -1 {
				t.Errorf("Expected rail %d to be empty:\n"+
					"Actual: %v\n",
					n,
					actual)
				break
			}
		}
	}
}
