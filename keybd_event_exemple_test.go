package keybd_event

import (
	"time"
	"runtime"
)

func ExempleNewKeyBounding(){
	kb,err := NewKeyBounding()
	if err!=nil {
		panic(err)
	}
	// for linux is very important wait 2 second
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	kb.SetKeys(VK_A,VK_B) //set keys

	kb.HasSHIFT(true) //set shif is pressed

	err = kb.Launching() //launch
	if err!=nil {
		panic(err)
	}
	//Ouput : AB
}
