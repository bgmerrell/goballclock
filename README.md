SUMMARY
=======

Here's my software simulation of a ball-clock that outputs the number of
days (24-hour periods) which elapse before the clock returns to its initial
ordering.

Please note that this is not an optimal solution to the ball clock problem
(http://www.chilton.com/~jimw/ballclk.html).  This solution does, however,
fully simulate a ball clock.  That is, the balls move through the positions in
the clock until the balls are all back in their original order in the queue.

RUNNING THE CLOCK
=================

The program takes input from stdin, so after running the above command, simply
provide input (i.e., number of balls for the clock) and press enter.
Alternatively, you can process several inputs by doing something like this:

	cat clock-input.txt | goballclock > output.txt

Where clock-input.txt is a text file containing one input per line.

If you wanted to time the program, simply prefix the previous command with
"time ".

RUNNING THE TESTS
=================

If you'd like to run the unit tests, you can run the following from
src/ballclock:

	go test ./...

go1 appears to have a bug where it doesn't run the test in the current working
directory, so if you want to run all the tests from go1, just run the above
command in src (instead of src/ballclock).
