# :duck: ascii-art-web :duck:

## Description
ASCII Art Web consists in creationg and running a server, in which it will be possible to use a web GUI version of ASCII Art project

## Authors

Rakhat && Zhandos 
@Rshezarr @Duck1en
## Usage
To start the server from root directory:
```
$ go run ./cmd/server.go
```

## Implementation details
<!-- Initialize a map with rune keys.
Assing every 10 new lines to a key from 32 to 126 runes removing top and bottom lines.
 -->


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
⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⢸⣿⣿⣿⣿⣿⣿⣿⢸⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿s
```
