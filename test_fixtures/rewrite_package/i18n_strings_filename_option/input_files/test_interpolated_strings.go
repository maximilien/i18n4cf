package input_files

import (
	"fmt"
)

func Interpolated() string {
	name := "cruel"
	myName := "evil"
	fmt.Println("Hello %s world!", name)
	fmt.Println("Hello %s world!, bye from %s", name, myName)

	fmt.Println("Hello %d(%s) world!, bye from %s", 10, name, "Evil")

	args := []string{"oh no", "really?", "that's awful!!!"}
	fmt.Println("%s %s %s", args...)
}
