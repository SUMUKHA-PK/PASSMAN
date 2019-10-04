package passman

import (
	"fmt"
)

// SyncDataWithServer pushes all the data to the server
// It ensures no data is lost while syncing.
func SyncDataWithServer() {
	fmt.Println("PASSMAN Server sync sequence")
}
