package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Incorrect number of arguments")
		return
	}
	file1 := args[0]
	file2 := args[1]

	f, err := os.Open(file1)
	if err != nil {
		fmt.Print(file1)
		fmt.Println(" file does not exist")
		return
	}
	defer f.Close()

	arr := bufio.NewScanner(f)

	data := ""
	var tmp []byte
	var ans []byte
	for arr.Scan() {
		data = arr.Text()
		var slice []string

		slice = strings.Fields(data)
		slice = articles(slice)

		slice = correction(slice)
		points := regexp.MustCompile(`(\s*)((\.\.\.)|(!\?)|(\.)|(,)|(!)|(\?)|(;)|(:))(\s*)`)
		data = points.ReplaceAllString(strings.Join(slice, " "), "$2 ")

		if len(data) != 0 && data[len(data)-1] == ' ' {
			data = data[:(len(data) - 1)]
		}

		data = apostroph(data)

		tmp = []byte(data)
		tmp = append(tmp, '\n')
		ans = append(ans, tmp...)
	}

	ioutil.WriteFile(file2, ans, 0644)
}

func articles(slice []string) []string {
	for i := 0; i < len(slice); i++ {
		if slice[i] == "an" {
			slice[i] = "a"
		}
		if slice[i] == "An" {
			slice[i] = "A"
		}
		vowels := []byte{'e', 'y', 'u', 'i', 'o', 'a', 'h'}
		for j := 0; j < len(vowels); j++ {
			if slice[i] == "A " && slice[i+1][0] == vowels[j] {
				slice[i] = "An"
			}
			if slice[i] == "a" && slice[i+1][0] == vowels[j] {
				slice[i] = "an"
			}
		}

	}
	return slice
}

func correction(slice []string) []string {
	for i := 0; i < len(slice); i++ {

		if slice[0] == "(up," || slice[0] == "(cap," || slice[0] == "(low," {
			slice = slice[2:]
			continue
		}
		if strings.Contains(slice[i], "(cap,") {
			number, _ := strconv.Atoi(strings.Trim(slice[i+1], ")"))
			for j := 1; j <= number; j++ {
				if slice[i-j+1] == slice[0] {
					break
				}
				if slice[i-j] == "." || slice[i-j] == "," || slice[i-j] == "!" || slice[i-j] == "?" || slice[i-j] == ";" || slice[i-j] == ":" || slice[i-j] == "..." || slice[i-j] == "!?" {
					number++
				} else {
					slice[i-j] = strings.ToLower(slice[i-j])
					slice[i-j] = strings.Title(slice[i-j])
				}
			}
			slice = append(slice[:i], slice[i+2:]...)
			i--
		}
		if slice[0] == "(cap)" || slice[0] == "(cap)," {
			slice[0] = strings.Trim(slice[0], "(cap)")
		} else if slice[i] == "(cap)" || slice[i] == "(cap)," {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i-1] = strings.Title(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}

		if slice[0] == "(bin)" || slice[0] == "(bin)," {
			slice[0] = strings.Trim(slice[0], "(bin)")
		} else if slice[i] == "(bin)" && i == len(slice)-1 {
			if isBinary(slice[i-1]) {
				tmp, _ := strconv.ParseInt(slice[i-1], 2, 64)
				slice[i-1] = strconv.Itoa(int(tmp))
				slice = slice[:i]
				i--
			} else {
				slice[i] = strings.Trim(slice[i], "(bin)")
			}
		} else if slice[i] == "(bin)" {
			if isBinary(slice[i-1]) {
				tmp, _ := strconv.ParseInt(slice[i-1], 2, 64)
				slice[i-1] = strconv.Itoa(int(tmp))
				slice = append(slice[:i], slice[i+1:]...)
				i--
			} else {
				slice[i] = strings.Trim(slice[i], "(bin)")
			}
		}

		if slice[0] == "(hex)" || slice[0] == "(hex)," {
			slice[0] = strings.Trim(slice[0], "(hex)")
		} else if slice[i] == "(hex)" && i == len(slice)-1 {
			if isHex(slice[i-1]) {
				tmp, _ := strconv.ParseInt(slice[i-1], 16, 64)
				slice[i-1] = strconv.Itoa(int(tmp))
				slice = slice[:i]
				i--
			} else {
				slice[i] = strings.Trim(slice[i], "(hex)")
			}
		} else if slice[i] == "(hex)" {
			if isHex(slice[i-1]) {
				tmp, _ := strconv.ParseInt(slice[i-1], 16, 64)
				slice[i-1] = strconv.Itoa(int(tmp))
				slice = append(slice[:i], slice[i+1:]...)
				i--
			} else {
				slice[i] = strings.Trim(slice[i], "(hex)")
			}
		}

		if slice[0] == "(low)" || slice[0] == "(low)," {
			slice[0] = strings.Trim(slice[0], "(low)")
		} else if slice[i] == "(low)" || slice[i] == "(low)," {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
		if strings.Contains(slice[i], "(low,") {
			number, _ := strconv.Atoi(strings.Trim(slice[i+1], ")"))
			for j := 1; j <= number; j++ {
				if slice[i-j+1] == slice[0] {
					break
				}
				if slice[i-j] == "." || slice[i-j] == "," || slice[i-j] == "!" || slice[i-j] == "?" || slice[i-j] == ";" || slice[i-j] == ":" || slice[i-j] == "..." || slice[i-j] == "!?" {
					number++
				} else {
					slice[i-j] = strings.ToLower(slice[i-j])
				}
			}
			slice = append(slice[:i], slice[i+2:]...)
			i--
		}
		if strings.Contains(slice[i], "(up,") {

			number, _ := strconv.Atoi(strings.Trim(slice[i+1], ")"))
			for j := 1; j <= number; j++ {
				if slice[i-j+1] == slice[0] {
					break
				}
				if slice[i-j] == "." || slice[i-j] == "," || slice[i-j] == "!" || slice[i-j] == "?" || slice[i-j] == ";" || slice[i-j] == ":" || slice[i-j] == "..." || slice[i-j] == "!?" {
					number++
				} else {
					slice[i-j] = strings.ToUpper(slice[i-j])
				}
			}
			slice = append(slice[:i], slice[i+2:]...)
			i--
		}

		if slice[i] == "(up)" || slice[i] == "(up)," {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	return slice
}

func apostroph(s string) string {
	r := []rune(s)

	i := 0
	count := 0
	for ; i < len(r)-1; i++ {
		if r[i] == '\'' && r[i+1] == ' ' || r[i] == 8216 && r[i+1] == ' ' {
			for j := i + 1; r[j] == ' '; j++ {
				count++
			}
			r = append(r[:i+1], r[i+count+1:]...)
			break
		}
	}
	count = 0
	i++
	for ; i < len(r); i++ {
		if r[i] == '\'' && r[i-1] == ' ' || r[i] == 8216 && r[i-1] == ' ' {
			for j := i - 1; r[j] == ' '; j-- {
				r = append(r[:i-1], r[i:]...)
				i = i - 1
			}
		}
	}

	return string(r)
}

func isBinary(s string) bool {
	r := []rune(s)
	for i := 0; i < len(r); i++ {
		if r[i] != '0' && r[i] != '1' {
			return false
		}
	}
	return true
}

func isHex(s string) bool {
	r := []rune(s)
	notHex := false
	for i := 0; i < len(r); i++ {
		if (r[i] >= '0' && r[i] <= '9') || (r[i] >= 'a' && r[i] <= 'f') || (r[i] >= 'A' && r[i] <= 'F') {
			continue
		} else {
			notHex = true
		}
	}
	return !notHex
}
