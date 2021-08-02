package database

import (
	"testing"

	"github.com/adem522/tesodev/database"
)

func TestDeleteEmptyArg(t *testing.T) {
	err := database.Delete(nil, "Customer")
	if err == nil {
		t.Errorf("expected error, received err %v", err)
	}
}
