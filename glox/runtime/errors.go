package runtime

import (
	"fmt"
)

func RaiseError(line int, state *State) {
	state.HadError = true
	fmt.Printf("Error on %d\n", line)
}
