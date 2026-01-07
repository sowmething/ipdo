# ipdo

**ipdo** ( *Internet Protocol do* ) is a minimal, fast, and no-nonsense IP reset utility for **Windows**, written in **Go**.

It simply releases and renews the system IP address using native Windows networking commands, wrapped in a clean CLI tool.

---

## Features

*  Fast IP release & renew (`ipconfig /release` → `ipconfig /renew`)
*  Windows-only (native behavior)
*  Detects administrator privileges
*  `--silent` mode for scripting / automation
*  Colored output & ASCII banner (non-silent mode)
*  Clean exit codes

---

## Usage

### Normal mode

```bash
ipdo.exe
```

Output includes:

* ASCII banner
* Status messages
* Warnings if not ran as Administrator

---

### Silent mode

```bash
ipdo.exe --silent
```

Silent mode:

* No banner
* No colored logs
* Operation completed successfully.

Useful for:

* Scripts
* Scheduled tasks
* Automation pipelines

---

##  Administrator Rights

Running as **Administrator** is **recommended**.

Without admin rights:

* IP release/renew **may** fail
* Network adapters **may** not reset correctly

ipdo will **warn**, but will not exit.

---

## 🛠 Build

### Requirements

* Go 1.20+
* Windows (amd64)

### Build (64-bit, stripped)
- Clone the project and

```bash
go build -ldflags="-s -w" -o ipdo.exe
```
---

## Project Structure

```
ipdo/
├─ main.go        
├─ banner.go      
├─ go.mod
└─ README.md
```

---

## Notes

* ipdo does **not** change MAC addresses
* ipdo does **not** spoof traffic
* ipdo only resets the IP via standard Windows networking

This tool is intended for:

* Development
* Testing
* Network refresh

---

## Author

Made by **sowmething**
GitHub: [https://github.com/sowmething](https://github.com/sowmething)

---
