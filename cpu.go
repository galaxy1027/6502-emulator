package main

import (
	"fmt";
)

type Cpu struct {
	mem *Memory; // Pointer to Memory (contains RAM / ROM)

	cycles uint32;

	PC uint16; // Program Counter
	SP uint16; // Stack Pointer
	A uint8; // Accumulator
	X uint8; // Index Register
	Y uint8; // Index Register

	/* Flags */
	n bool; // Negative
	v bool; // Overflow
	b bool; // Break
	d bool; // Decimal
	i bool; // Interrupt disable
	z bool; // Zero
	c bool; // Carry
}

func (cpu *Cpu) Reset() {
	cpu.mem = new(Memory) // Clear cpu memory
	cpu.PC = 0xFFFC // Reset vector address
	cpu.SP = 0x0100 // Set to beginning of stack memory
	cpu.i = true
	cpu.d = false
	cpu.b = true
	cpu.cycles = 0
	fmt.Println("CPU reset")
}

func (cpu *Cpu) Fetch() uint8 {
	var val = cpu.mem.Read(cpu.PC)
	cpu.PC++
	cpu.cycles++
	return val
}

func (cpu *Cpu) Execute(instruction uint8) {
	switch instruction {
	case 0xA9:
		cpu.INS_LDA_IM()
	default:
		fmt.Printf("ERROR: Unrecognized instruction: 0x%X\n", instruction)
	}
}

func (cpu *Cpu) INS_LDA_IM() {
	var value uint8 = cpu.Fetch()
	cpu.A = value
	cpu.z = (cpu.A == 0)
	cpu.n = (cpu.A & 0b10000000) != 0

	cpu.cycles += 2
}
