<h1 align="center"> Go Converter 🔀 </h1>

<p align = "center">
  <kbd>
    <img src= "https://github.com/boydjawun/go-converter/blob/main/assets/converter.jpg" alt="Converter Image" width = "500" height = "500">
  </kbd>
</p>
  
> A small **Go CLI** that converts numbers between **decimal**, **hex**, and **binary**.

---

## What it does

- Accepts a decimal, hex (`-x`), or binary (`-b`) value from the command line
- Parses the input with Go’s standard library (`strconv`)
- Prints the same number in decimal, hex, and binary
- Uses only the Go standard library (no external dependencies)

---

## Usage
```
go run converter.go 
go run converter.go -x 
go run converter.go -b
```

## Example Output
``` 
Decimal: 255
Hex:     0xFF
Binary:  0b11111111
```


## Project Structure
```
go-converter/
├── converter.go     # CLI number converter
└── README.md
```

# How it Works

- Reads flags and arguments from `os.Args`
- Chooses the input base: `10` (default), `16` (`-x`), or `2` (`-b`)
- Parses the value with `strconv.ParseInt`
- Prints the number in decimal, hex (`0x`), and binary (`0b`)

## Build as an executable (optional)

1. Inside project folder 👉 ```go build -o converter converter.go```
2. Create a **folder** for the **<executable>.exe**, and move it in there
3. Then add the folder to your `PATH` if you want to run `converter` from anywhere by typing: ```converter <switch option> <value>```.

## Tech stack

| Layer | Technology |
|-------|------------|
| **Language** | Go |
| **CLI** | `os.Args` |
| **Parsing** | Go standard library (`strconv`) |
| **I/O** | Go standard library (`fmt`, `os`) |
| **Dependencies** | None (stdlib only) |
