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

func getRunnerVersion() string {

	version := "*"

	output, err := exec.Command("/bin/sh", "-c", `ps $(pidof Runner.Listener) | grep -ioE "/.+"`).Output()
	if err != nil {
		return version
	}

	if len(output) == 0 {
		return version
	}
	// output = []byte("/home/runner/runners/2.313.0/bin/Runner.Listener run")
	outputParts := strings.Split(string(output), " ")
	if len(outputParts) < 2 {
		return version
	}

	runnerListenerPath := outputParts[0]
	pathParts := strings.Split(runnerListenerPath, "/")

	version = pathParts[4]

	return version
}

func main(){

  fmt.Println("hello world !")
	
  fmt.Printf("Ecapture Installed: %v\n", isEcaptureInstalled())
  fmt.Printf("IsEcaptureRunning: %v\n", isEcaptureRunning())
  fmt.Printf("RunnerVersion: %v\n", getRunnerVersion())
	
  
}
