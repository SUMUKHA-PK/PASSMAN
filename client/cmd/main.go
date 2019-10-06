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
		options :=
			`1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
`
		color.Notice.Printf(options)
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
			passman.RemovePwd()
		case "5":
			passman.ChangeMasterPwd()
		case "6":
			passman.SyncDataWithServer()
		case "7":
			passman.ViewDataOnServer()
		case "8":
			passman.RemoveDataFromServer()
		default:
			fmt.Println("Please enter a valid option in the given list!")
		case "9":
			break
		}

		if option == "9" {
			color.Yellow.Println("Exiting PASSMAN. Bye :)")
			break
		}
	}

}
