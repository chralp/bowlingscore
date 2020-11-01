package bowling

import (
    "testing"
)

// Test includes marker test
func TestAddRoll(t *testing.T) {

	//Ensure no active games exists
	GAMES = nil

	Startgame(1)
	cg := GetCurrentGame("player1")

	// New Started Game should have frameState Started
	if(cg.Frames[cg.CurrentFrame - 1].FrameState != "Started") {
		t.Fail()
	}

	// Add roll to player1	
	GetCurrentGame("player1").AddRoll("x")

	// Played Frame should have FrameStatus completed
	if(cg.Frames[cg.CurrentFrame - 2].FrameState != "completed") {
		t.Fail()
	}

	// Upcoming NON played frame should have FrameStatus Next
	if(cg.Frames[cg.CurrentFrame - 1].FrameState != "next") {
		t.Fail()
	}

	// Two frames ahead should have status locked
	if(cg.Frames[cg.CurrentFrame].FrameState != "locked") {
		t.Fail()
	}

	// Add roll to player1
	GetCurrentGame("player1").AddRoll("9")
	
	// Current frame with roll [ - or 1 - 9] should have status inprogress
	if(cg.Frames[cg.CurrentFrame - 1].FrameState != "inprogress") {
		t.Fail()
	}

	// Current frame should have one entry in .Rolls[]
	if(len(cg.Frames[cg.CurrentFrame - 2].Rolls) != 1) {
		t.Fail()
	}
	
	// Add roll to player1
	GetCurrentGame("player1").AddRoll("/")

	// Current frame should have two entries in .Rolls[]
	if(len(cg.Frames[cg.CurrentFrame - 2].Rolls) != 2) {
		t.Fail()
	}
}
