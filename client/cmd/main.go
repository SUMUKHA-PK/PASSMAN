package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SUMUKHA-PK/Password-Manager/client/passman"
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
	fmt.Println(welcome)

	for {
		fmt.Printf("\nWhat do you want to do?\n\n")
		fmt.Printf("1. Register on PASSMAN.\n2. View saved passwords.\n3. Add password.\n4. Change master password.\n5. Sync data with server\n6. Remove data from server\n7. Exit PASSMAN :(\n\n")
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
			passman.RemoveDataFromServer()
		default:
			fmt.Println("Please enter a valid option in the given list!")
		case "7":
			break
		}

		if option == "7" {
			fmt.Println("Exiting PASSMAN. Bye :)")
			break
		}
	}

}
