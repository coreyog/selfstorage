The stupidest thing I've ever programmed.

Build it. Using `go run .` is pointless. Everytime you run it, it will print an
incremented number.

In words:

The program gets the current permissions for the executable and saves them for
later. It reads the contents of it's own executable. If the last 4 bytes are not
"STRG" then the file is fresh. The program prints 0 and appends 8 bytes onto the
contents of itself: { 0, 0, 0, 0, 'S', 'T', 'R', 'G' }. The buffer is saved to
disk on the same path with the same permissions. On a second run, the program
will detect the last 4 bytes being "STRG" so it will read the 4 bytes prior to
it as a uint32, increment it, print it, and modify the last 8 bytes to use the
new value.

This means every time you run the program it reads and modifies itself to
increment a number in "storage" at the tail end of the binary. Because why not?