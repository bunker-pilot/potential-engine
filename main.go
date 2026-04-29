package main

import(
	"github.com/bunker-pilot/potential-engine/internal/app"
)

func main(){
	app, err := app.Newapplication()
	if err !=nil{
		panic(err)
	}

	app.Logger.Println("We are running our app :)")
}