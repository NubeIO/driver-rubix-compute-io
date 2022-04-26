package outputs

import (
	"fmt"
	"testing"
)

func TestCommands(*testing.T) {

	exists, pin, err := SupportsPWM(OutputMaps, "UO3")
	fmt.Println("SupportsPWM", exists, err, "pin", pin)
	if err != nil {
		return
	}
	exists, err = ioExists(OutputMaps, "UO3")
	fmt.Println("ioExists", exists, err)
	if err != nil {
		return
	}

}
