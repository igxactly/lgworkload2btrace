package main

import (
	"bufio"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	//"io"
	"os"
)

func Uint64frombytes(bytes []byte) uint64 {
	bits := binary.LittleEndian.Uint64(bytes)
	return bits
}

func Uint32frombytes(bytes []byte) uint32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return bits
}

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

func main() {
	fmt.Println("Hello")

	var magic uint32
	var seq uint32
	var time, sector uint64
	var bytes, action, pid, dev, cpu uint32
	var errcode, pdu_len uint16

	magic = 0x61657407
	seq = 1
	time = 2
	sector = 3
	bytes = 4096
	action = 5
	pid = 6
	dev = 0x10300000
	cpu = 0
	errcode = 7
	pdu_len = 0

	fmt.Printf("0x%08x\n", magic)
	fmt.Printf("0x%08x\n", seq)
	fmt.Printf("0x%016x\n", time)
	fmt.Printf("0x%016x\n", sector)
	fmt.Printf("0x%08x\n", bytes)
	fmt.Printf("0x%08x\n", action)
	fmt.Printf("0x%08x\n", pid)
	fmt.Printf("0x%08x\n", dev)
	fmt.Printf("0x%08x\n", cpu)
	fmt.Printf("0x%04x\n", errcode)
	fmt.Printf("0x%04x\n", pdu_len)

	fmt.Println("Now Little endian")

	le_bytes := Uint32bytes(magic)
	fmt.Println(hex.EncodeToString(le_bytes))

	le_bytes = Uint32bytes(dev)
	fmt.Println(hex.EncodeToString(le_bytes))
	/*
	   f_in, err := os.Open(os.Args[1])
	   if err != nil {
	       fmt.Println(err)
	       os.Exit(1)
	   }
	   defer f_in.Close()

	   reader := bufio.NewReader(f_in)
	*/

	//	/*
	f_out, err := os.OpenFile(os.Args[2], os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f_out.Close()
	writer := bufio.NewWriter(f_out)

	_, err = writer.Write(le_bytes)
	writer.Flush()
	//  */
}
