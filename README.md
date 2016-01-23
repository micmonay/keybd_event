# keybd_event


## For simulate key press in Linux, Windows and Mac in golang (actualy in beta)

### An exemple :
```go
    kb,err := keybd.NewKeyBounding()
    if err!=nil {
        panic(err)
    }
    // for linux is very important wait 2 second
    if runtime.GOOS == "linux" {
        time.Sleep(2 * time.Second)
    }
    kb.SetKeys(keybd.VK_A,keybd.VK_B) //set keys

    kb.HasSHIFT(true) //set shif is pressed

    err = kb.Launching() //launch
    if err!=nil {
        panic(err)
    }
    //Ouput : AB
```
