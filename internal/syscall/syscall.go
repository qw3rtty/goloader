// Package syscall provides functions to make system calls.
package syscall

import (
    "fmt"
    "strings"
    "os/exec"
    "bytes"
)

func ExecuteSystemCall(command string) bool {
    // Check if command exists
    if len(command) == 0 {
        return false
    }

    // Split command on whintespace
    args := strings.Fields(command)
    fmt.Println("[i] Execute command:", args)
    fmt.Println()

    // Call command dynamiclly with n arguments
    cmd := exec.Command(args[0], args[1:]...)

    // Create buffers to capture stdout and stderr
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr

    // Execute command
    cmd.Run()
    fmt.Println(stdout.String())
    fmt.Println(stderr.String())

    return true
}
