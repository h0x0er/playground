package main

import "fmt"
import "os/exec"
import "os"
import "strings"

func isEcaptureRunning() bool {
	resp, err := exec.Command("/bin/sh","-c" ,"sudo pidof ecapture").Output()
	if err != nil {
		fmt.Printf("[Error] %v\n", err)
		return false
	}
	out := string(resp)
	fmt.Printf("out: %v\n", out)
	return len(strings.TrimSpace(out)) > 0

}

func isEcaptureInstalled() bool {

	if _, err := os.Stat("/usr/local/bin/ecapture"); os.IsNotExist(err) {
		return false
	}

	return true
}


func main(){

  fmt.Println("hello world !")
	
  fmt.Printf("Ecapture Installed: %v\n", isEcaptureInstalled())
  fmt.Printf("IsEcaptureRunning: %v\n", isEcaptureRunning())
  
}
