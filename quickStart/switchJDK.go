package main

// Computer\HKEY_LOCAL_MACHINE\SYSTEM\ControlSet001\Control\Session Manager\Environment
// https://stackoverflow.com/questions/46542609/create-a-registry-key-string-with-go

import (
    //"golang.org/x/sys/windows/registry"
    //"log"
    "fmt"
    "os"
    "flag"
)

const JAVA_HOME = "JAVA_HOME"
const JDK8 = "C:\\Program Files\\Java\\jdk8"
const JDK11 = "C:\\Program Files\\Java\\jdk11"

func main() {
    intPtr := flag.Int("jdk", 11, "an int")
    flag.Parse()
    if *intPtr == 11 {
        os.Setenv(JAVA_HOME, JDK11)
    } else {
        os.Setenv(JAVA_HOME, JDK8)
    }
    fmt.Println(JAVA_HOME + ":", os.Getenv("JAVA_HOME"))
}