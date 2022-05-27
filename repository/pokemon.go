package repository

import (
	"os"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model/mapper"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service"
)

//go:generate mockery --name PokemonRepository --filename pokemon.go --outpkg mocks --structname PokemonRepositoryMock --disable-version-string
type PokemonRepository interface {
	GetAll() ([]*model.PokemonDTO, *model.ErrorHandler)
	GetByID(id int) (*model.PokemonDTO, *model.ErrorHandler)
	StoreToCSV(pokemonAPI model.PokemonAPI) (*model.PokemonDTO, *model.ErrorHandler)
	GetCSVDataInMemory() (map[int]model.PokemonCSV, *model.ErrorHandler)
}

// PokemonRepository structure for repository, contains the csv file's name
type pokemonRepository struct {
	csvService  service.CSV
	fileService service.File
}

// NewPokemonRepository method for create a Repository instance
func NewPokemonRepository(csvServ service.CSV, fileServ service.File) *pokemonRepository {
	return &pokemonRepository{
		csvService:  csvServ,
		fileService: fileServ,
	}
}

// GetAll get all pokemons from csv file
func (p *pokemonRepository) GetAll() ([]*model.PokemonDTO, *model.ErrorHandler) {
	pokemonFile, err := p.fileService.OpenFile(os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, model.NewOpenFileError(err.Error())
	}
	pokemons := []model.PokemonCSV{}

	if err := p.csvService.UnmarshalFile(pokemonFile, &pokemons); err != nil {
		return nil, model.NewUnmarshalFileError(err.Error())
	}

	defer p.fileService.Close()
	return mapper.PokemonsCSVToPokemonsDTO(pokemons), nil
}

// GetByID get pokemon from csv by id
func (p pokemonRepository) GetByID(id int) (*model.PokemonDTO, *model.ErrorHandler) {
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}

	for _, pokemon := range pokemons {
		if pokemon.ID == id {
			return pokemon, nil
		}
	}

	return nil, model.NewNotFoundPokemonError(id)
}

// StoreToCSV store pokemon to csv
func (p *pokemonRepository) StoreToCSV(pokemonAPI model.PokemonAPI) (*model.PokemonDTO, *model.ErrorHandler) {
	pokemonMap, err := p.GetCSVDataInMemory()
	if err != nil {
		return nil, err
	}
	pokemon := mapper.PokemonAPItoPokemonCSV(pokemonAPI)
	pokemonMap[pokemon.ID] = pokemon
	pokemons := make([]model.PokemonCSV, 0)
	for _, pokemonObj := range pokemonMap {
		pokemons = append(pokemons, pokemonObj)
	}

	pokemonFile, errFile := p.fileService.OpenFile(os.O_RDWR|os.O_CREATE, os.ModePerm)
	if errFile != nil {
		return nil, model.NewOpenFileError(errFile.Error())
	}
	if err := p.csvService.MarshalFile(&pokemons, pokemonFile); err != nil {
		return nil, model.NewAccesingCSVFileError(err.Error())
	}
	defer p.fileService.Close()
	return mapper.PokemonAPIToPokemonDTO(pokemonAPI), nil
}

// getCSVDataInMemory store pokemons from csv to memory
func (p pokemonRepository) GetCSVDataInMemory() (map[int]model.PokemonCSV, *model.ErrorHandler) {
	pokemonMap := make(map[int]model.PokemonCSV)
	pokemons, err := p.GetAll()
	if err != nil {
		return nil, err
	}
	for _, pokemon := range pokemons {
		pokemonMap[pokemon.ID] = mapper.PokemonDTOToPokemonCSV(pokemon)
	}
	return pokemonMap, nil
}
