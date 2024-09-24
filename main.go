package main

func main() {
	var cpu *Cpu = new(Cpu)
	cpu.Reset()

	var running bool = true
	var timer int = 0

	for running {
		timer++
		running = (timer < 5) // Limit how long emulator runs, delete later
		var instruction = cpu.Fetch()
		cpu.Execute(instruction)
	}

}
