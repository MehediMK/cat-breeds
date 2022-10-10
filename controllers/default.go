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
	api_key := goDotEnvVariable("API_KEY")

	// here get all api breeds
	breeds_url := "https://api.thecatapi.com/v1/breeds"

	req := httplib.Get(breeds_url)
	res, err := req.String()
	if err != nil {
		fmt.Println(err)
	}

	// here define structure for breeds
	breeds := breed.Breeds{}
	err_1 := json.Unmarshal([]byte(res), &breeds)
	if err_1 != nil {
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
	breed_image_url := "https://api.thecatapi.com/v1/images/search?limit=25&breed_ids=" + id + "&api_key=" + api_key
	req_img := httplib.Get(breed_image_url)
	res_img, err_img := req_img.String()
	if err_img != nil {
		fmt.Println(err_img)
	}

	// here define structure for breedsimage
	breeds_image := breed.Breeds_images{}
	errimg1 := json.Unmarshal([]byte(res_img), &breeds_image)
	if errimg1 != nil {
		fmt.Println("Here some error get")
	}

	c.Data["name"] = breeds_image[0].Breeds[0].Name
	c.Data["cat_id"] = breeds_image[0].Breeds[0].ID
	c.Data["description"] = breeds_image[0].Breeds[0].Description
	c.Data["temperament"] = breeds_image[0].Breeds[0].Temperament
	c.Data["origin"] = breeds_image[0].Breeds[0].Origin
	c.Data["weight"] = breeds_image[0].Breeds[0].Weight.Metric
	c.Data["life_span"] = breeds_image[0].Breeds[0].LifeSpan
	c.Data["wikipedia_url"] = breeds_image[0].Breeds[0].WikipediaURL

	fmt.Println(breeds_image[0].Breeds[0].Name)

	c.Data["breeds"] = breeds
	c.Data["breeds_images"] = breeds_image
	c.TplName = "index.html"
}
