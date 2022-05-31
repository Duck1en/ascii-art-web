# :duck: ascii-art-web :duck:

## Description

ASCII Art Web consists in creationg and running a server, in which it will be possible to use a web GUI version of ASCII Art project

## Authors

Rakhat (@Rshezarr) && Zhandos (@Duck1en)

## Usage

To start the server from root directory:

```
$ go run ./cmd/server.go
```

## Implementation details

```
Assign banner to a map with rune keys(' '-'~')

Replace "\r\n" with "\n"
check the input for non-printable characters

split the input by '\n'

loop over every line in array {
	if the line is empty {
        add a new line to the result
    }
    loop through every character in that line {
		loop through 8th new lines in given BIG LETTER {
			add the line to the temporary result
        }
    }
	assemble the result line by line from the temporary result
}

return the result to the server
```

```
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⡿⢋⣩⣭⣶⣶⣮⣭⡙⠿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠿⡙⣭⣮⣶⣶⣭⣩⢋⡿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⠿⣋⣴⣿⣿⣿⣿⣿⣿⣿⣿⣿⣦⡙⢿⣿⣿⣿⣿⣿⣿⣿⢿⡙⣦⣿⣿⣿⣿⣿⣿⣿⣿⣿⣴⣋⠿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⡃⠄⠹⡿⣿⣿⣿⣿⠟⠛⣿⣿⣿⣿⣷⡌⢿⣿⣿⣿⣿⣿⢿⡌⣷⣿⣿⣿⣿⠛⠟⣿⣿⣿⣿⡿⠹⠄⡃⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⠐⣠⡶⣶⣲⡎⢻⣿⣤⣴⣾⣿⣿⣿⣿⣿⠸⣿⣿⣿⣿⣿⠸⣿⣿⣿⣿⣿⣾⣴⣤⣿⢻⡎⣲⣶⡶⣠⠐⣿⣿⣿⣿⣿
⣿⠟⣋⡥⡶⣞⡯⣟⣾⣺⢽⡧⣥⣭⣉⢻⣿⣿⣿⣿⣿⣆⢻⣿⣿⣿⢻⣆⣿⣿⣿⣿⣿⢻⣉⣭⣥⡧⢽⣺⣾⣟⡯⣞⡶⡥⣋⠟⣿
⡃⣾⢯⢿⢽⣫⡯⣷⣳⢯⡯⠯⠷⠻⠞⣼⣿⣿⣿⣿⣿⣿⡌⣿⣿⣿⡌⣿⣿⣿⣿⣿⣿⣼⠞⠻⠷⠯⡯⢯⣳⣷⡯⣫⢽⢿⢯⣾⡃
⣦⣍⡙⠫⠛⠕⣋⡓⠭⣡⢶⠗⣡⣶⡝⣿⣿⣿⣿⣿⣿⣿⣧⢹⣿⢹⣧⣿⣿⣿⣿⣿⣿⣿⡝⣶⣡⠗⢶⣡⠭⡓⣋⠕⠛⠫⡙⣍⣦
⣿⣿⣿⣿⣿⣿⣘⣛⣋⣡⣵⣾⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣾⣵⣡⣋⣛⣘⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿
```
