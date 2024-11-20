package main

import (
	"fmt"
	oops "goPrac/Oops"
	solidprinciples "goPrac/SolidPrinciples"
)

func main() {
	fmt.Println("Main File Call")
	solidprinciples.RunSingleResponsibility()
	solidprinciples.RunOpenClosePrinciple()
	solidprinciples.RunLiskovSubstituion()
	solidprinciples.RunInterfaceSegregation()
	solidprinciples.RunDependencyInversionPrinciple()

	oops.RunEncapsulation()
	oops.RunGoRoutineAndChannel()
	oops.RunMutex()
}
