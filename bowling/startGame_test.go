package bowling

import (
    "testing"
)

func TestStartOneGame(t *testing.T) {
	GAMES = nil
	Startgame(1)
	if len(GAMES) != 1 {
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
