package xkcd_comics

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"time"
)

type Comic struct {
	Title      string `json:"safe_title"`
	Number     int    `json:"num"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Link       string `json:"img"`
	Transcript string `json:"transcript"`
}

type lastComic struct {
	Number int `json:"num"`
}

func (comic Comic) printInfo() {
	_, err := fmt.Printf("Title: %s\nDate: %s/%s/%s\nNumber: %d\nLink: %s\n", comic.Title, comic.Year, comic.Month, comic.Day, comic.Number, comic.Link)
	if err != nil {
		fmt.Println(err)
	}
}

func getLastComic() int {
	res, err := http.Get("https://xkcd.com/info.0.json")
	if err != nil {
		log.Fatal("Error getting the latest comic number. Try again later.")
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	comic := lastComic{}
	err = json.Unmarshal(body, &comic)
	if err != nil {
		log.Fatal(err)
	}

	return comic.Number
}

func GetComicByNumber(number int) Comic {
	lastComicNumber := getLastComic()
	if lastComicNumber < number {
		log.Fatal("The number specified was bigger than the number of the last comic. Please insert one lower to this: ", lastComicNumber)
	}

	url := fmt.Sprintf("https://xkcd.com/%d/info.0.json", number)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	comic := Comic{}
	err = json.Unmarshal(body, &comic)
	if err != nil {
		log.Fatal(err)
	}

	download(comic)
	return comic
}

func GetRandomComic() Comic {
	lastComicNumber := getLastComic()
	rand.Seed(time.Now().Unix())

	randomNumber := 1 + rand.Intn(lastComicNumber-1) + 1
	randomComic := GetComicByNumber(randomNumber)

	return randomComic
}

func download(comic Comic) {
	comic.printInfo()
	res, err := http.Get(comic.Link)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	fileName := fmt.Sprintf("%s.png", comic.Title)
	image, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	//defer image.Close()
	_, err = io.Copy(image, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	image.Close()
	displayImage := exec.Command("feh", fileName)
	err = displayImage.Run()
	os.Exit(0)
}
