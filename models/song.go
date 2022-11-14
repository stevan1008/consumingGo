package models

import "encoding/xml"

type RandomSong struct {
	Results []Song
}

type Song struct {
	ID       uint 				`json:"id" `
	IdSong   string             `json:"idSong"`
	Name     string             `json:"name"`
	Artist   string				`json:"artist"`	
	Duration string             `json:"duration"`
	Album    string             `json:"album"`
	Artwork  string             `json:"artwork"`
	Price    string             `json:"price"`
	Origin   string             `json:"origin"`
}

type SongResponse struct {
	IdSong   string 	`json:"idSong"`
	Name     string 	`json:"name"`
	Artist 	 string		`json:"artist"`
	Duration string 	`json:"duration"`
	Album    string 	`json:"album"`
	Artwork  string 	`json:"artwork"`
	Price    string 	`json:"price"`
	Origin   string 	`json:"origin"`
}

type ResponseClientApple struct {
	ResultCount int              `json:"resultCount"`
	Results     []map[string]any `json:"results"`
}

type ResponseClientChartlyrics struct {
	XMLName           xml.Name `xml:"ArrayOfSearchLyricResult"`
	Text              string   `xml:",chardata"`
	Xsd               string   `xml:"xsd,attr"`
	Xsi               string   `xml:"xsi,attr"`
	Xmlns             string   `xml:"xmlns,attr"`
	SearchLyricResult []struct {
		Text          string `xml:",chardata"`
		Nil           string `xml:"nil,attr"`
		TrackId       string `xml:"TrackId"`
		LyricChecksum string `xml:"LyricChecksum"`
		LyricId       string `xml:"LyricId"`
		SongUrl       string `xml:"SongUrl"`
		ArtistUrl     string `xml:"ArtistUrl"`
		Artist        string `xml:"Artist"`
		Song          string `xml:"Song"`
		SongRank      string `xml:"SongRank"`
		TrackChecksum string `xml:"TrackChecksum"`
	} `xml:"SearchLyricResult"`
}

