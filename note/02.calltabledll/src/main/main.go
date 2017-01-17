package main

import (
    "syscall"
    "unsafe"
)

var (
    times = 0
    dll   = syscall.NewLazyDLL("user32.dll")
    proc  = dll.NewProc("MessageBoxW")
)

func main() {
    proc.Call(uintptr(0), 
                uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("test"))),
                uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("test message"))),
    uintptr(0))
}