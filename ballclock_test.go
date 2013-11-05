package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

const TESTDATADIR = "test/data"

// Run the ball clock using the contents of path as input.
// If validateInputOnly is true, the ball clock does not actually run.
// The output from running the clock is returned a string.
// An error is also returned, but is nil if there were no problems.
func runFromPath(t *testing.T, path string, validateInputOnly bool) (output string, err error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Fatalf("No test file: %s\n", path)
	}

	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("Failed to open test file: %s\n", path)
	}

	tempf, err := ioutil.TempFile(os.TempDir(), "ballclock")
	defer tempf.Close()
	defer os.Remove(tempf.Name())
	if err != nil {
		t.Fatalf("Failed to open temp file: %s\n", err.Error())
	}
	err = run(bufio.NewScanner(f), tempf, validateInputOnly)
	if !validateInputOnly {
		tempf.Seek(0, 0)
		bytes, err := ioutil.ReadFile(tempf.Name())
		if err != nil {
			t.Fatalf("Failed to read from temp file: %s\n", err.Error())
		}
		output = string(bytes)
	}
	return output, err
}

func TestGoodInputFile(t *testing.T) {
	path := filepath.Join(TESTDATADIR, "good-input-file.txt")
	output, err := runFromPath(t, path, false)
	if err != nil {
		t.Errorf("Unexpected failure parsing good input file (%s): %s\n", path, err.Error())
	}

	// validate output
	const EXPECTED1 = "30 balls cycle after 15 days."
	const EXPECTED2 = "45 balls cycle after 378 days."
	expected := fmt.Sprintf("%s\n%s\n", EXPECTED1, EXPECTED2)
	if output != expected {
		t.Errorf("Unexpected run output:\n"+
			"Actual: %s\n"+
			"Expected: %s",
			output,
			expected)
	}
}

func TestTooFewBalls(t *testing.T) {
	path := filepath.Join(TESTDATADIR, "input-file-too-few-balls.txt")
	_, err := runFromPath(t, path, true)
	if err == nil {
		t.Fatalf("Unexpected successful parsing bad input file (%s)", err.Error())
	}
	expected := "Malformed input (Too few balls, 26 < 27)"
	if err.Error() != expected {
		t.Errorf("Unexpected failure:\n"+
			"Actual: %s\n"+
			"Expected: %s",
			err.Error(),
			expected)
	}
}

func TestTooManyBalls(t *testing.T) {
	path := filepath.Join(TESTDATADIR, "input-file-too-many-balls.txt")
	_, err := runFromPath(t, path, true)
	if err == nil {
		t.Fatalf("Unexpected successful parsing bad input file (%s)", err.Error())
	}
	expected := "Malformed input (Too many balls, 128 > 127)"
	if err.Error() != expected {
		t.Errorf("Unexpected failure:\n"+
			"Actual: %s\n"+
			"Expected: %s",
			err.Error(),
			expected)
	}
}

func TestEmpty(t *testing.T) {
	path := filepath.Join(TESTDATADIR, "input-file-empty.txt")
	_, err := runFromPath(t, path, true)
	if err == nil {
		t.Fatalf("Unexpected successful parsing bad input file (%s)", err.Error())
	}
	expected := "Malformed input (empty)"
	if err.Error() != expected {
		t.Errorf("Unexpected failure:\n"+
			"Actual: %s\n"+
			"Expected: %s",
			err.Error(),
			expected)
	}
}

func TestValTooLarge(t *testing.T) {
	path := filepath.Join(TESTDATADIR, "input-file-val-too-large.txt")
	_, err := runFromPath(t, path, true)
	if err == nil {
		t.Fatalf("Unexpected successful parsing bad input file (%s)", err.Error())
	}
	expected := "Malformed input (failed to parse \"256\" as uint8)"
	if err.Error() != expected {
		t.Errorf("Unexpected failure:\n"+
			"Actual: %s\n"+
			"Expected: %s",
			err.Error(),
			expected)
	}
}

func TestNegativeVal(t *testing.T) {
	path := filepath.Join(TESTDATADIR, "input-file-negative-val.txt")
	_, err := runFromPath(t, path, true)
	if err == nil {
		t.Fatalf("Unexpected successful parsing bad input file (%s)", err.Error())
	}
	expected := "Malformed input (failed to parse \"-1\" as uint8)"
	if err.Error() != expected {
		t.Errorf("Unexpected failure:\n"+
			"Actual: %s\n"+
			"Expected: %s",
			err.Error(),
			expected)
	}
}
