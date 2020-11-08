package goton

import "testing"

func TestAbiMethods(t *testing.T) {
	config := NewConfig(0)
	client, err := InitClient(config)
	if err != nil {
		t.Errorf("test Failed - Init client error: %s", err)
	}
	defer client.Destroy()

}
