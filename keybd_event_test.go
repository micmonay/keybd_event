package keybd_event

import (
	"runtime"
	"time"
)

func ExampleNewKeyBonding() {
	kb, err := NewKeyBonding()
	if err != nil {
		panic(err)
	}
	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	kb.SetKeys(VK_SP2) //set keys

	kb.HasALTGR(true) //set shif is pressed

	err = kb.Launching() //launch
	if err != nil {
		panic(err)
	}
	// Output: AB
}
