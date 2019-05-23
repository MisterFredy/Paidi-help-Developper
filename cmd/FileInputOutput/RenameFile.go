package FileInputOutput

import (
	"log"
	"os"
)

func RenameFile(framework string, nama string) {
	if framework == "codeigniter" {
		err := os.Rename("CodeIgniter-develop", nama)

		if err != nil {
			log.Print(err)
		}
	} else if framework == "laravel" {
		err := os.Rename("laravel-master", nama)

		if err != nil {
			log.Print(err)
		}
	}

}
