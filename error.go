package shipstation

import "fmt"

func check(err error, origin string) {
	if err != nil {
		message := fmt.Sprintf("at %s -> detail\n%s", origin, err.Error())
		panic(message)
	}
}
