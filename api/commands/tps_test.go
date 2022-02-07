package commands

import (
	"os"
	"testing"
)

func TestTps(t *testing.T) {
	client := NewCommandClient(os.Getenv("API_KEY"))
	run, err := client.GetTps("D_L_1", true).Run()
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(run)
}
