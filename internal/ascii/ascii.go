package ascii

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

var ErrNotFound = errors.New("font does not exist")

func AsciiDrawer(input string, font string) (string, int, error) {
	ascii, code, err := AssingFont(font)
	if err != nil {
		// log.Fatalf("Main: %v\n", err.Error())
		return "", code, err
	}

	input, code, err = CheckInput(input)
	if err != nil {
		// log.Fatalf("Main: %v\n", err.Error())
		return "", code, err
	}

	res := OutputAscii(input, ascii)
	return res, code, nil
}

// assigns a font to map
func AssingFont(str string) (map[rune]string, int, error) {
	str = "internal/fonts/" + str + ".txt"

	m := make(map[rune]string)
	content, err := ioutil.ReadFile(str)
	if err != nil {
		return nil, 404, fmt.Errorf("AssignFont: %v", err.Error())
	}

	//
	s := string(content)
	temp := ""
	count := 0
	j := ' '
	for _, v := range s {
		temp += string(v)
		if v == '\n' {
			count++
		}
		if count == 9 {
			m[j] = temp[1 : len(temp)-1]
			j++
			count = 0
			temp = ""
		}
	}
	// checks if all ascii symbols are in banner file
	if len(m) != 95 {
		return nil, 500, fmt.Errorf("AssingFont: Banner file corrupted")
	}
	return m, 200, nil
}

func OutputAscii(str string, ascii map[rune]string) string {
	res := ""
	for _, j := range strings.Split(str, "\n") {
		if j == "" {
			res += "\n"
			continue
		}
		temp := [8]string{}
		for _, s := range j {
			for i, v := range strings.Split(ascii[s], "\n") {
				temp[i] += v
			}
		}
		for _, s := range temp {
			res += s + "\n"
		}
	}

	return res[:len(res)-1]
}

func CheckInput(str string) (string, int, error) {
	str = strings.ReplaceAll(str, "\r\n", "\n")

	for _, s := range str {
		if s != '\n' && (s < 32 || s > 126) {
			return str, 400, fmt.Errorf("checkInputStr: incorrect symbols in %s", str)
		}
	}
	return str, 200, nil
}
