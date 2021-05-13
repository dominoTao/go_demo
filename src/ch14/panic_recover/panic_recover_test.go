package panic_recover

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer fmt.Println("start ")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from ", err, reflect.TypeOf(err))
		}
	}()
	//os.Exit(-1)
	panic(errors.New("something wrong"))
}