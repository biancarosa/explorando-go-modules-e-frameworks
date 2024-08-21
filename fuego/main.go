package main

import (
	"fmt"

	"github.com/biancarosa/explorando-go-modules-e-frameworks/controllers"
	"github.com/go-fuego/fuego"
)

func main() {
	fmt.Println("Hello world!")

	s := fuego.NewServer()

	// bia preferiria se fosse s.Get()
	fuego.Get(s, "/", helloWorld)

	// bia preferiria se fosse s.AddRoute(controllers.RecipesRessources)
	controllers.RecipesRessources{}.Routes(s)

	s.Run()
}

func helloWorld(ctx fuego.ContextNoBody) (string, error) {
	return "Hello world", nil
}
