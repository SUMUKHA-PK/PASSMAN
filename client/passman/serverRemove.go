package passman

import (
	"fmt"

	"github.com/gookit/color"
)

// RemoveDataFromServer removes the data from the server
// It also ensures that any updated data in the server is not lost.
func RemoveDataFromServer() {
	color.Info.Println("PASSMAN remove data from server sequence")
	err := serverCommunication("remove")
	if err != nil {
		return
	}
	fmt.Printf("\n\n")
	color.Success.Println("Data removed from server.")
}
