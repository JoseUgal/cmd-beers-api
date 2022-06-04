package cmd

import (
	. "fmt"

	"github.com/JoseUgal/cmd-beers-api/internal/errors"
	store "github.com/JoseUgal/cmd-beers-api/internal/store/api"
	i "github.com/tockins/interact"
)

// Enum that includes all the application's option
type _QuestResponse = int
const (
	_Unknown _QuestResponse = iota
	_ALL
	_ONE
    _INGREDIENTS
)

var apiRepository = store.NewRepository()

// Method that allows to print the menu of the
// application by console
func DrawUI(){

	Println("🚧 This application uses internet resources for its operation")
	Println("👨‍💻 API: https://punkapi.com/")
	Println()
    
    // Execute APP Quiz with infinite loop
    for {
      cmdQuiz("main")
    }
}


func cmdQuiz( option string ){
    switch option {
        case "main":
                i.Run(&i.Interact{
                Questions: []*i.Question{
                    {
                        Quest: i.Quest{
                            Msg:     "🧐 What do you want to get?",
                            Choices: i.Choices{
                                Alternatives: []i.Choice{
                                    {
                                        Text: "🍻 All beers",
                                        Response: _ALL,
                                    },
                                    {
                                        Text: "🍺 A beer",
                                        Response: _ONE,
                                    },
                                    {
                                        Text: "🍋 Ingredients of a beer",
                                        Response: _INGREDIENTS,
                                    },
                                },
                            },
                        },
                        Action: func(c i.Context) interface{} {
                            val, _ := c.Ans().Int()
                            
                            switch val {
                            case int64(_ALL):
                                beers, err := apiRepository.GetAllBeers()

                                if errors.IsDataUnreacheable(err) {
                                    Println(err)
                                } else {
                                    Println(beers)
                                }
                            case int64(_ONE):
                                cmdQuiz("beer")
                            default:
                                Println("🚫 La opción seleccionada no está disponible en estos momentos")
                            }
        
                            return nil
                        },
                    },
                },
            }) 
        case "beer":
            i.Run(&i.Interact{
                Questions: []*i.Question{
                    {
                        Quest: i.Quest{
                            Msg: "🍺 what beer do you want to look for?",
                        },
                        Action: func(c i.Context) interface{} {
                            val, _ := c.Ans().Int()
                            
                            beer, err := apiRepository.GetBeer(int(val))

                            if errors.IsDataUnreacheable(err) {
                                Println(err)
                            }else{
                                Println(beer)
                            }
        
                            return nil
                        },
                    },
                },
            }) 
        default:
            Println("La vista que intentas cargar no existe.")
    }
   
}