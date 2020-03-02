# keybd_event
This library simulates the key press on the keyboard.

**Important :** 
- The keys change in the different keyboard layout configuration of the target computer.
- I have tested this code on the different system and I don't find the error, but I don't granted this update have no bug. If you have a bug, please create an issue.

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

	//set shift is pressed
	kb.HasSHIFT(true) 

	//launch
	err = kb.Launching() 
	if err != nil {
		panic(err)
	}
	//Ouput : AB
}
```

For easy access of all keys on the virtual keyboard, I have added more special keycodes constants `VK_SP*`. 

The next picture is a good solution to understand

![keyboard.png](./keyboard.png)

## Linux

On Linux this library use **uinput**, but generally, on the major distributions, the uinput is only for the root user. 

The easy solution is executing on root user or change permission by `chmod`, but it is not good.

You can follow the next example, for more security.

```bash
sudo groupadd uinput
sudo usermod -a -G uinput my_username
sudo udevadm control --reload-rules
echo "SUBSYSTEM==\"misc\", KERNEL==\"uinput\", GROUP=\"uinput\", MODE=\"0660\"" | sudo tee /etc/udev/rules.d/uinput.rules
echo uinput | sudo tee /etc/modules-load.d/uinput.conf
```

Another subtlety on Linux, it is important after creating keybd_event, to waiting 2 seconds before running first keyboard actions

## Darwin (MAC OS)
This library depends on the frameworks Apple, I did not find a solution for cross-compilation.
