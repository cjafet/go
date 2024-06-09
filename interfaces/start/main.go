package main

import (
	"fmt"

	"./parts"
)

func MonitorSummary(monitor parts.Monitor) string {
	return monitor.Specs() + "\n" + monitor.Price()
}

func HardDriveSummary(drive parts.HardDrive) string {
	return drive.Specs() + "\n" + drive.Price()
}

func KeyboardSummary(keyboard parts.Keyboard) string {
	return keyboard.Specs() + "\n" + keyboard.Price()
}

func main() {
	monitor := parts.Monitor{"HDMI", "1080p", 249.99}
	fmt.Println(MonitorSummary(monitor))
	fmt.Println("------------------------")
	drive := parts.HardDrive{"SSD", "SATA", 149.99}
	fmt.Println(HardDriveSummary(drive))
	fmt.Println("------------------------")
	keyboard := parts.Keyboard{108, "Mechanical", 99.99}
	fmt.Println(KeyboardSummary(keyboard))
	fmt.Println("------------------------")
}
