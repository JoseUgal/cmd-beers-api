package cli

type Beer struct {
	Id						int 			`json:"id"`
	Name					int 			`json:"name"`
	Tagline					string			`json:"tagline"`
    FirstbBrewed			string			`json:"first_brewed"`
    Description				string			`json:"description"`
    ImageUrl				string			`json:"image_url"`
    Abv						int				`json:"abv"`
    Ibu						int				`json:"ibu"`
    Target_fg				int				`json:"target_fg"`
    Target_og				int				`json:"target_og"`
    Ebc						int				`json:"ebc"`
    Srm						int				`json:"srm"`
    Ph						int				`json:"ph"`
    AttenuationLevel		int				`json:"attenuation_level"`
    Volume					BoilVolume		`json:"volume"`
    BoilVolume				BoilVolume		`json:"boil_volume"`
    Method					Method			`json:"method"`
    Ingredients				Ingredients		`json:"ingredients"`
    FoodPairing				[]string		`json:"food_pairing"`
    BrewersTips			string				`json:"brewers_tips"`
    ContributedBy			string			`json:"contributed_by"`
}

type BoilVolume struct {
	Value					int 			`json:"value"`
	Unit					string 			`json:"unit"`
}

type Ingredients struct {
	Malt					[]Malt 			`json:"malt"`
	Hops					[]Hop 			`json:"hops"`
	Yeast					string 			`json:"yeast"`
}

type Hop struct {
	Name					int 			`json:"name"`
	Amount					BoilVolume 		`json:"amount"`
	Add						string 			`json:"add"`
	Attribute				string 			`json:"attribute"`
}

type Malt struct {
	Name					string 			`json:"name"`
	Amount					BoilVolume 		`json:"amount"`
}

type Method struct {
	MashTemp				[]MashTemp 		`json:"mash_temp"`
	Fermentation			Fermentation 	`json:"fermentation"`
}

type Fermentation struct {
	Temp					BoilVolume 		`json:"temp"`
}

type MashTemp struct {
	Temp					BoilVolume 		`json:"temp"`
	Duration				int 			`json:"duration"`
}

// Define API Repository Interface
type ApiRepository interface {
	GetAllBeers() ([]Beer, error)
	GetBeer( id int ) (Beer, error)
}