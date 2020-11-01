package bowling

// Mark Frame or game with status
func (cg* Game) marker(typeof string, status string) {

	switch typeof {
	case "frame":
		if(status == statusCompleted) {
			// Set CurrentFrame as completed
			cg.Frames[cg.CurrentFrame - 1].FrameState = status;
			// Mark next frame as Next
			cg.Frames[cg.CurrentFrame].FrameState = statusNext

			// Increase CurrentFrame int
			cg.CurrentFrame++
		} else {
			cg.Frames[cg.CurrentFrame - 1].FrameState = status;
		}
		break;
	case "game":
		cg.Frames[cg.CurrentFrame - 1].FrameState = statusCompleted;
		cg.GameState = statusCompleted
		break;
	}
	return;
}

// Handle each roll and append to frame[Number].Rolls[roll]
func (cg* Game) AddRoll(roll string) {

	// Current frame
	cf := &cg.Frames[cg.CurrentFrame - 1]

	indexPos := cg.CurrentFrame -1

	switch cg.CurrentFrame {

	case 1, 2, 3, 4, 5, 6, 7, 8, 9: // If frame is 1-9 we want the same outcome
		
		// Append roll to Rolls
		// We will mark the frame as complete if the first roll is a strike.
		// If the first roll is not a strike, then we will wait until secnd roll before marking frame as complete.		
		if(len(cf.Rolls) == 0) {
			cg.marker("frame", statusInprogress)
			cf.Rolls = append(cf.Rolls, roll)
			if(roll == "x") {
				cg.marker("frame", statusCompleted)
			}
			break;
		}

		if(len(cf.Rolls) == 1) {
			cf.Rolls = append(cf.Rolls, roll)
			cg.marker("frame", statusCompleted)
			break;			
		}

	case 10: // If frame is 10 we want to be able to handle three rolls

		// Append input to Rolls
		// We will mark the frame as complete if the first roll ( input ) is a strike.
		// If the first roll is not a strike, then we will wait until secnd roll before marking frame as complete.
		
		if (len(cf.Rolls) == 0) {
			cg.marker("frame", statusInprogress)
			cf.Rolls = append(cf.Rolls, roll)
			break;	
		}
		
 		if (len(cf.Rolls) == 1) {
			cf.Rolls = append(cf.Rolls, roll)
			if !(roll == "/" || roll == "x") {
				cg.marker("game", statusCompleted)
			}
			break;
		}

		if (len(cf.Rolls) == 2){
			cf.Rolls = append(cf.Rolls, roll)
			cg.marker("game", statusCompleted)
			break;
		}
	}

	// Pass frameNumber to start the count from.
	cg.countScore(indexPos)
	
	return
}
