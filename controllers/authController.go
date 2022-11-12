package controllers

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"api/ex/v2/models"
	"api/ex/v2/database"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
	"net/http"
	"encoding/json"
	"fmt"
)

var client *http.Client
const SecretKey = "AlgoSecreto"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := models.User{
		Name: data["name"],
		Email: data["email"],
		Password: password,
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string
	
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "PAssword Incorrect",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), //1 day
	})

	token, err := claims.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not login",
		}) 
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "Perfectly",
	})
}

func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	claims := token.Claims.(*jwt.StandardClaims)

	var user models.User

	database.DB.Where("id = ?", claims.Issuer).First(&user)

	return c.JSON(user)
} 

func Logout (c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name: 		"jwt",
		Value: 		"",
		Expires:	time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Message": "All success",
	})
}


func GetJson(url string, target interface{}) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func GetRandomMusic(c *fiber.Ctx) error {

	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// database.DB.Where("id = ?", claims.Issuer).First(&user)

	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	url := "https://itunes.apple.com/search?term=jack+johnson"

	var music models.RandomMusic
	client = &http.Client{Timeout: 10 * time.Second}
	er := GetJson(url, &music)

	if er != nil {
		// fmt.Printf("error getting music: %s\n", err.Error())
		return c.JSON(fiber.Map{
			"Message": "error getting music",
		})
	} else {
		// fmt.Println(music.Results)
		musicStored := []models.Music{}
		for _, mus := range music.Results {
			musicStored = append(musicStored, mus)
		}
		fmt.Println(musicStored)

		for _, elem := range musicStored {
			var music = models.Music {
				TrackId: elem.TrackId,
				TrackName: elem.TrackName,
				ArtistName: elem.ArtistName,
				TrackTimeMillis: elem.TrackTimeMillis,
				CollectionName: elem.CollectionName,
				ArtworkUrl30: elem.ArtworkUrl30,
				TrackPrice: elem.TrackPrice,
				Country: elem.Country,
			}
			database.DB.Create(&music)
		}
		// artist := music.Results[0].ArtistName
		// artWork := music.Results[0].ArtworkUrl30
		// collectionName := music.Results[0].CollectionName
		// country := music.Results[0].Country
		// trackname := music.Results[0].TrackName
		// trackId := music.Results[0].TrackId
		// price := music.Results[0].TrackPrice
		// duration := music.Results[0].TrackTimeMillis

		
		return c.JSON(fiber.Map{
			"Message": "Success",
		})

		// fmt.Printf("Music here: %s %s %s %s %s %s %s %s\n", music.Results[0].ArtistName,
		// 	music.Results[0].ArtworkUrl30,
		// 	music.Results[0].CollectionName,
		// 	music.Results[0].Country,
		// 	music.Results[0].TrackName,
		// 	music.Results[0].TrackId,
		// 	music.Results[0].TrackPrice,
		// 	music.Results[0].TrackTimeMillis)
	}
}