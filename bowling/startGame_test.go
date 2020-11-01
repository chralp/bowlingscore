package bowling

import (
    "testing"
	"fmt"
)

func TestStartOneGame(t *testing.T) {
	Startgame(1)
	if len(GAMES) != 2 {
		fmt.Println("GAMES STARTED", len(GAMES))
		t.Fail()
	}
}

func TestCreateEmptyFrames(t *testing.T) {
	emptyFrames := CreateEmptyFrames()
	if len(emptyFrames) != 10 {
		t.Fail()
	}
}

func TestGetLane(t *testing.T) {
	lane := GetLane()
	if !(lane >= 1 && lane <= 12) {
		t.Fail()
	}

}
