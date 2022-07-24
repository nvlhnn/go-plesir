package seeders

import (
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gosimple/slug"
	"github.com/lib/pq"
	"github.com/nvlhnn/go-plesir/model/domain"

	"gorm.io/gorm"
)

func PlaceSeed(db *gorm.DB){

	if db.Migrator().HasTable(&domain.Place{}){
		// if err := db.First(&domain.Day{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {	

		// seed := os.Getenv("SEEDING")

		// switch seed {
		// case :
			
		// }

		var places []domain.Place
		iterator := [20]uint{}
		for index, _ := range iterator {
			url := "https://www.tiket.com/to-do/api/filtered-products?startingPriceInCentsFrom=0&pageNumber="+strconv.Itoa(index+1)+"&sortAttributes=popularityScore&sortDirection=DESC&pageSize=50&excludes=operationalHours%2Csections%2Cfeatures%2Cpackages&productCategoryCodes=ATTRACTION&lang=id"
			res := getResponse(url)
			places = append(places, res...)
		}

		err := db.Create(&places).Error
		if err != nil {
			log.Println("[seeding error] : ",err )
		}
				
	}

	return

}


func getResponse(url string) []domain.Place {
	log.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		log.Println("migrate failed on get")
		log.Println(err)
		return nil
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)	
	if err != nil {
		log.Println("migrate failed on read")
		log.Println(err)

		return nil
	}

	type Img struct {
		Img string `json:"medium"`
	}
	type placeDays struct {
		Day string `json:"day"`
		Open string `json:"openTime"`
		Close string `json:"closeTime"`
	}
	type place struct {
		Name string `json:"title"`
		Region string `json:"region"`
		Price       float32 `json:"startPrice"`
		Images []Img  `json:"media"`
		PlaceDays *[]placeDays `json:"openDayTimes"`
	} 

	type data struct{
		Place []place `json:"products"`
	}

	type response struct {
		Data data `json:"data"` 
	}

	time.Sleep(20 * time.Second)
	var res response

	// log.Println(body)
	err = json.Unmarshal(body, &res)

	var places []domain.Place			

	if err != nil {		
		log.Println("migrate failed on unmarshal")
		log.Println(err)
		log.Printf("error decoding sakura response: %v", err)
		if e, ok := err.(*json.SyntaxError); ok {
			log.Printf("syntax error at byte offset %d", e.Offset)
		}
		log.Printf("sakura response: %q", res)
		return nil
	}else{
		for _, v := range res.Data.Place {		
			if v.PlaceDays != nil {
				// set images
				var images []string
				for _, img := range v.Images {
					images = append(images, img.Img)
				}

				// set days
				var placeDays []domain.PlaceDays
				for _, day := range *v.PlaceDays {
					var dayId uint
					switch day.Day {
					case "MONDAY":
						dayId = 1				
					case "TUESDAY":
						dayId = 2				
					case "WEDNESDAY":
						dayId = 3				
					case "THURSDAY":
						dayId = 4
					case "FRIDAY":
						dayId = 5
					case "SATURDAY":
						dayId = 6
					case "SUNDAY":
						dayId = 7
					}

					work := domain.PlaceDays{
						DayID: dayId,
						OpenTime: day.Open,
						CloseTime: day.Close,
					}
					placeDays = append(placeDays, work)

				}

				rand := rand.Intn(99999-11111) + 11111
				slug := slug.Make(v.Name) + "-" + strconv.Itoa(rand)

				thePlace := domain.Place{
					Name : v.Name,
					Region: v.Region,
					Description: "lorem ipsum",
					Slug: slug,
					Price: v.Price,
					Images: pq.StringArray(images),
					UserID: 1,
					PlaceDays: placeDays,

				}

				places = append(places, thePlace)			
			}				
					
		}
	}

	// err = db.Create(&places).Error
	// if err != nil {
	// 	log.Println("[seeding error] : ",err )
	// }
	return places
}