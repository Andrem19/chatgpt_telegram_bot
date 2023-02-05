package helpers

import (
	"fmt"
	"os"
)


func AddToLog(log string) error {
	f, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("%s\n", log)); err != nil {
		panic(err)
	}
	return nil
}