package controllers

import (
	breed "cat-breeds/models"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/beego/beego/httplib"
	"github.com/beego/beego/v2/server/web"

	// for env
	"github.com/joho/godotenv"
)

type MainController struct {
	web.Controller
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func (c *MainController) Get() {

	// load .env file
	apikey := goDotEnvVariable("API_KEY")

	// here get all api breeds
	breedsurl := "https://api.thecatapi.com/v1/breeds"

	req := httplib.Get(breedsurl)
	res, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	// here define structure for breeds
	breeds := breed.Breeds{}
	err1 := json.Unmarshal([]byte(res), &breeds)
	if err1 != nil {
		fmt.Println("Here some error get")
	}
	fmt.Println(breeds[0].Id)

	// get request params
	var id string
	c.Ctx.Input.Bind(&id, "breed")

	if id == "" {
		id = breeds[0].Id
	}
	c.Data["reid"] = id
	fmt.Println("URL: " + id)

	// here api call for images
	breedimageurl := "https://api.thecatapi.com/v1/images/search?limit=25&breed_ids=" + id + "&api_key=" + apikey
	reqimg := httplib.Get(breedimageurl)
	resimg, errimg := reqimg.String()
	if errimg != nil {
		fmt.Println(errimg)
	}

	// here define structure for breedsimage
	breedsimage := breed.Breedsimage{}
	errimg1 := json.Unmarshal([]byte(resimg), &breedsimage)
	if errimg1 != nil {
		fmt.Println("Here some error get")
	}

	c.Data["name"] = breedsimage[0].Breeds[0].Name
	c.Data["catid"] = breedsimage[0].Breeds[0].ID
	c.Data["description"] = breedsimage[0].Breeds[0].Description
	c.Data["temperament"] = breedsimage[0].Breeds[0].Temperament
	c.Data["origin"] = breedsimage[0].Breeds[0].Origin
	c.Data["weight"] = breedsimage[0].Breeds[0].Weight.Metric
	c.Data["life_span"] = breedsimage[0].Breeds[0].LifeSpan
	c.Data["wikipedia_url"] = breedsimage[0].Breeds[0].WikipediaURL

	fmt.Println(breedsimage[0].Breeds[0].Name)

	c.Data["breeds"] = breeds
	c.Data["breedsimages"] = breedsimage
	c.TplName = "index.html"
}
