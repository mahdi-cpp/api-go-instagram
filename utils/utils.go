package utils

import (
	"bufio"
	"fmt"
	"image"
	"io"
	"io/fs"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func GetImageDimension(imagePath string) (int, int) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	image, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}
	return image.Width, image.Height
}

var numberOnly = regexp.MustCompile("^[0-9-_!@#$%^&*()]*$")

func isOnlyNumber(str string) bool {
	return numberOnly.MatchString(str)
}

func checkMovieName(fileName string) (*string, *string) {

	var funcName = "checkName:"

	var name = strings.ToLower(fileName)
	var year = 1900

	if strings.HasSuffix(name, ".mp4") || strings.HasSuffix(name, ".mkv") {

	} else {
		fmt.Println(funcName, "filetype must be .mp4 or .mkv suffix -->", "{"+name+"}")
		return nil, nil
	}

	name = strings.ReplaceAll(name, ".", " ")
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.ReplaceAll(name, "(", "")
	name = strings.ReplaceAll(name, ")", "")
	name = strings.ReplaceAll(name, "!", "")
	name = strings.ReplaceAll(name, "@", "")
	name = strings.ReplaceAll(name, "#", "")
	name = strings.ReplaceAll(name, "%", "")
	name = strings.ReplaceAll(name, "&", "")

	//fmt.Println(name)

	if isOnlyNumber(fileName) {
		fmt.Println(funcName, "filetype must be contain alphabetic characters -->", "{"+name+"}")
		return nil, nil
	}

	for i := 1900; i < 2050; i++ {
		if strings.Contains(name, strconv.Itoa(i)) {
			year = i
			index := strings.Index(name, strconv.Itoa(year))
			if index > 2 {
				name = substr(name, 0, index)
			} else {
				fmt.Println(funcName, "filename name and year not valid -->", "{"+name+"}")
				return nil, nil
			}
			break
		}
	}

	//fmt.Println(funcName, "FileName is Valid -->", "{"+name+"}")

	var year2 = strconv.Itoa(year)
	return &name, &year2
}

func checkSeriesName(fileName string) (*string, *string, *string) {

	var funcName = "checkName:"

	var name = strings.ToLower(fileName)
	var season = "01"
	var episode = "01"

	if strings.HasSuffix(name, ".mp4") || strings.HasSuffix(name, ".mkv") {
		name = strings.ReplaceAll(name, ".", " ")
		name = strings.ReplaceAll(name, "_", " ")
		name = strings.ReplaceAll(name, "(", "")
		name = strings.ReplaceAll(name, ")", "")
		name = strings.ReplaceAll(name, "!", "")
		name = strings.ReplaceAll(name, "@", "")
		name = strings.ReplaceAll(name, "#", "")
		name = strings.ReplaceAll(name, "%", "")
		name = strings.ReplaceAll(name, "&", "")
	} else {
		//fmt.Println(funcName, "filetype must be .mp4 or .mkv suffix -->", "{"+name+"}")
		return nil, nil, nil
	}

	//fmt.Println(name)

	if isOnlyNumber(fileName) {
		//fmt.Println(funcName, "filetype must be contain alphabetic characters -->", "{"+name+"}")
		return nil, nil, nil
	}

	for i := 1; i < 30; i++ {
		var ss = "s0" + strconv.Itoa(i)
		if strings.Contains(name, ss) {
			index := strings.Index(name, ss)
			episode = substr(name, index+3, 3)
			if index > 5 {
				name = substr(name, 0, index)
				//fmt.Println(name, ss, episode)
			} else {
				fmt.Println(funcName, "filename name and season and episode not valid -->", "{"+name+"}")
				return nil, nil, nil
			}
			return &name, &season, &episode
		}
	}

	return nil, nil, nil
}

func download(url string, PosterPath string) {

	//if isExistFile(PosterPath) {
	//	return
	//}

	PosterPath = strings.ReplaceAll(PosterPath, "/", "")

	// don't worry about errors
	response, e := http.Get(url)
	if e != nil {
		log.Fatal(e)
	}
	defer response.Body.Close()

	//open a file for writing
	file, err := os.Create("/var/instagram/posters/series/" + PosterPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Use io.Copy to just dump the response body to the file. This supports huge files
	_, err = io.Copy(file, response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success!")
}

func isExistFile(fileName string) bool {
	files, err := ioutil.ReadDir("/var/instagram/posters/00")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		//fmt.Println(file.Name(), file.IsDir())
		if strings.Contains(fileName, file.Name()) {
			fmt.Println("File is exist:", file.Name())
			return true
		}
	}
	return false
}

func readOfFile() {
	file, err := os.Open("/home/mahdi/photos/ali.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var url = "https://image.tmdb.org/t/p/original" + scanner.Text()
		fmt.Println(url)
		//fmt.Println("....")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// NOTE: this isn't multi-Unicode-codepoint aware, like specifying skintone or
//
//	gender of an emoji: https://unicode.org/emoji/charts/full-emoji-modifiers.html
func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}

func directory(path string, directory fs.DirEntry) {

	var root = path + "/" + directory.Name()
	fmt.Println("is Directory", root)
	files, err := os.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
