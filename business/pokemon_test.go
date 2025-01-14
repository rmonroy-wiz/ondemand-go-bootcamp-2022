package business

import (
	"reflect"
	"testing"

	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/model"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/repository"
	repoMocks "github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/repository/mocks"
	"github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service"
	serviceMocks "github.com/rmonroy-wiz/ondemand-go-bootcamp-2022/service/mocks"
)

func Test_pokemonBusiness_GetAll(t *testing.T) {
	mockRepo1 := new(repoMocks.PokemonRepositoryMock)
	mockRepo1.On("GetAll").Return([]model.PokemonDTO{}, nil)

	mockRepo2 := new(repoMocks.PokemonRepositoryMock)
	mockRepo2.On("GetAll").Return(nil, new(model.ErrorHandler))

	want1 := []model.PokemonDTO{}

	type fields struct {
		pokemonRepository repository.PokemonRepository
		serviceAPI        service.ExternalPokemonAPI
	}
	tests := []struct {
		name      string
		fields    fields
		want      []model.PokemonDTO
		wantError *model.ErrorHandler
	}{
		{
			name: "Get all pokemons",
			fields: fields{
				pokemonRepository: mockRepo1,
				serviceAPI:        nil,
			},
			want:      want1,
			wantError: nil,
		},
		{
			name: "Error getting pokemons",
			fields: fields{
				pokemonRepository: mockRepo2,
				serviceAPI:        nil,
			},
			want:      make([]model.PokemonDTO, 0),
			wantError: &model.ErrorHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pokemonBusiness{
				pokemonRepository: tt.fields.pokemonRepository,
				serviceAPI:        tt.fields.serviceAPI,
			}
			got, gotError := s.GetAll()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pokemonBusiness.GetAll() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotError, tt.wantError) {
				t.Errorf("pokemonBusiness.GetAll() gotError = %v, wantError %v", gotError, tt.wantError)
			}
		})
	}
}

func Test_pokemonBusiness_StoreByID(t *testing.T) {
	identifier := int(5)

	mockService1 := new(serviceMocks.ExternalPokemonAPIMock)
	mockService1.On("GetPokemonFromAPI", identifier).Return(model.PokemonAPI{}, new(model.ErrorHandler))

	pokemonAPI := model.PokemonAPI{
		ID:             identifier,
		BaseExperience: 50,
		Height:         20,
		Weight:         30,
		Name:           "Porygon",
		Types: []model.TypeSlot{
			{
				Type: model.Type{
					Name: "Grass",
				},
			},
		},
	}

	pokemonDTO := model.PokemonDTO{
		ID:             identifier,
		BaseExperience: 50,
		Height:         20,
		Weight:         30,
		Name:           "Porygon",
		PrimaryType:    "Grass",
		SecondaryType:  "",
	}

	mockService2 := new(serviceMocks.ExternalPokemonAPIMock)
	mockService2.On("GetPokemonFromAPI", identifier).Return(pokemonAPI, nil)

	mockRepo2 := new(repoMocks.PokemonRepositoryMock)
	mockRepo2.On("StoreToCSV", pokemonAPI).Return(model.PokemonDTO{}, new(model.ErrorHandler))

	mockService3 := new(serviceMocks.ExternalPokemonAPIMock)
	mockService3.On("GetPokemonFromAPI", identifier).Return(pokemonAPI, nil)

	mockRepo3 := new(repoMocks.PokemonRepositoryMock)
	mockRepo3.On("StoreToCSV", pokemonAPI).Return(pokemonDTO, nil)

	type fields struct {
		pokemonRepository repository.PokemonRepository
		serviceAPI        service.ExternalPokemonAPI
	}
	type args struct {
		id int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      model.PokemonDTO
		wantError *model.ErrorHandler
	}{
		{
			name: "Error in the GetPokemonAPI service",
			fields: fields{
				pokemonRepository: nil,
				serviceAPI:        mockService1,
			},
			args: args{
				id: identifier,
			},
			want:      model.PokemonDTO{},
			wantError: &model.ErrorHandler{},
		},
		{
			name: "Error in the StoreToCSV repository",
			fields: fields{
				pokemonRepository: mockRepo2,
				serviceAPI:        mockService2,
			},
			args: args{
				id: identifier,
			},
			want:      model.PokemonDTO{},
			wantError: &model.ErrorHandler{},
		},
		{
			name: "Happy path",
			fields: fields{
				pokemonRepository: mockRepo3,
				serviceAPI:        mockService3,
			},
			args: args{
				id: identifier,
			},
			want:      pokemonDTO,
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pokemonBusiness{
				pokemonRepository: tt.fields.pokemonRepository,
				serviceAPI:        tt.fields.serviceAPI,
			}
			got, gotError := s.StoreByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pokemonBusiness.StoreByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotError, tt.wantError) {
				t.Errorf("pokemonBusiness.StoreByID() gotError = %v, wantError %v", gotError, tt.wantError)
			}
		})
	}
}

func Test_pokemonBusiness_GetByID(t *testing.T) {
	identifier := int(5)

	mockRepo1 := new(repoMocks.PokemonRepositoryMock)
	mockRepo1.On("GetByID", identifier).Return(model.PokemonDTO{}, new(model.ErrorHandler))

	pokemonDTO := model.PokemonDTO{
		ID:             identifier,
		BaseExperience: 50,
		Height:         20,
		Weight:         30,
		Name:           "Porygon",
		PrimaryType:    "Grass",
		SecondaryType:  "",
	}

	mockRepo2 := new(repoMocks.PokemonRepositoryMock)
	mockRepo2.On("GetByID", identifier).Return(pokemonDTO, nil)

	type fields struct {
		pokemonRepository repository.PokemonRepository
		serviceAPI        service.ExternalPokemonAPI
	}
	type args struct {
		id int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		want      model.PokemonDTO
		wantError *model.ErrorHandler
	}{
		{
			name: "Error in the GetByID repository",
			fields: fields{
				pokemonRepository: mockRepo1,
				serviceAPI:        nil,
			},
			args: args{
				id: identifier,
			},
			want:      model.PokemonDTO{},
			wantError: new(model.ErrorHandler),
		},
		{
			name: "Happy path",
			fields: fields{
				pokemonRepository: mockRepo2,
				serviceAPI:        nil,
			},
			args: args{
				id: identifier,
			},
			want:      pokemonDTO,
			wantError: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := pokemonBusiness{
				pokemonRepository: tt.fields.pokemonRepository,
				serviceAPI:        tt.fields.serviceAPI,
			}
			got, gotError := s.GetByID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pokemonBusiness.GetByID() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(gotError, tt.wantError) {
				t.Errorf("pokemonBusiness.GetByID() gotError = %v, wantError %v", gotError, tt.wantError)
			}
		})
	}
}

func TestNewPokemonBusiness(t *testing.T) {
	mockRepo := new(repoMocks.PokemonRepositoryMock)
	mockService := new(serviceMocks.ExternalPokemonAPIMock)
	type args struct {
		repository repository.PokemonRepository
		service    service.ExternalPokemonAPI
	}
	tests := []struct {
		name string
		args args
		want *pokemonBusiness
	}{
		{
			name: "Create new pokemon business instance",
			args: args{
				repository: mockRepo,
				service:    mockService,
			},
			want: &pokemonBusiness{mockRepo, mockService},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPokemonBusiness(tt.args.repository, tt.args.service); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPokemonBusiness() = %v, want %v", got, tt.want)
			}
		})
	}
}
