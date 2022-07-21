# Get Host Data
---
## Structure and technologies used

### What purpose?
> Small command-line app, for the purpose of studying and creating a host capture information tool.
> Initially capturing information of the types: IP, Nameserver, MX and TXT.

### What was used?
- [Golang](https://go.dev/)
- [package cli](https://pkg.go.dev/github.com/urfave/cli@v1.22.9)
- [package net](https://pkg.go.dev/net)

### Basic structure:
- [package generate](/generate)
- [package options](/options)

### Basic usage:
> Linux:
> 
> ``` ./gethostdata <command> --host <your.host>```
![](/images/gethostdata_use_linux.png)

> Windows
> 
> ``` ./gethostdata.exe <command> --host <your.host>```
![](/images/gethostdata_use_windows.png)

---
---
_end of README.md_