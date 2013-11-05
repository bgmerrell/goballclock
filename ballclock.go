/*
ballclocks's main package

Parse command line arguments and let the fun begin!
*/
package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"github.com/bgmerrell/goballclock/clock"
	"os"
	"path"
	"strconv"
)

const NARGS = 0
const MAXBALLS = 127
const MINBALLS = 27
const END_OF_INPUT_VAL = 0

func usage() {
	name := path.Base(os.Args[0])
	msg := fmt.Sprintf("Usage: %s\n\n"+
		"%s takes no arguments and accepts input from stdin.\n", name, name)
	fmt.Fprintf(os.Stderr, msg)
}

func parseCommandLine() {
	flag.Parse()
}

// Take a bufio Scanner and parse scanned input.
// An error is returned if there is a problem parsing the input.
func run(scanner *bufio.Scanner, file *os.File, validateInputOnly bool) error {
	// Only need uint8, but strconv.ParseUint returns a uint64.
	var nBalls uint64
	var err error

	for scanner.Scan() {
		// parsed value is base 10 and should fit within 8 bits
		text := scanner.Text()
		if nBalls, err = strconv.ParseUint(text, 10, 8); err != nil {
			msg := fmt.Sprintf("Malformed input (failed to parse \"%s\" as uint8)", text)
			fmt.Fprintf(os.Stderr, msg)
			return errors.New(msg)
		}
		if nBalls == END_OF_INPUT_VAL {
			return nil
		} else if nBalls > MAXBALLS {
			msg := fmt.Sprintf("Malformed input (Too many balls, %d > %d)", nBalls, MAXBALLS)
			fmt.Fprintln(os.Stderr, msg)
			return errors.New(msg)
		} else if nBalls < MINBALLS {
			msg := fmt.Sprintf("Malformed input (Too few balls, %d < %d)", nBalls, MINBALLS)
			fmt.Fprintln(os.Stderr, msg)
			return errors.New(msg)
		} else {
			if !validateInputOnly {
				fmt.Fprintf(file, "%d balls cycle after %d days.\n",
					nBalls,
					clock.GetDaysUntilCycle(uint8(nBalls)))
			}
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from input:", err.Error())
		return err
	} else if nBalls == 0 {
		msg := "Malformed input (empty)"
		fmt.Fprintln(os.Stderr, msg)
		return errors.New(msg)
	} else if nBalls != 0 {
		msg := fmt.Sprintf("Malformed input (zero should signify the end of input, got %d)", nBalls)
		fmt.Fprintln(os.Stderr, msg)
		return errors.New(msg)
	}
	return nil
}

func main() {
	flag.Usage = usage
	parseCommandLine()
	if flag.NArg() != NARGS {
		usage()
		os.Exit(1)
	}

	// The input may be of an unspecified length, so we'll use buffered IO
	// and compute the ball cycles as we receive input
	if err := run(bufio.NewScanner(os.Stdin), os.Stdout, false); err != nil {
		os.Exit(1)
	}
}
