# Go-Do

Go-Do is a task management CLI utility that is designed to help you keep track of
incoming tasks that are due during any period of time.

Set priorities and tags to help you sort and filter task respectively.

Create recurring tasks that need to be done multiple times or forever.

## Installation

### 1. Install Go

You need the Golang toolchain installed in your system, minimum version required is v1.25.0.

I recommend you install Go through the [Webi installer](https://webinstall.dev/golang/) for ease of use but you could
also use installers from the [official page](https://go.dev/) or your system's package manager.

### 2. Install the Go-Do CLI

Use the following command to download and install the Go-Do CLI in your Go's toolchain `bin` directory.

```bash
go install github.com/JorgeToAn/go-do@latest
```

To verify if the CLI was installed corrrectly, use `go-do version`.
