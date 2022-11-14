package controllers

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"crypto/tls"
	"net/http"
	"log"
	"fmt"
	"api/ex/v2/models"
	"api/ex/v2/config"
	"api/ex/v2/database"
)

type filter struct {
	name   string
	artist string
	album  string
}

var Config, _ = config.LoadConfig("./")

func GetSongOfAllClients(name string, artist string, album string) (element []any) {

	filterApple := "term=" + name
	filterChart := "song=" + name + "&artist=" + artist

	if len(name) > 0 {
		filterApple = "term=" + name
	}

	if len(artist) > 0 {
		filterApple = filterApple + ""
	}

	if len(album) > 0 {
		filterApple = filterApple + ""
	}

	fmt.Println("filter Apple : ", filterApple)
	fmt.Println("filter Chartlyrics : ", filterChart)
	dataChart := GetSongClientChart(filterChart)
	dataApple := GetSongApple(filterApple)

	element = append(dataChart, dataApple...)
	fmt.Println(len(element))

	return

}

func GetSongClientChart(url string) (songs []any) {

	var resp = getSongSoap(Config.ClientChartlyricsApi + url)

	for _, song := range resp.SearchLyricResult {
		maping := make(map[string]interface{})
		maping["IdSong"] = song.TrackId
		maping["Name"] = song.Song
		maping["artist"] = song.Artist
		maping["Duration"] = song.SongRank
		maping["Album"] = song.Song
		maping["Artwork"] = song.SongUrl
		maping["Price"] = song.SongRank
		maping["Origin"] = "chartlyrics"

		songs = append(songs, maping)
	}
	return
}

func GetSongApple(url string) (songs []any) {

	var resp = getSongRest(Config.ClientAppleApi + url)

	for _, song := range resp.Results {
		maping := make(map[string]interface{})
		maping["IdSong"] = song["trackId"]
		maping["Name"] = song["trackName"]
		maping["artist"] = song["artistName"]
		maping["Duration"] = song["trackTimeMillis"]
		maping["Album"] = song["collectionName"]
		maping["Artwork"] = song["previewUrl"]
		maping["Price"] = song["trackPrice"]
		maping["Origin"] = "Apple"

		songs = append(songs, maping)
	}
	return
}

func getSongRest(url string) (respSong models.ResponseClientApple) {

	resp, err := http.Get(url)

	if err != nil {
		fmt.Print(err.Error())
	}

	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err.Error())
	}

	json.Unmarshal(respData, &respSong)

	return
}

func getSongSoap(url string) (respSong models.ResponseClientChartlyrics) {

	c := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	html, err := c.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(html.Body)
	if err != nil {
		log.Fatal(err)
	}
	html.Body.Close()

	xml.Unmarshal(data, &respSong)

	return
}

func GetSong(name string, artist string, album string) (songs *models.Song, err error) {

	songFilter := make(map[string]interface{})

	if len(name) > 0 {
		songFilter["name"] = name
	}

	if len(artist) > 0 {
		songFilter["artist"] = artist
	}

	if len(album) > 0 {
		songFilter["album"] = album
	}
	jsonStr, err := json.Marshal(songFilter)

	var mapData map[string]interface{}
	json.Unmarshal(jsonStr, &mapData)
	fmt.Println("Somethin", mapData)
	database.DB.Where("name = ?", mapData["name"]).Or("artist = ?", mapData["artist"]).Or("album = ?", mapData["album"]).Find(&songs)
	return
}