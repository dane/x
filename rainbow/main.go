package main

import (
	"bytes"
	"compress/zlib"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

/*
|1| Frame Type
|1| Version
|1| Control Type
| Flags (8) |
| Header Length (8 byte)|
| Headers compressed|
| Content-Length (24 byte)|
*/

func main() {
	var (
		Type_CONTROL byte   = 0
		Version      uint16 = 1

		Control_TBD byte = 0

		// Note: may need 0 to be a nil flag. See what SPDY does
		Flags_FIN   byte = 0
		Flags_MULTI byte = 1
	)

	var (
		buf bytes.Buffer
		hdr [8]byte
	)

	hdr[0] = Flags_FIN
	hdr[1] = Flags_MULTI

	buf.WriteByte(Type_CONTROL)
	binary.Write(&buf, binary.LittleEndian, Version)
	buf.WriteByte(Control_TBD)
	binary.Write(&buf, binary.LittleEndian, hdr)

	frameType, err := buf.ReadByte()
	errorIf(err)
	fmt.Printf("Frame Type: %d\n", frameType)

	var version uint16
	err = binary.Read(&buf, binary.LittleEndian, &version)
	errorIf(err)
	fmt.Printf("Version: %d\n", version)

	controlType, err := buf.ReadByte()
	errorIf(err)
	fmt.Printf("Control Type: %d\n", controlType)

	var headers [8]byte
	err = binary.Read(&buf, binary.LittleEndian, &headers)
	errorIf(err)
	fmt.Printf("Headers: %d\n", headers)

	fmt.Printf("Len: %d\n", buf.Len())

	f, err := os.Open("raw.json")
	errorIf(err)

	var compressed bytes.Buffer
	w := zlib.NewWriter(&compressed)
	size, err := io.Copy(w, f)
	errorIf(err)

	fmt.Printf("Compressed size: %d", size)
}

func errorIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
