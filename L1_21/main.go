package main

import "fmt"

type Printer interface {
	Print(text string)
}

type OldPrinter struct{}

func (op *OldPrinter) PrintOld(text string) {
	fmt.Println("OldPrinter:", text)
}

type PrinterAdapter struct {
	oldPrinter *OldPrinter
}

func (pa *PrinterAdapter) Print(text string) {
	pa.oldPrinter.PrintOld(text)
}

func main() {
	oldPrinter := &OldPrinter{}
	adapter := &PrinterAdapter{oldPrinter: oldPrinter}

	var printer Printer = adapter

	printer.Print("Hello, Adapter Pattern!")
}
