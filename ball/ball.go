/*
a ball of the clock
*/
package ball

type Ball struct {
	// The original position of the ball in a ball holder
	Id uint8
}

func New(id uint8) Ball {
	return Ball{id}
}
