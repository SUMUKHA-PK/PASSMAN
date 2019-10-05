package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SUMUKHA-PK/PASSMAN/client/passman"
	"github.com/gookit/color"
)

func main() {

	welcome := `
+-----+ +----+ +----- +----- |\    /| +----+ |\    |
|     |	|    | |      |      | \  / | |    | | \   |
+-----+	|----| +----+ +----+ |  \/  | |----| |  \  |
|       |    |      |      | |      | |    | |   \ |
|       |    | -----+ -----+ |      | |    | |    \|

Welcome to PASSMAN, your locally hosted Password Manager!

`
	color.Yellow.Println(welcome)

	for {
		color.Notice.Printf("\n\nWhat do you want to do?\n\n")
		color.Notice.Printf("1. Register on PASSMAN.\n2. View saved passwords.\n3. Add password.\n4. Change master password.\n5. Sync data with server\n6. View data on server\n7. Remove data from server\n8. Exit PASSMAN :(\n\n")
		reader := bufio.NewReader(os.Stdin)

		option, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Error in reading input: %v", err)
		}
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			passman.Register()
		case "2":
			passman.Access()
		case "3":
			passman.AddPwd()
		case "4":
			passman.ChangeMasterPwd()
		case "5":
			passman.SyncDataWithServer()
		case "6":
			passman.ViewDataOnServer()
		case "7":
			passman.RemoveDataFromServer()
		default:
			fmt.Println("Please enter a valid option in the given list!")
		case "8":
			break
		}

		if option == "8" {
			color.Yellow.Println("Exiting PASSMAN. Bye :)")
			break
		}
	}

}
