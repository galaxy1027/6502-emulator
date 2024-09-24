package main

import (

)

type Memory struct {
	ram [65536] uint8 // 64k bytes of ram
}

func (mem *Memory) Write(address uint16, value uint8) {
	mem.ram[address] = value;
}

func (mem *Memory) Read(address uint16) uint8 {
	return mem.ram[address]
}

func (mem *Memory) Read16(address uint16) uint16 {
	var data uint16 = 0
	var lowByte uint16 = uint16(mem.Read(address))
	var highByte uint16 = uint16(mem.Read(address + 1)) << 8
	data = highByte | lowByte
	return data
}

func (mem *Memory) Write16(address uint16, value uint16) {
	var lowByte uint8 = uint8(value & 0x00FF)
	var highByte uint8 = uint8((value & 0xFF00) >> 8)
	mem.Write(address, lowByte)
	mem.Write(address + 1, highByte)
}
