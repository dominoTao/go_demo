package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer fmt.Println("start ")

	defer func() {
		fmt.Println("finally")
	}()
	//os.Exit(-1)
	panic(errors.New("something wrong"))
}