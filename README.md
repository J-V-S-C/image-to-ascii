
# Image to ASCII Converter (Terminal-Based)

This is a simple but powerful image-to-ASCII converter written in Go.  
It transforms image files into colorful or plain ASCII/ANSI representations directly in your terminal.

![Demo](./example.gif)

---

## Features

- Intelligent resizing and aspect ratio handling
- Multiple modes:
  - ANSI colored blocks
  - Colored ASCII
  - Plain grayscale ASCII
- Supports JPG, PNG
- Command-line interaction with mode selection

---


## Prerequisites

Make sure you have the following installed on your system:

- Go (https://golang.org/doc/install)
- Git
- A terminal that supports ANSI colors (like kitty, alacritty, wezterm, etc.)

For Arch Linux users, you can install Go with:

    sudo pacman -S go

---

## Installation

### 1. Clone this repository
```bash
    # Clone the directory
    git clone https://github.com/J-V-S-C/image-to-ascii.git
    cd image-to-ascii
```
### 2. (Optional) Initialize Go modules
```bash
#If you're using Go modules (recommended):
    go mod init main
    go mod tidy
```



### 3. Run the project
    go run main.go path/to/image.jpg
```bash
# example
    go run main.go astronaut-on-moon.jpg

```


You'll be prompted to choose a conversion mode:

    1 = ANSI
    2 = Colored ASCII
    3 = Plain ASCII

Choose an option and enjoy the output directly in your terminal!
