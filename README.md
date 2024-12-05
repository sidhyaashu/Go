## Init

```bash
go mod init <project_name>
```
```bash
go run ./main.go
```
Here’s a `README.md` file that you can use for your Go programming language tutorial project on GitHub, tailored for working on Windows 11 using VS Code:

```markdown
# Go Programming Language Tutorial

This repository contains a step-by-step tutorial to help you get started with Go programming language. The tutorial covers everything from installation to writing basic programs, data types, control structures, functions, and much more.

## Prerequisites

To follow this tutorial, you'll need the following:

- **Windows 11** (but should work on other operating systems too)
- **Visual Studio Code (VS Code)** installed
- **Go Programming Language** installed on your machine

### Installation Steps

#### 1. Install Go

1. Go to the official [Go download page](https://go.dev/dl/).
2. Download the Go installer for Windows.
3. Run the installer and follow the prompts to install Go.

#### 2. Verify Go Installation

After installation, open **Command Prompt** or **PowerShell**, and check the Go version to verify that the installation was successful.

```bash
go version
```
```

This should display the installed version of Go, like so:
```
go version go1.19.5 windows/amd64
```

#### 3. Install Visual Studio Code (VS Code)

1. Go to the official [Visual Studio Code download page](https://code.visualstudio.com/).
2. Download and install the Windows version of VS Code.
3. Once installed, open VS Code.

#### 4. Install Go Extension for VS Code

1. Open VS Code.
2. Go to the **Extensions** view by clicking on the Extensions icon on the Activity Bar on the side of the window.
3. Search for **Go** and install the official Go extension by **Go Team at Google**.

---

## Setting Up Your Go Workspace

1. **Set the Go Path**  
   Open **Command Prompt** or **PowerShell**, and set the `GOPATH` environment variable:

   ```bash
   setx GOPATH "%USERPROFILE%\go"
   setx PATH "%PATH%;%GOPATH%\bin"
   ```

2. **Create a Project Folder**  
   Navigate to your workspace directory (e.g., `%USERPROFILE%\go\src`), and create a new folder for your project:

   ```bash
   mkdir %USERPROFILE%\go\src\hello
   cd %USERPROFILE%\go\src\hello
   ```

3. **Create Your First Go File**  
   Inside your project folder (`hello`), create a file named `main.go`. You can do this from within VS Code or using the command line.

---

## Writing Your First Go Program

1. Open `main.go` in VS Code, and add the following code:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

2. **Run the Program**

In your **Command Prompt** or **PowerShell**, navigate to the project folder (`hello`) and run:

```bash
go run main.go
```

You should see the output:
```
Hello, World!
```

---

## Tutorial Overview

### 1. **Variables and Constants**

In Go, variables can be declared using `var` or shorthand `:=`. Constants are declared using the `const` keyword.

```go
package main

import "fmt"

func main() {
    var x int = 10
    y := 20
    const z = 30

    fmt.Println(x, y, z)
}
```

### 2. **Control Structures**

Go supports standard control structures like `if`, `for`, and `switch`.

```go
// if-else
if x > 10 {
    fmt.Println("x is greater than 10")
} else {
    fmt.Println("x is 10 or less")
}

// for loop
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

### 3. **Functions**

Functions in Go are defined using the `func` keyword. Here's an example:

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(5, 3)
    fmt.Println(result)
}
```

### 4. **Structs and Interfaces**

Go supports struct-based object-oriented programming. Here's an example of a struct and an interface:

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

type Shape interface {
    Area() float64
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

func main() {
    p := Person{Name: "John", Age: 30}
    fmt.Println(p)

    var s Shape = Circle{Radius: 5}
    fmt.Println("Area:", s.Area())
}
```

---

## Build and Install Your Program

To build your Go program into an executable file:

1. Run the following command in your project directory:

```bash
go build
```

This will generate an executable file in your project directory (e.g., `hello.exe` on Windows).

2. To run the built executable:

```bash
./hello
```

---

## Additional Resources

- [Go Documentation](https://go.dev/doc/)
- [Go Playground](https://go.dev/play/)
- [Go Tour](https://tour.golang.org/)

---

## Conclusion

This repository provides the basic steps to get started with Go. From here, you can explore more advanced topics such as concurrency with goroutines, building APIs, and working with Go modules.

Feel free to fork this repository and contribute any improvements or additions to the tutorial!

---

