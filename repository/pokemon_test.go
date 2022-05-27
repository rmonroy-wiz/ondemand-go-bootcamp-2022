package repository

import (
	"os"
	"reflect"
	"testing"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service"
	serviceMocks "github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service/mocks"
	"github.com/stretchr/testify/mock"
)

func TestNewPokemonRepository(t *testing.T) {
	type args struct {
		csvServ  service.CSV
		fileServ service.File
	}
	tests := []struct {
		name string
		args args
		want *pokemonRepository
	}{
		{
			name: "Create PokemonRepository instance",
			args: args{
				csvServ:  nil,
				fileServ: nil,
			},
			want: new(pokemonRepository),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPokemonRepository(tt.args.csvServ, tt.args.fileServ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPokemonRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pokemonRepository_GetAll(t *testing.T) {
	mockFileServ1 := new(serviceMocks.FileMock)
	mockFileServ1.On("OpenFile", mock.Anything, mock.Anything).Return(new(os.File), nil)
	mockFileServ1.On("Close")

	pokemonsCSV := []model.PokemonCSV{{
		ID:             1,
		Name:           "bulbasaur",
		Height:         100,
		Weight:         200,
		BaseExperience: 2,
		PrimaryType:    "Grass",
		SecondaryType:  "Poison",
	}}
	mockCSVServ1 := new(serviceMocks.CSVMock)
	mockCSVServ1.On("UnmarshalFile", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(1).(*[]model.PokemonCSV)
		*arg = pokemonsCSV
	}).Return(nil)

	pokemons := []*model.PokemonDTO{
		{
			ID:             1,
			Name:           "bulbasaur",
			Height:         100,
			Weight:         200,
			BaseExperience: 2,
			PrimaryType:    "Grass",
			SecondaryType:  "Poison",
		},
	}

	type fields struct {
		csvService  service.CSV
		fileService service.File
	}
	tests := []struct {
		name      string
		fields    fields
		want      []*model.PokemonDTO
		wantError *model.ErrorHandler
	}{
		{
			name: "Happy Path",
			fields: fields{
				csvService:  mockCSVServ1,
				fileService: mockFileServ1,
			},
			want:      pokemons,
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pokemonRepository{
				csvService:  tt.fields.csvService,
				fileService: tt.fields.fileService,
			}
			got, got1 := p.GetAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pokemonRepository.GetAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.wantError) {
				t.Errorf("pokemonRepository.GetAll() got1 = %v, want %v", got1, tt.wantError)
			}
		})
	}
}
