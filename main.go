package main

import (
	"github.com/gofiber/fiber/v2"
	"api/ex/v2/database"
	"api/ex/v2/routes"
)

// type RandomMusic struct {
// 	Results []Music
// }

// type Music struct {
// 	TrackId         uint
// 	TrackName       string
// 	ArtistName      string
// 	TrackTimeMillis uint
// 	CollectionName  string
// 	ArtworkUrl30    string
// 	TrackPrice      float32
// 	Country         string 
// }

// func GetJson(url string, target interface{}) error {
// 	resp, err := client.Get(url)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()

// 	return json.NewDecoder(resp.Body).Decode(target)
// }

// func GetRandomMusic() {
// 	url := "https://itunes.apple.com/search?term=jack+johnson"

// 	var music models.RandomMusic

// 	err := GetJson(url, &music)

// 	if err != nil {
// 		fmt.Printf("error getting music: %s\n", err.Error())
// 		return
// 	} else {
// 		fmt.Printf("Music here: %s %s %s %s %s %s %s %s\n", music.Results[0].ArtistName,
// 			music.Results[0].ArtworkUrl30,
// 			music.Results[0].CollectionName,
// 			music.Results[0].Country,
// 			music.Results[0].TrackName,
// 			music.Results[0].TrackId,
// 			music.Results[0].TrackPrice,
// 			music.Results[0].TrackTimeMillis)
// 	}
// }

func main() {

	database.Connect()


	// client = &http.Client{Timeout: 10 * time.Second}
	// controllers.GetRandomMusic()

	app := fiber.New()
	routes.SetUp(app)	
	app.Listen(":3100")
}