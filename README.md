# keybd_event


## For simulate key press in Linux, Windows and Mac in golang

### An example :
```go
    kb,err := keybd.NewKeyBounding()
    if err!=nil {
        panic(err)
    }
    // For linux, it is very important wait 2 seconds
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
