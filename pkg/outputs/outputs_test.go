package outputs

import (
	"fmt"
	"testing"
)

func TestCommands(*testing.T) {

	exists, err := ioExists(OutputMaps, "UO1")
	fmt.Println(exists, err)
	if err != nil {
		return
	}

}
