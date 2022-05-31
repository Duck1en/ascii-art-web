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

for _, j := range strings.Split(str, "\n") { //run the output on each line
	if j == "" { // if the line is empty add a new line to the result
		res += "\n"
		continue
	}
	temp := [8]string{}
	for _, s := range j { // run through every character in the line
		for i, v := range strings.Split(ascii[s], "\n") { // split the big ASCII letter by new lines
			temp[i] += v // add to the temporary result line by line
		}
	}
	for _, s := range temp { // assemble the result
		res += s + "\n"
	} // repeat
}

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
