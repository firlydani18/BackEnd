package service

import (
	"KosKita/features/kos"
	"KosKita/features/user"
	"KosKita/mocks"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	input := kos.CoreInput{
		UserID:      1,
		Name:        "KOS",
		Description: "kos ini murah sekali",
		Category:    "putra",
		Price:       100000,
		Rooms:       10,
		Address:     "Jl. Kost Murah No.1",
		Longitude:   "123.456",
		Latitude:    "98.768",
		KosFacilities: []kos.KosFacilityCore{
			{
				Facility: "AC",
			},
			{
				Facility: "TV",
			},
		},
		KosRules: []kos.KosRuleCore{
			{
				Rule: "Dilarang Merokok",
			},
			{
				Rule: "Dilarang Membawa Hewan Peliharaan",
			},
		},
	}

	userIdLogin := int(input.UserID)

	t.Run("user is not owner", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "renter"}, nil).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		kosId, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.Equal(t, uint(0), kosId)
		assert.EqualError(t, err, "anda bukan owner")
	})

	t.Run("error from userService", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(nil, errors.New("user not found")).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		kosId, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.Equal(t, uint(0), kosId)
		assert.EqualError(t, err, "user not found")

		srvUser.AssertExpectations(t)
	})

	t.Run("name is empty", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		input.Name = ""
		_, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "name anda kosong")
		input.Name = "Kos Murah"
	})

	t.Run("category is empty", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		input.Category = ""
		_, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "category anda kosong")
		input.Category = "putra"
	})

	t.Run("price is zero or less", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		input.Price = 0
		_, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "harga anda kosong")
		input.Price = 1000000
	})

	t.Run("rooms is zero or less", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		input.Rooms = 0
		_, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "rooms anda kosong")
		input.Rooms = 10
	})

	t.Run("address is empty", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		input.Address = ""
		_, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "alamat anda kosong")
		input.Address = "Jl. Kost Murah No.1"
	})

	t.Run("error from repository", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil)
		repo.On("Insert", userIdLogin, input).Return(uint(0), errors.New("database error")).Once()

		inputInvalid := input
	inputInvalid.Name = "" 

		kosId, err := srv.Create(userIdLogin, input)

		assert.Error(t, err)
		assert.NotNil(t, kosId)
		assert.Equal(t, uint(0), kosId)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil)
		repo.On("Insert", userIdLogin, input).Return(input.ID, nil).Once()

		kosId, err := srv.Create(userIdLogin, input)

		assert.NoError(t, err)
		assert.Equal(t, input.ID, kosId)

		repo.AssertExpectations(t)
	})
}

func TestPut(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	input := kos.CoreInput{
		UserID:      1,
		Name:        "KOS",
		Description: "kos ini murah sekali",
		Category:    "putra",
		Price:       100000,
		Rooms:       10,
		Address:     "Jl. Kost Murah No.1",
		Longitude:   "123.456",
		Latitude:    "98.768",
		KosFacilities: []kos.KosFacilityCore{
			{
				Facility: "AC",
			},
			{
				Facility: "TV",
			},
		},
		KosRules: []kos.KosRuleCore{
			{
				Rule: "Dilarang Merokok",
			},
			{
				Rule: "Dilarang Membawa Hewan Peliharaan",
			},
		},
	}

	userIdLogin := int(input.UserID)

	t.Run("user is not owner", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "renter"}, nil).Once()

		err := srv.Put(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "anda bukan owner")

		srvUser.AssertExpectations(t)
	})

	t.Run("error from userService", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(nil, errors.New("user not found")).Once()

		err := srv.Put(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")

		srvUser.AssertExpectations(t)
	})

	t.Run("error from repository", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil)
		repo.On("Update", userIdLogin, input).Return(errors.New("database error")).Once()

		err := srv.Put(userIdLogin, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil)
		repo.On("Update", userIdLogin, input).Return(nil).Once()

		err := srv.Put(userIdLogin, input)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}

func TestCreateRating(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	inputRating := kos.RatingCore{
		Score:           5,
		BoardingHouseID: 1,
		UserID:          1,
	}

	userIdLogin := int(inputRating.UserID)
	kosId := int(inputRating.BoardingHouseID)

	t.Run("score out of range", func(t *testing.T) {
		inputRating.Score = 6

		err := srv.CreateRating(userIdLogin, kosId, inputRating)

		assert.Error(t, err)
		assert.EqualError(t, err, "skor hanya bisa 1 sampai 5")

		repo.AssertExpectations(t)
	})

	t.Run("error from CekRating", func(t *testing.T) {
		inputRating.Score = 5
		repo.On("CekRating", userIdLogin, kosId).Return(nil, errors.New("database error")).Once()

		err := srv.CreateRating(userIdLogin, kosId, inputRating)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("existing rating", func(t *testing.T) {
		inputRating.Score = 5
		repo.On("CekRating", userIdLogin, kosId).Return(&inputRating, nil).Once()
		err := srv.CreateRating(userIdLogin, kosId, inputRating)

		assert.Error(t, err)
		assert.EqualError(t, err, "anda sudah pernah memberikan rating untuk kos ini")
	})

	t.Run("error from InsertRating", func(t *testing.T) {
		repo.On("CekRating", userIdLogin, kosId).Return(nil, nil).Once()
		repo.On("InsertRating", userIdLogin, kosId, inputRating).Return(errors.New("database error")).Once()

		err := srv.CreateRating(userIdLogin, kosId, inputRating)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("success", func(t *testing.T) {
		repo.On("CekRating", userIdLogin, kosId).Return(nil, nil).Once()
		repo.On("InsertRating", userIdLogin, kosId, inputRating).Return(nil).Once()

		err := srv.CreateRating(userIdLogin, kosId, inputRating)

		assert.NoError(t, err)
	})
}

func TestGetByRating(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	kosData := []kos.Core{
		{
			ID:          1,
			Name:        "Kos",
			Description: "Kos Murah",
			Category:    "campur",
			Price:       10000,
			Rooms:       11,
			Address:     "Jakarta selatan",
			Longitude:   "98.768",
			Latitude:    "98.768",
			KosFacilities: []kos.KosFacilityCore{
				{
					Facility: "AC",
				},
			},
			KosRules: []kos.KosRuleCore{
				{
					Rule: "GABOLEH MANDI",
				},
			},
			PhotoMain:       "https://cloudinary.com/foto",
			PhotoFront:      "https://cloudinary.com/foto",
			PhotoBack:       "https://cloudinary.com/foto",
			PhotoRoomFront:  "https://cloudinary.com/foto",
			PhotoRoomInside: "https://cloudinary.com/foto",
			CreatedAt:       time.Date(2034, 11, 11, 0, 0, 0, 0, time.UTC),
			UpdatedAt:       time.Date(2043, 11, 11, 0, 0, 0, 0, time.UTC),
		},
	}

	t.Run("error from repo", func(t *testing.T) {
		repo.On("SelectByRating").Return(nil, errors.New("database error")).Once()

		result, err := srv.GetByRating()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		repo.On("SelectByRating").Return(kosData, nil).Once()

		result, err := srv.GetByRating()

		assert.NoError(t, err)
		assert.Equal(t, kosData, result)

		repo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	input := kos.CoreInput{
		UserID: 1,
		ID:     1,
	}

	userIdLogin := int(input.UserID)
	kosId := int(input.ID)

	t.Run("error from SelectById", func(t *testing.T) {
		repo.On("SelectById", kosId).Return(nil, errors.New("database error")).Once()

		err := srv.Delete(userIdLogin, kosId)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("kos not found", func(t *testing.T) {
		repo.On("SelectById", kosId).Return(nil, errors.New("record not found")).Once()

		err := srv.Delete(userIdLogin, kosId)

		assert.Error(t, err)
		assert.EqualError(t, err, "kos id tidak ada")

		repo.AssertExpectations(t)
	})

	t.Run("kos not owned by user", func(t *testing.T) {
		kosData := &kos.Core{
			ID:     uint(kosId),
			UserID: uint(userIdLogin + 1),
		}
		repo.On("SelectById", kosId).Return(kosData, nil).Once()

		err := srv.Delete(userIdLogin, kosId)

		assert.Error(t, err)
		assert.EqualError(t, err, "kos ini bukan milik Anda")

		repo.AssertExpectations(t)
	})

	t.Run("error from repo", func(t *testing.T) {
		kosData := &kos.Core{
			ID:     uint(kosId),
			UserID: uint(userIdLogin),
			User:   user.Core{ID: uint(userIdLogin)},
		}
		repo.On("SelectById", kosId).Return(kosData, nil).Once()
		repo.On("Delete", userIdLogin, kosId).Return(errors.New("database error")).Once()

		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		err := srv.Delete(userIdLogin, kosId)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")
	})

	t.Run("success", func(t *testing.T) {
		kosData := &kos.Core{
			ID:     uint(kosId),
			UserID: uint(userIdLogin),
			User:   user.Core{ID: uint(userIdLogin)},
		}
		repo.On("SelectById", kosId).Return(kosData, nil).Once()
		repo.On("Delete", userIdLogin, kosId).Return(nil).Once()

		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()

		err := srv.Delete(userIdLogin, kosId)

		assert.NoError(t, err)
	})
}

func TestGetById(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	kosId := 1

	t.Run("error from repo", func(t *testing.T) {
		repo.On("SelectById", kosId).Return(nil, errors.New("database error")).Once()

		_, err := srv.GetById(kosId)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		kosData := &kos.Core{
			ID:     uint(kosId),
			UserID: uint(1),
		}
		repo.On("SelectById", kosId).Return(kosData, nil).Once()

		result, err := srv.GetById(kosId)

		assert.NoError(t, err)
		assert.Equal(t, kosData, result)

		repo.AssertExpectations(t)
	})
}

func TestGetByUserId(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	userIdLogin := 1

	t.Run("error from repo", func(t *testing.T) {
		repo.On("SelectByUserId", userIdLogin).Return(nil, errors.New("database error")).Once()

		_, err := srv.GetByUserId(userIdLogin)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		kosData := []kos.Core{
			{
				ID:     uint(1),
				UserID: uint(userIdLogin),
			},
		}
		repo.On("SelectByUserId", userIdLogin).Return(kosData, nil).Once()

		result, err := srv.GetByUserId(userIdLogin)

		assert.NoError(t, err)
		assert.Equal(t, kosData, result)

		repo.AssertExpectations(t)
	})
}

func TestSearchKos(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	query := "query"
	category := "category"
	minPrice := 100
	maxPrice := 200

	t.Run("error from repo", func(t *testing.T) {
		repo.On("SearchKos", query, category, minPrice, maxPrice).Return(nil, errors.New("database error")).Once()

		_, err := srv.SearchKos(query, category, minPrice, maxPrice)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("no kos found", func(t *testing.T) {
		repo.On("SearchKos", query, category, minPrice, maxPrice).Return([]kos.Core{}, nil).Once()

		_, err := srv.SearchKos(query, category, minPrice, maxPrice)

		assert.Error(t, err)
		assert.EqualError(t, err, "tidak ada kos yang ditemukan dengan filter yang dipilih")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		kosData := []kos.Core{
			{
				ID:     uint(1),
				UserID: uint(1),
			},
		}
		repo.On("SearchKos", query, category, minPrice, maxPrice).Return(kosData, nil).Once()

		result, err := srv.SearchKos(query, category, minPrice, maxPrice)

		assert.NoError(t, err)
		assert.Equal(t, kosData, result)

		repo.AssertExpectations(t)
	})
}

func TestCreateImage(t *testing.T) {
	repo := new(mocks.KosDataInterface)
	srvUser := new(mocks.UserServiceInterface)
	srv := New(repo, srvUser)

	userIdLogin := 1
	kosId := 1
	input := kos.CoreFoto{
		UserID:          uint(userIdLogin),
		PhotoMain:       "main.jpg",
		PhotoFront:      "front.jpg",
		PhotoBack:       "back.jpg",
		PhotoRoomFront:  "room_front.jpg",
		PhotoRoomInside: "room_inside.jpg",
	}

	t.Run("error from userService", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(nil, errors.New("user not found")).Once()

		err := srv.CreateImage(userIdLogin, kosId, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "user not found")

		srvUser.AssertExpectations(t)
	})

	t.Run("user is not owner", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "renter"}, nil).Once()

		err := srv.CreateImage(userIdLogin, kosId, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "anda bukan owner")

		srvUser.AssertExpectations(t)
	})

	t.Run("error from repo", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()
		repo.On("InsertImage", userIdLogin, kosId, input).Return(errors.New("database error")).Once()

		err := srv.CreateImage(userIdLogin, kosId, input)

		assert.Error(t, err)
		assert.EqualError(t, err, "database error")

		repo.AssertExpectations(t)
	})

	t.Run("success", func(t *testing.T) {
		srvUser.On("GetById", userIdLogin).Return(&user.Core{Role: "owner"}, nil).Once()
		repo.On("InsertImage", userIdLogin, kosId, input).Return(nil).Once()

		err := srv.CreateImage(userIdLogin, kosId, input)

		assert.NoError(t, err)

		repo.AssertExpectations(t)
	})
}
