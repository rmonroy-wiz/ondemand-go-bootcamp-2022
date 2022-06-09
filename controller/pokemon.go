package controller

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/business"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
)

type pokemon struct {
	pokemonBusiness business.PokemonBusiness
	baseController
}

func NewPokemonController(pokemonBusiness business.PokemonBusiness) *pokemon {
	return &pokemon{
		pokemonBusiness: pokemonBusiness,
	}
}

// Get all pokemons
func (ctrl pokemon) GetAllPokemons(c *gin.Context) {
	pokemons, err := ctrl.pokemonBusiness.GetAll()

	if err != nil {
		ctrl.ResponseError(c, err)
	}

	ctrl.ResponseSucess(c, http.StatusOK, pokemons)
}

// GetPokemonByID get pokemon based on ID
func (ctrl pokemon) GetPokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("pokemonId"))
	if err != nil {
		ctrl.ResponseError(c, model.NewURLParameterDoesNotFound("pokemonId"))
		return
	}
	pokemon, errBusiness := ctrl.pokemonBusiness.GetByID(id)
	if errBusiness != nil {
		ctrl.ResponseError(c, errBusiness)
		return
	}

	ctrl.ResponseSucess(c, http.StatusOK, pokemon)
}

// StorePokemonByID get pokemon based on ID
func (ctrl pokemon) StorePokemonByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("pokemonId"))
	if err != nil {
		ctrl.ResponseError(c, model.NewURLParameterDoesNotFound("pokemonId"))
		return
	}
	pokemon, errBusiness := ctrl.pokemonBusiness.StoreByID(id)
	if errBusiness != nil {
		ctrl.ResponseError(c, errBusiness)
		return
	}

	ctrl.ResponseSucess(c, http.StatusCreated, pokemon)
}

// SearchPokemon
func (ctrl pokemon) SearchPokemon(c *gin.Context) {
	typeParam := c.Query("type")
	if typeParam == "" {
		ctrl.ResponseError(c, model.NewQueryParameterDoesNotFound("type"))
		return
	}
	typeValidValues := []string{"odd", "even"}
	if contains(typeValidValues, typeParam) {
		ctrl.ResponseError(c, model.NewQueryParameterDoesNotContainsValidValues("type", typeValidValues))
	}

	items := c.Query("items")
	if items == "" {
		ctrl.ResponseError(c, model.NewQueryParameterDoesNotFound("items"))
		return
	}
	if reflect.TypeOf(items).Kind().String() == "int" {
		ctrl.ResponseError(c, model.NewQueryParameterDoesNotHaveValidTypeValue("items", "int"))
		return
	}

	itemsPerWorkers := c.Query("items_per_workers")
	if itemsPerWorkers == "" {
		ctrl.ResponseError(c, model.NewQueryParameterDoesNotFound("items"))
		return
	}
	if reflect.TypeOf(items).Kind().String() == "int" {
		ctrl.ResponseError(c, model.NewQueryParameterDoesNotHaveValidTypeValue("items_per_workers", "int"))
		return
	}

	itemsInt, _ := strconv.Atoi(items)
	itempsPerWorkersInt, _ := strconv.Atoi(itemsPerWorkers)

	pokemons, errBusiness := ctrl.pokemonBusiness.SearchPokemonSingle(typeParam, itemsInt, itempsPerWorkersInt)

	if errBusiness != nil {
		ctrl.ResponseError(c, errBusiness)
		return
	}

	ctrl.ResponseSucess(c, http.StatusOK, pokemons)
}

func contains(array []string, search string) bool {
	for _, value := range array {
		if search == value {
			return true
		}
	}
	return false
}
