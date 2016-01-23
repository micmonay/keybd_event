# keybd_event


## For simulate key press in Linux, Windows and Mac in golang (actualy in beta)

### An exemple :
```go
    kb,err := keybd.NewKeyBounding()
	if err!=nil {
		panic(err)
	}
	// for linux is very important wait 2 second
	time.Sleep(2*time.Second)
	// set keys
	kb.SetKeys(keybd.VK_0,keybd.VK_A)
	// Set shif is pressed
	kb.HasSHIFT(true)
	// Launch event
	err = kb.Launching()
	if err!=nil {
		panic(err)
	}
```
