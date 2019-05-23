package FileInputOutput

import (
	"os"
)

func DeleteFile(framework string) error {
	// delete file
	err := os.Remove(framework + ".zip")

	if err != nil {
		return err
	}

	return err
}
