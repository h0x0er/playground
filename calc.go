package main

import "fmt"
import "os/exec"

func isEcaptureRunning() bool {
	resp, err := exec.Command("/bin/sh", "sudo pidof ecapture").Output()
	if err != nil {
		return false
	}
	out := string(resp)
	// fmt.Printf("out: %v\n", out)
	return len(strings.TrimSpace(out)) > 0

}

func main(){

  fmt.Println("hello world !")
  fmt.Println("IsEcaptureRunning: %v\n", isEcaptureRunning())
  
}
