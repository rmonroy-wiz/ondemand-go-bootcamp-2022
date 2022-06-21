package repository

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"
	"sync"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model/mapper"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service"
)

//go:generate mockery --name PokemonRepository --filename pokemon.go --outpkg mocks --structname PokemonRepositoryMock --disable-version-string
type PokemonRepository interface {
	GetAll() ([]model.PokemonDTO, *model.ErrorHandler)
	GetByID(id int) (model.PokemonDTO, *model.ErrorHandler)
	StoreToCSV(pokemonAPI model.PokemonAPI) (model.PokemonDTO, *model.ErrorHandler)
	GetCSVDataInMemory() (map[int]model.PokemonCSV, *model.ErrorHandler)
	WorkerPoolSearchPokemon(typeSearch string, items int, itemsPerWorker int) ([]model.PokemonDTO, *model.ErrorHandler)
}

// PokemonRepository structure for repository, contains the csv file's name
type pokemonRepository struct {
	csvService     service.CSV
	fileService    service.File
	readerCSVMutex sync.Mutex
}

// NewPokemonRepository method for create a Repository instance
func NewPokemonRepository(csvServ service.CSV, fileServ service.File) *pokemonRepository {
	return &pokemonRepository{
		csvService:  csvServ,
		fileService: fileServ,
	}
}

// GetAll get all pokemons from csv file
func (p *pokemonRepository) GetAll() ([]model.PokemonDTO, *model.ErrorHandler) {
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
func (p *pokemonRepository) GetByID(id int) (model.PokemonDTO, *model.ErrorHandler) {
	pokemons, err := p.GetAll()
	if err != nil {
		return model.PokemonDTO{}, err
	}

	for _, pokemon := range pokemons {
		if pokemon.ID == id {
			return pokemon, nil
		}
	}

	return model.PokemonDTO{}, model.NewNotFoundPokemonError(id)
}

// StoreToCSV store pokemon to csv
func (p *pokemonRepository) StoreToCSV(pokemonAPI model.PokemonAPI) (model.PokemonDTO, *model.ErrorHandler) {
	pokemonMap, err := p.GetCSVDataInMemory()
	if err != nil {
		return model.PokemonDTO{}, err
	}
	pokemon := mapper.PokemonAPItoPokemonCSV(pokemonAPI)
	pokemonMap[pokemon.ID] = pokemon
	pokemons := make([]model.PokemonCSV, 0)
	for _, pokemonObj := range pokemonMap {
		pokemons = append(pokemons, pokemonObj)
	}

	pokemonFile, errFile := p.fileService.OpenFile(os.O_RDWR|os.O_CREATE, os.ModePerm)
	if errFile != nil {
		return model.PokemonDTO{}, model.NewOpenFileError(errFile.Error())
	}
	if err := p.csvService.MarshalFile(&pokemons, pokemonFile); err != nil {
		return model.PokemonDTO{}, model.NewAccesingCSVFileError(err.Error())
	}
	defer p.fileService.Close()
	return mapper.PokemonAPIToPokemonDTO(pokemonAPI), nil
}

// getCSVDataInMemory store pokemons from csv to memory
func (p *pokemonRepository) GetCSVDataInMemory() (map[int]model.PokemonCSV, *model.ErrorHandler) {
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

func (p *pokemonRepository) WorkerPoolSearchPokemon(typeSearch string, items int, itemsPerWorker int) ([]model.PokemonDTO, *model.ErrorHandler) {
	var wg sync.WaitGroup
	file, err := p.fileService.OpenFile(os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, model.NewOpenFileError(err.Error())
	}
	csvReader := csv.NewReader(file)

	_, err = csvReader.Read()
	if err != nil {
		return nil, model.NewAccesingCSVFileError(err.Error())
	}

	defer p.fileService.Close()

	resultChannel := make(chan model.PokemonDTO, items)
	itemsWorkers := items
	w := 1
	for itemsWorkers > 0 {
		wg.Add(1)
		if (itemsWorkers) > itemsPerWorker {
			go p.workerSearchPokemon(w, &wg, csvReader, typeSearch, itemsPerWorker, resultChannel)
		} else {
			go p.workerSearchPokemon(w, &wg, csvReader, typeSearch, itemsWorkers, resultChannel)
		}
		itemsWorkers = itemsWorkers - itemsPerWorker
		w++
	}

	wg.Wait()
	pokemons := make([]model.PokemonDTO, 0)
	for x := 0; x < items; x++ {
		pokemon := <-resultChannel
		pokemons = append(pokemons, pokemon)
	}
	return pokemons, nil
}

func (p *pokemonRepository) workerSearchPokemon(id int, wg *sync.WaitGroup, csvReader *csv.Reader, typeSearch string, itemsPerWorker int, result chan<- model.PokemonDTO) {
	defer wg.Done()

	var count int = 0
	for {
		if count == itemsPerWorker {
			return
		}
		pokemon, err := p.getCSVRow(csvReader)
		if err != nil || err == io.EOF {
			return
		}
		if p.isTypeOf(pokemon, typeSearch) {
			result <- pokemon
			count = count + 1
		}
	}
}

func (p *pokemonRepository) getCSVRow(csvReader *csv.Reader) (model.PokemonDTO, error) {
	p.readerCSVMutex.Lock()
	defer p.readerCSVMutex.Unlock()

	data, err := csvReader.Read()
	if err != nil {
		return model.PokemonDTO{}, err
	}
	idValue, err := strconv.Atoi(data[0])
	if err != nil {
		return model.PokemonDTO{}, err
	}
	heightValue, err := strconv.Atoi(data[2])
	if err != nil {
		return model.PokemonDTO{}, err
	}
	weightValue, err := strconv.Atoi(data[3])
	if err != nil {
		return model.PokemonDTO{}, err
	}
	baseExperienceValue, err := strconv.Atoi(data[4])
	if err != nil {
		return model.PokemonDTO{}, err
	}
	return model.PokemonDTO{
		ID:             idValue,
		Name:           data[1],
		Height:         heightValue,
		Weight:         weightValue,
		BaseExperience: baseExperienceValue,
		PrimaryType:    data[5],
		SecondaryType:  data[6],
	}, nil
}

func (p *pokemonRepository) isTypeOf(pokemon model.PokemonDTO, typeSearch string) bool {
	if typeSearch == "odd" {
		return (pokemon.ID%2 == 0)
	} else {
		return (pokemon.ID%2 != 0)
	}
}
