# goloader - Offensive Loader, Obfuscator and Bypass Tool written in GoLang
Here you find the basic informations about the loader.

**IMPORTANT:** This project is Work-in-Progress - WIP.
It is my first project regarding malware development.


## Evasion Modules
The evasion modules provides the following bypasses:
 - AMSIBypass(processName) -> Bypass AMSI based on defined Memory Pattern so search for

Flags to use:
```bash
$ goloader --byAMSI
```

## Obfuscator Modules
The current obfuscator provides the possiblity to obfuscatea/deobfuscate 
the payload to/from IPv4 format to hide it from the basic detection 
mechanisms. 

Flags to use:
```bash
# Obfuscate given payload
$ goloader --payload <PAYLOAD> --toIPv4

# Deobfuscate given payload
$ goloader --payloadFile <PATH-TO-FILE> --fromIPv4
```

## SysCall Modules
The syscall module provides one function to excute system commands
based on current operating system (Windows, UNIX, Mac)

Flags to use:
```bash
# Basic execution of given payload
$ goloader --payload <COMMAND/PAYLOAD>

# Execute payload from file 
$ goloader --payloadFile <PATH-TO-FILE>

# Execute payload from file and deobfuscate it from IPv4 format
$ goloader --payloadFile <PATH-TO-FILE> --fromIPv4
```


## Helper Modules
The helper module provides the following helper functions:
 - GetContentFromFileWithChecks(filePath) -> Returns the whole file content 

More helpers will added by need.
