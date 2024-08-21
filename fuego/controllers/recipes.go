package controllers

import (
	"github.com/go-fuego/fuego"
)

type RecipesRessources struct {
	// TODO add ressources
	RecipesService RecipesServiceImpl
}

type Recipes struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RecipesCreate struct {
	Name string `json:"name"`
}

type RecipesUpdate struct {
	Name string `json:"name"`
}

func (rs RecipesRessources) Routes(s *fuego.Server) {
	recipesGroup := fuego.Group(s, "/recipes")

	fuego.Get(recipesGroup, "/", rs.getAllRecipes)
	fuego.Post(recipesGroup, "/", rs.postRecipes)

	fuego.Get(recipesGroup, "/{id}", rs.getRecipes)
	fuego.Put(recipesGroup, "/{id}", rs.putRecipes)
	fuego.Delete(recipesGroup, "/{id}", rs.deleteRecipes)
}

func (rs RecipesRessources) getAllRecipes(c fuego.ContextNoBody) ([]Recipes, error) {
	return rs.RecipesService.GetAllRecipes()
}

func (rs RecipesRessources) postRecipes(c *fuego.ContextWithBody[RecipesCreate]) (Recipes, error) {
	body, err := c.Body()
	if err != nil {
		return Recipes{}, err
	}

	new, err := rs.RecipesService.CreateRecipes(body)
	if err != nil {
		return Recipes{}, err
	}

	return new, nil
}

func (rs RecipesRessources) getRecipes(c fuego.ContextNoBody) (Recipes, error) {
	id := c.PathParam("id")

	return rs.RecipesService.GetRecipes(id)
}

func (rs RecipesRessources) putRecipes(c *fuego.ContextWithBody[RecipesUpdate]) (Recipes, error) {
	id := c.PathParam("id")

	body, err := c.Body()
	if err != nil {
		return Recipes{}, err
	}

	new, err := rs.RecipesService.UpdateRecipes(id, body)
	if err != nil {
		return Recipes{}, err
	}

	return new, nil
}

func (rs RecipesRessources) deleteRecipes(c *fuego.ContextNoBody) (any, error) {
	return rs.RecipesService.DeleteRecipes(c.PathParam("id"))
}

type RecipesService interface {
	GetRecipes(id string) (Recipes, error)
	CreateRecipes(RecipesCreate) (Recipes, error)
	GetAllRecipes() ([]Recipes, error)
	UpdateRecipes(id string, input RecipesUpdate) (Recipes, error)
	DeleteRecipes(id string) (any, error)
}

type RecipesServiceImpl struct {
	RecipesService
}

func (s RecipesServiceImpl) GetAllRecipes() ([]Recipes, error) {
	return []Recipes{
		{
			ID:   "1",
			Name: "Freijão fradinho",
		},
	}, nil
}

func (s RecipesServiceImpl) GetRecipes(id string) (Recipes, error) {
	return Recipes{ID: id, Name: "nome"}, nil

}

func (s RecipesServiceImpl) CreateRecipes(RecipesCreate) (Recipes, error) {
	recipe := new(Recipes)
	return *recipe, nil
}

func (s RecipesServiceImpl) UpdateRecipes(id string, input RecipesUpdate) (Recipes, error) {
	return Recipes{
		ID:   "1",
		Name: "Freijão fradinho",
	}, nil
}

func (s RecipesServiceImpl) DeleteRecipes(id string) (any, error) {
	return []Recipes{}, nil
}
