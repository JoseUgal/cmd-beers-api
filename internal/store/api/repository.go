package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/JoseUgal/cmd-beers-api/cli"
	"github.com/JoseUgal/cmd-beers-api/internal/errors"
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


func ( r *repository ) GetBeer( id int ) (beer cli.Beer,err error) {

	var beers []cli.Beer

	endpoint := fmt.Sprintf("%v%v", reqOneBeer, id)

	req, err := r.client.Get( endpoint )

	if err != nil {
		return cli.Beer{}, errors.WrapDataUnreacheable(err, "error getting response to %s", endpoint )
	}

	body, _ := ioutil.ReadAll(req.Body)
	
	_ = json.Unmarshal(body, &beers)

	beer = beers[0]

	return beer, nil
}