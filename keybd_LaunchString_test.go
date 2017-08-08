package keybd_event

import (
	"fmt"
	"runtime"
	"time"
)

func ExampleLaunchString() {
	if runtime.GOOS != "darwin" {
		fmt.Println("LaunchString() is implemeted only on darwin os")
		return
	}

	kb, err := NewKeyBonding()
	if err != nil {
		panic(err)
	}
	// For linux, it is very important wait 2 seconds
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	kb.LaunchString("1234567890")
	kb.LaunchString("!@#$%^&*()")
	kb.LaunchString("QWERTYUIOP")
	kb.LaunchString("asdfghjkl")
	kb.LaunchString("-=[]\\;',./'")
	kb.LaunchString("_+{}|:\"<>?")

	// Output:
	//1234567890!@#$%^&*()QWERTYUIOPasdfghjkl-=[]\;',./'_+{}|:"<>?
}
