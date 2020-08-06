# clipboard-image
Copy image to clipboard

## Supported OS
- Windows
- Mac
- Linux/Unix

## Requirements
- xclip(linux only)

## Usage
Copy image file to clipboard.

```go
f, err := os.Open("image.png")
if err != nil {
	log.Fatal(err)
}
deder f.Close()

if err := Rclipboard.CopyToClipboard(f); err != nil {
	log.Fatal(err)
}
```

Read image file from clipboard.

```go
r, err := Rclipboard.eadFromClipboard()
if err != nil {
	log.Fatal(err)
}

f, err := os.Create("image.png")
if err != nil {
	log.Fatal(err)
}
defer f.Close()

if _, err := io.Copy(f, r); err != nil {
	log.Fatal(err)
}
```

## Author
skanehira
