package uuid

import (
	"fmt"
	"testing"

	uuid "github.com/aeud/go-uuid"
)

// TestUUID generates a new UUID
func TestUUID(t *testing.T) {
	id, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	fmt.Println(id)
}
