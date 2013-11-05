package ball

import (
	"testing"
)

func TestNewBall(t *testing.T) {
	const BALL_ID = 0
	ball := New(BALL_ID)
	if ball.Id != BALL_ID {
		t.Errorf("Unexpected ball ID (Actual: %d, Expected: %d)\n", ball.Id, BALL_ID)
	}
}
