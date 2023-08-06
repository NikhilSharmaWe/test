package main

import (
	"bytes"
	"encoding/binary"
)

func (n Name) ComposeNameBinary() []byte {
	b := new(bytes.Buffer)

	binary.Write(b, binary.LittleEndian, int32(len(n.Firstname)))

	binary.Write(b, binary.LittleEndian, []byte(n.Firstname))

	binary.Write(b, binary.LittleEndian, int32(len(n.Lastname)))

	binary.Write(b, binary.LittleEndian, []byte(n.Lastname))

	binary.Write(b, binary.LittleEndian, int32(n.Age))

	return b.Bytes()
}

func ParseNameBinary(b []byte) Name {
	n := Name{}

	var (
		fsLen int32
		lsLen int32
		age   int32
	)

	r := bytes.NewReader(b)
	binary.Read(r, binary.LittleEndian, &fsLen)

	fs := make([]byte, fsLen)
	binary.Read(r, binary.LittleEndian, &fs)

	binary.Read(r, binary.LittleEndian, &lsLen)

	ls := make([]byte, lsLen)
	binary.Read(r, binary.LittleEndian, &ls)

	binary.Read(r, binary.LittleEndian, &age)

	n.Firstname = string(fs)
	n.Lastname = string(ls)
	n.Age = int(age)

	return n
}
