package keybd_event_test

import (
	"time"
	"runtime"
	"github.com/micmonay/keybd_event"
)

func ExempleNewKeyBounding(){
	kb,err := keybd_event.NewKeyBounding()
	if err!=nil {
		panic(err)
	}
	// for linux is very important wait 2 second
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	kb.SetKeys(keybd_event.VK_A,keybd_event.VK_B) //set keys

	kb.HasSHIFT(true) //set shif is pressed

	err = kb.Launching() //launch
	if err!=nil {
		panic(err)
	}
	// Output: AB
}
