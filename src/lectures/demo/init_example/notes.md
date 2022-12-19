# init Function
- Common to have an initialization step in programs
    - Creating a function named `init()` will do
    - `init()` runs **before** the `main()` function
- Allows creation and creation and validation of program state before executio begins
    e.g.)
    - Check network connections
    - Check db connections
    - Cache expensive operations
    - etc.

# Further Info on init Function
- Each package can have its own `init()` function
- All packages will execute `init()` before `main()` runs
- even imported multiple times, `init()` will run only once inside its corresponding package