# Json Comparator

This comparator was built to compare json files exported from databases.

The top format is an array with json objects. The order of these objects and
their keys may be in any order and the comparator may still identify them as
similar.

Examples of the format can be seen in the `comparator/test_examples/` folder.


## Project Structure
```md
|cmd              - Folder with the commands which will turn into executables
|--/comparator    - Command that runs the comparator code
|comparator       - Main json similarity comparator functionality
|--/test_examples - Files used as examples to test the comparator functionality
```


## Build
To compile the program run
```sh
make
```


## Run
The program can be run with
```sh
bin/comparator <file1.json> <file2.json>
```


## Test
To run the tests to make sure the program works as expected run
```sh
make test
```

