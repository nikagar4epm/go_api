package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		Username:  "alex",
	},
	"jason": {
		AuthToken: "456DEF",
		Username:  "jason",
	},
	"marie": {
		AuthToken: "789GHI",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    100,
		Username: "alex",
	},
	"jason": {
		Coins:    200,
		Username: "jason",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

var mockClientProfiles = map[string]UserDetails{
	"alex": {
		Email:    "alex@example.com",
		Id:       "123",
		Username: "alex",
		Name:     "Alexander",
		Token:    "123ABC",
	},
	"jason": {
		Email:    "jason@example.com",
		Id:       "456",
		Username: "jason",
		Name:     "Jason",
		Token:    "456DEF",
	},
	"marie": {
		Email:    "marie@example.com",
		Id:       "789",
		Username: "marie",
		Name:     "Marie",
		Token:    "789GHI",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserProfile(username string) *UserDetails {
	time.Sleep(time.Second * 1)

	var clientData = UserDetails{}
	clientData, ok := mockClientProfiles[username]

	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetUserProfile(username string, profile UserDetails) *UserDetails {
	time.Sleep(time.Second * 1)

	var clientData = UserDetails{}
	clientData, ok := mockClientProfiles[username]

	if !ok {
		return nil
	}

	mockClientProfiles[username] = profile

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
