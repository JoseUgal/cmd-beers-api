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

	req, err := r.client.Get(reqAllBers)

	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s" )
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error parsing response to bytes")
	}
	
	err = json.Unmarshal(body, &beers)

	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error marshaling bytes to 'Beers")
	}

	return beers, nil
}


func ( r *repository ) GetBeer( id int ) (beer cli.Beer,err error) {

	var beers []cli.Beer

	endpoint := fmt.Sprintf("%v%v", reqOneBeer, id)

	req, err := r.client.Get( endpoint )

	if err != nil {
		return cli.Beer{}, errors.WrapDataUnreacheable(err, "error getting response to %s", endpoint )
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return cli.Beer{}, errors.WrapDataUnreacheable(err, "error parsing response to bytes")
	}
	
	err = json.Unmarshal(body, &beers)

	if err != nil {
		return cli.Beer{}, errors.WrapDataUnreacheable(err, "error marshaling bytes to 'Beer")
	}

	beer = beers[0]

	return beer, nil
}