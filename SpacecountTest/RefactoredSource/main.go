package main

import (
	"fmt"

	types "github.com/fuatto/UnitTestsGolang/SpacecountTest/RefactoredSource/type"
)

func main() {
	liveClient := types.LiveGetWebRequest{}
	number := types.GetAstronauts(liveClient)

	fmt.Println(number)
}
