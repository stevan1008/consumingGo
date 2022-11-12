package models

type Music struct {
	TrackId         uint	
	TrackName       string	
	ArtistName      string	
	TrackTimeMillis uint	
	CollectionName  string	
	ArtworkUrl30    string	
	TrackPrice      float32	
	Country         string 	
}

type RandomMusic struct {
	Results []Music
}