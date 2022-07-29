package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Oooh, you have to run the programm with 2 arguments, 1st argument is 'sample.txt' and the 2nd argument is 'result.txt' JUST DO IT!")
		return
	}
	if os.Args[1] != "sample.txt" || os.Args[2] != "result.txt" {
		fmt.Println("Hey? Wrong input files :(")
		return
	}
	content, err := ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("The system cannot find the file specifiied, please create file with name 'sample.txt' JUST TRY AGAIN :)")
		return
	}
	if len(string(content)) == 0 {
		fmt.Println("Oooh, file is empty :(")
		return
	}
	vowels := []rune{'a', 'e', 'y', 'u', 'i', 'o', 'h'}
	arr := EmptySpaces(strings.Split(string(content), " "))

	for k := 0; k <= len(arr); k++ {
		for i, key := range arr {
			if key == "(cap)" {
				if i == 0 {
					arr = CheckFirstKey(i, arr)
				} else {
					arr = Cap(i, arr[i-1], arr)
				}
				break
			}
			if key == "(up)" {
				if i == 0 {
					arr = CheckFirstKey(i, arr)
				} else {
					arr = Up(i, arr[i-1], arr)
				}
				break
			}
			if key == "(low)" {
				if i == 0 {
					arr = CheckFirstKey(i, arr)
				} else {
					arr = Low(i, arr[i-1], arr)
				}
				break
			}
			if key == "(hex)" {
				if i == 0 {
					arr = CheckFirstKey(i, arr)
				} else {
					arr = Hex(i, arr[i-1], arr)
				}
				break
			}
			if key == "(bin)" {
				if i == 0 {
					arr = CheckFirstKey(i, arr)
				} else {
					arr = Bin(i, arr[i-1], arr)
				}
				break
			}
			if key == "(low," {
				if i == 0 {
					arr = CheckFirstKey2(i, arr)
				} else {
					arr = LowInt(i, arr)
				}
				break
			}
			if key == "(up," {
				if i == 0 {
					arr = CheckFirstKey2(i, arr)
				} else {
					arr = UpInt(i, arr)
				}
				break
			}
			if key == "(cap," {
				if i == 0 {
					arr = CheckFirstKey2(i, arr)
				} else {
					arr = CapInt(i, arr)
				}
				break
			}
			if key == "A" || key == "a" {
				for _, w := range vowels {
					if rune(arr[i+1][0]) == w {
						replace := An(i, arr[i], arr)
						arr = replace
						break
					}
				}
			}
			if key == "An" || key == "an" || key == "AN" || key == "aN" {
				for _, w := range vowels {
					if rune(arr[i+1][0]) == w {
						replace := AnReverse(i, arr[i], arr)
						arr = replace
						break
					}
				}
			}
		}
	}

	arr = SignFormatter(arr)
	result := strings.Join(arr, " ")
	txt := os.WriteFile(args[1], []byte(result), 0644)
	if txt != nil {
		fmt.Println("error: Cannot to create final file")
	}
}

func EmptySpaces(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

func CheckFirstKey(i int, arr []string) []string {
	var newArr []string
	for i := 1; i < len(arr); i++ {
		newArr = append(newArr, arr[i])
	}
	return newArr
}

func CheckFirstKey2(i int, arr []string) []string {
	var newArr []string
	for i := 2; i < len(arr); i++ {
		newArr = append(newArr, arr[i])
	}
	return newArr
}

func Cap(i int, str string, arr []string) []string {
	arr[i-1] = strings.Title(strings.ToLower(str))
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}

func Up(i int, str string, arr []string) []string {
	arr[i-1] = strings.ToUpper(str)
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}

func Low(i int, str string, arr []string) []string {
	arr[i-1] = strings.ToLower(str)
	arr = append(arr[:i], arr[i+1:]...)
	return arr
}

func Hex(i int, str string, arr []string) []string {
	decimal, err := strconv.ParseInt(str, 16, 32)
	if err == nil {
		arr[i-1] = strconv.Itoa(int(decimal))
		arr = append(arr[:i], arr[i+1:]...)
		return arr
	} else {
		arr = append(arr[:i], arr[i+1:]...)
		return arr
	}
}

func Bin(i int, str string, arr []string) []string {
	decimal, err := strconv.ParseInt(str, 2, 64)
	if err == nil {
		arr[i-1] = strconv.Itoa(int(decimal))
		arr = append(arr[:i], arr[i+1:]...)
		return arr
	} else {
		arr = append(arr[:i], arr[i+1:]...)
		return arr
	}
}

func CapInt(i int, arr []string) []string {
	str := strings.Trim(string(arr[i+1]), arr[i+1][1:])

	number, _ := strconv.Atoi(string(str))

	for j := 1; j <= number; j++ {
		if number > i-1 {
			number = i
		}
		arr[i-j] = strings.Title(strings.ToLower(arr[i-j]))
	}

	arr = append(arr[:i], arr[i+2:]...)
	return arr
}

func UpInt(i int, arr []string) []string {
	str := strings.Trim(string(arr[i+1]), arr[i+1][1:])

	number, _ := strconv.Atoi(string(str))

	for j := 1; j <= number; j++ {
		if number > i-1 {
			number = i
		}
		arr[i-j] = strings.ToUpper(arr[i-j])
	}

	arr = append(arr[:i], arr[i+2:]...)
	return arr
}

func LowInt(i int, arr []string) []string {
	str := strings.Trim(string(arr[i+1]), arr[i+1][1:])

	number, _ := strconv.Atoi(string(str))

	for j := 1; j <= number; j++ {
		if number > i-1 {
			number = i
		}
		arr[i-j] = strings.ToLower(arr[i-j])
	}

	arr = append(arr[:i], arr[i+2:]...)
	return arr
}

func An(i int, str string, arr []string) []string {
	res := str + "n"
	arr[i] = res
	arr = append(arr[:i+1], arr[i+1:]...)
	return arr
}

func AnReverse(i int, str string, arr []string) []string {
	var res string
	if str == "An" {
		res = strings.ReplaceAll(str, "An", "A")
	} else if str == "an" {
		res = strings.ReplaceAll(str, "an", "a")
	} else if str == "AN" {
		res = strings.ReplaceAll(str, "AN", "A")
	} else if str == "aN" {
		res = strings.ReplaceAll(str, "aN", "a")
	}
	arr[i] = res
	arr = append(arr[:i+1], arr[i+1:]...)
	return arr
}

func SignFormatter(s []string) []string {
	signs := []string{".", ",", "!", "?", ":", ";"}
	for i, word := range s {
		for _, key := range signs {
			if string(word[0]) == key && string(word[len(word)-1]) != key {
				s[i-1] += key
				s[i] = word[1:]
			}
		}
	}
	for i, word := range s {
		for _, key := range signs {
			if (string(word[0]) == key) && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}
	for i, word := range s {
		for _, key := range signs {
			if string(word[0]) == key && string(word[len(word)-1]) == key && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}
	count := 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count += 1
			s[i+1] = word + s[i+1]
			s = append(s[:i], s[i+1:]...)
		}
	}
	for i, word := range s {
		if word == "'" {
			s[i-1] = s[i-1] + word
			s = append(s[:i], s[i+1:]...)
		}
	}

	return s
}
