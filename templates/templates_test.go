package templates

import "testing"

func TestBuildCommand(t *testing.T) {
	input := "echo hello"
	expectedOutput := "echo hello\n"

	realizedOutput := BuildCommand(input)
	if realizedOutput != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, realizedOutput)
	}
}

func TestBuildCommandEmpty(t *testing.T) {
	input := ""
	expectedOutput := ""

	realizedOutput := BuildCommand(input)
	if realizedOutput != expectedOutput {
		t.Errorf("Expected %s, got %s", expectedOutput, realizedOutput)
	}
}
