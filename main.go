package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func main() {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}

	nfo, err := os.Stat(path)
	if err != nil {
		panic(err)
	}

	raw, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	hasStorage := bytes.HasSuffix(raw, []byte("STRG"))

	if !hasStorage {
		fmt.Println(0)

		buf := &bytes.Buffer{}
		err = binary.Write(buf, binary.LittleEndian, uint32(0))
		if err != nil {
			panic(err)
		}

		err = binary.Write(buf, binary.LittleEndian, []byte("STRG"))
		if err != nil {
			panic(err)
		}

		raw = append(raw, buf.Bytes()...)

		err = os.Remove(path)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(path, raw, nfo.Mode())
		if err != nil {
			panic(err)
		}
	} else {
		rawNum := raw[len(raw)-8 : len(raw)-4]
		num := binary.LittleEndian.Uint32(rawNum)
		num++

		fmt.Println(num)

		raw = raw[:len(raw)-8]

		buf := &bytes.Buffer{}
		err = binary.Write(buf, binary.LittleEndian, num)
		if err != nil {
			panic(err)
		}

		err = binary.Write(buf, binary.LittleEndian, []byte("STRG"))
		if err != nil {
			panic(err)
		}

		raw = append(raw, buf.Bytes()...)

		err = os.Remove(path)
		if err != nil {
			panic(err)
		}

		err = os.WriteFile(path, raw, nfo.Mode())
		if err != nil {
			panic(err)
		}
	}
}
