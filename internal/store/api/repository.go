package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/JoseUgal/cmd-beers-api/cli"
)

type repository struct {
	client			*http.Client
}

func NewRepository() cli.ApiRepository {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	return &repository{
		client,
	}
}

const (
	reqAllBers = "https://api.punkapi.com/v2/beers/"
	reqOneBeer = "https://api.punkapi.com/v2/beers/"
)

func ( r *repository ) GetAllBeers() ([]cli.Beer, error) {
	
	var beers []cli.Beer 

	req, _ := r.client.Get(reqAllBers)

	body, _ := ioutil.ReadAll(req.Body)
	
	_ = json.Unmarshal(body, &beers)

	fmt.Println(beers)

	return beers, nil
}


func ( r *repository ) GetBeer( id int ) (cli.Beer, error) {

	var beers []cli.Beer

	req, _ := r.client.Get( reqOneBeer + strconv.Itoa(id))

	body, _ := ioutil.ReadAll(req.Body)
	
	_ = json.Unmarshal(body, &beers)

	fmt.Println(beers[0])

	return beers[0], nil
}