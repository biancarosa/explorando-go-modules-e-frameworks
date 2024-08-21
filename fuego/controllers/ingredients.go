package controller

import (
	"github.com/go-fuego/fuego"
)

type IngredientsRessources struct {
	// TODO add ressources
	IngredientsService IngredientsService
}

type Ingredients struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type IngredientsCreate struct {
	Name string `json:"name"`
}

type IngredientsUpdate struct {
	Name string `json:"name"`
}

func (rs IngredientsRessources) Routes(s *fuego.Server) {
	ingredientsGroup := fuego.Group(s, "/ingredients")

	fuego.Get(ingredientsGroup, "/", rs.getAllIngredients)
	fuego.Post(ingredientsGroup, "/", rs.postIngredients)

	fuego.Get(ingredientsGroup, "/{id}", rs.getIngredients)
	fuego.Put(ingredientsGroup, "/{id}", rs.putIngredients)
	fuego.Delete(ingredientsGroup, "/{id}", rs.deleteIngredients)
}

func (rs IngredientsRessources) getAllIngredients(c fuego.ContextNoBody) ([]Ingredients, error) {
	return rs.IngredientsService.GetAllIngredients()
}

func (rs IngredientsRessources) postIngredients(c *fuego.ContextWithBody[IngredientsCreate]) (Ingredients, error) {
	body, err := c.Body()
	if err != nil {
		return Ingredients{}, err
	}

	new, err := rs.IngredientsService.CreateIngredients(body)
	if err != nil {
		return Ingredients{}, err
	}

	return new, nil
}

func (rs IngredientsRessources) getIngredients(c fuego.ContextNoBody) (Ingredients, error) {
	id := c.PathParam("id")

	return rs.IngredientsService.GetIngredients(id)
}

func (rs IngredientsRessources) putIngredients(c *fuego.ContextWithBody[IngredientsUpdate]) (Ingredients, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return Ingredients{}, err
	}

	new, err := rs.IngredientsService.UpdateIngredients(id, body)
	if err != nil {
		return Ingredients{}, err
	}

	return new, nil
}

func (rs IngredientsRessources) deleteIngredients(c *fuego.ContextNoBody) (any, error) {
	return rs.IngredientsService.DeleteIngredients(c.PathParam("id"))
}

type IngredientsService interface {
	GetIngredients(id string) (Ingredients, error)
	CreateIngredients(IngredientsCreate) (Ingredients, error)
	GetAllIngredients() ([]Ingredients, error)
	UpdateIngredients(id string, input IngredientsUpdate) (Ingredients, error)
	DeleteIngredients(id string) (any, error)
}
