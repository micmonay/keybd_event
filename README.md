# keybd_event


## For simulate key press in Linux, Windows and Mac in golang

### An example :
```go
package main

import (
	"runtime"
	"time"
	"github.com/micmonay/keybd_event"
)

func main() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		panic(err)
	}

	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}
	
	//set keys
	kb.SetKeys(keybd_event.VK_A, keybd_event.VK_B) 

	//set shif is pressed
	kb.HasSHIFT(true) 

	//launch
	err = kb.Launching() 
	if err != nil {
		panic(err)
	}
	//Ouput : AB
}
```
