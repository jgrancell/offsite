package app

import "testing"

func TestLoad(t *testing.T) {
	a := &App{}
	a.Load(nil, "0.0.1")
	if a.Version != "0.0.1" {
		t.Errorf("expected version to equal 0.0.1, got %s", a.Version)
	}
}

func TestParseFlags(t *testing.T) {
	args := make([]string, 0)
	args = append(args, "-h")
	args = append(args, "--foo")
	args = append(args, "bar")
	args = append(args, "--fizz-buzz")
	args = append(args, "baz")
	args = append(args, "mycommand")

	a := App{}
	err := a.ParseFlags(args)
	if err != nil {
		t.Errorf("expected nil result, got error: %s", err.Error())
	}

	// Check help flag
	checkFlag := a.Flags["h"]
	if checkFlag.Type != "boolean" {
		t.Errorf("expected h flag to be set to type boolean, got %s", checkFlag.Type)
	}

	if checkFlag.Value.(bool) != true {
		t.Errorf("expected h flag to be set to value true, got %v", checkFlag.Value.(bool))
	}

	// Check foo flag
	checkFlag = a.Flags["foo"]
	if checkFlag.Type != "assignment" {
		t.Errorf("expected foo flag to be set to type assignment, got %s", checkFlag.Type)
	}

	if checkFlag.Value.(string) != "bar" {
		t.Errorf("expected foo flag to be set to value bar, got %s", checkFlag.Value.(string))
	}

	// Check fizz-buzz flag
	checkFlag = a.Flags["fizz-buzz"]
	if checkFlag.Type != "assignment" {
		t.Errorf("expected fizzbuzz flag to be set to type assignment, got %s", checkFlag.Type)
	}

	if checkFlag.Value.(string) != "baz" {
		t.Errorf("expected foo flag to be set to value baz, got %s", checkFlag.Value.(string))
	}
}

func ConvertBoolToInterface(foo bool) interface{} {
	return foo
}
