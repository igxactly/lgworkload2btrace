package main

import (
	"bufio"
	"encoding/binary"
	"encoding/csv"
	// "encoding/hex"
	"fmt"
	"os"
	"strconv"
)

/*
func Uint64frombytes(bytes []byte) uint64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return bits
}

func Uint32frombytes(bytes []byte) uint32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return bits
}

func Uint16frombytes(bytes []byte) uint16 {
	bits := binary.LittleEndian.Uint16(bytes)
	return bits
}
*/

func Uint64bytes(bits uint64) []byte {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func Uint32bytes(bits uint32) []byte {
	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, bits)
	return bytes
}

func Uint16bytes(bits uint16) []byte {
	bytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(bytes, bits)
	return bytes
}

func main() {
	var magic uint32
	var seq uint32
	var time, sector uint64
	var bytes, action, pid, dev, cpu uint32
	var errcode, pdu_len uint16

	magic = 0x65617407
	seq = 0
	time = 0
	sector = 0
	bytes = 4096
	action = 0x9999999
	pid = 0x00007070
	dev = 0x10300000
	cpu = 0
	errcode = 0
	pdu_len = 0

	f_in, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_in.Close()

	buf_in := bufio.NewReader(f_in)
	csv_in := csv.NewReader(buf_in)

	f_out, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_out.Close()

	buf_out := bufio.NewWriter(f_out)

	var r []string
	r, err = csv_in.Read()

	seq = 0
	for r, err = csv_in.Read(); err == nil; r, err = csv_in.Read() {

		// set values
		time, _ = strconv.ParseUint(r[0], 10, 64)
		time *= 1000
		sector, _ = strconv.ParseUint(r[2], 10, 64)
		action = func() uint32 {
			switch r[1] {
			case "CMD18":
				return 0x00010001
			case "CMD25":
				return 0x00020001
			default: /* CMD05? */
				return 0x00000000
			}
		}()

		if action != 0 {
			var record []byte

			record = append(record, Uint32bytes(magic)...)
			record = append(record, Uint32bytes(seq)...)
			record = append(record, Uint64bytes(time)...)
			record = append(record, Uint64bytes(sector)...)
			record = append(record, Uint32bytes(bytes)...)
			record = append(record, Uint32bytes(action)...)
			record = append(record, Uint32bytes(pid)...)
			record = append(record, Uint32bytes(dev)...)
			record = append(record, Uint32bytes(cpu)...)
			record = append(record, Uint16bytes(errcode)...)
			record = append(record, Uint16bytes(pdu_len)...)

			_, err = buf_out.Write(record)
			seq++
		}
	}
	buf_out.Flush()
}
