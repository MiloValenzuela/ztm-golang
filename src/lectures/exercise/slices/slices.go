//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

func printAssemblyLine(line []Part) {
	fmt.Println("Assembly line Contents:")
	for i := 0; i < len(line); i++ {
		part := line[i]
		fmt.Printf("%d: %s\n", i+1, part)
	}
	fmt.Println()
}

type Part string

func main() {
	//  - Create an assembly line having any three parts
	assemblyLine := []Part{"Bolt", "Nut", "Washer"}

	//  - Print the contents of the assembly line
	printAssemblyLine(assemblyLine)

	//  - Add two new parts to the line
	assemblyLine = append(assemblyLine, "Gear", "Screw")

	printAssemblyLine(assemblyLine)

	//  - Slice the assembly line so it contains only the two new parts
	assemblyLine = assemblyLine[len(assemblyLine)-2:]

	printAssemblyLine(assemblyLine)
}
