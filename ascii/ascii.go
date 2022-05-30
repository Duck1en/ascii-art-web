package asciiart

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func AsciiDrawer(input string, font string) (string, error) {
	ascii, err := AssingFont(font)
	if err != nil {
		// log.Fatalf("Main: %v\n", err.Error())
		return "", err
	}

	input, err = CheckInput(input)
	if err != nil {
		// log.Fatalf("Main: %v\n", err.Error())
		return "", err
	}

	res := OutputAscii(input, ascii)
	return res, nil
}

// assigns a font to map
func AssingFont(str string) (map[rune]string, error) {
	str = "fonts/" + str

	m := make(map[rune]string)
	content, err := ioutil.ReadFile(str)
	if err != nil {
		return nil, fmt.Errorf("AssignFont: %v", err.Error())
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
		return nil, fmt.Errorf("AssingFont: Banner file corrupted")
	}
	return m, nil
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

func CheckInput(str string) (string, error) {
	str = strings.ReplaceAll(str, "\r\n", "\n")

	for _, s := range str {
		if s != '\n' && (s < 32 || s > 126) {
			return str, fmt.Errorf("checkInputStr: incorrect symbols in %s", str)
		}
	}
	return str, nil
}
