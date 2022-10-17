# Dups CLI

Remove duplicates takes a `--src` json file path and saves the output to a `--dest` file path (`clean_application.json` if not specified) with the ability to print out to stdout and run in concurrency mode if `-c true` was passed.

## How to Run

### Using Go
You can build the application if you have `go` installed in your machine. run the following to build and run:

1- Clone the application

```
git clone https://github.com/msolimans/dups
```

2- Build and run
```
go run ./cmd/main.go 
```

#### Examples

- Remove duplicates in non-concurrency mode `default mode` and saves output to `clean_application.json` since `-d` was not specified
```
go run cmd/main.go remove -s ./mock_application.json  -o     

Non-concurrent processing
---------- Summary ------------

Field count: 
         before: 31 
         after: 26
View count: 
         before: 10 
         after: 9
Done!
```

- Remove duplicates in concurrency mode and output summary
```
go run cmd/main.go remove -s ./mock_application.json -c -o 

Concurrent processing

 ---------- Summary ------------

Field count: 
         before: 31 
         after: 26
View count: 
         before: 10 
         after: 9
Done!
```

- Saves output into a file named `new.json`
```
go run cmd/main.go remove -s ./mock_application.json -d new.json 

Non-concurrent processing

Done!
```

For full list of commands and flags, please take a look at the below table.

### Using Docker

1- Clone the application

```
git clone https://github.com/msolimans/dups
```

2- Build docker image

```
docker build -t dups .
```

3- Run

```
docker run -it --rm dups 
```

## CLI commands and parameters:

| Command  |  Flag(s)  | Short Flag(s) |  Description             |           Required      |  Default
|----------|-----------|---------------|------------|-------------|-----|
| help     |  help     |    -h         | Prints help message               |  |
| version  |           |    -h         | Prints version details               |  |
| remove  |           |               | Remove duplicates               |  |
|          |  --src  |    -s         | source file path  | Y |  - |
|          |  --dest    |    -d         | destination file path | N |  clean_application.json |
|          |  --pretty   |    -p         | pretty format - enable indentation |  N  | true |
|          |  --concurrent   |    -c         | enable concurrency  |  N | false |
|          |  --out   |    -o         | output summary |  N | false |
