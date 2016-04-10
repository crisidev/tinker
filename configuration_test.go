package main

import (
	"testing"
)

func TestLoadConfigurationErr(t *testing.T) {
	config := TinkerConfiguration{}
	err := config.LoadFromFile("non-existent-file.json")
	if err == nil {
		t.Fatalf("Expected err == nil, got %s", err)
	}
}

func TestLoadConfigurationOk(t *testing.T) {
	config := TinkerConfiguration{}
	err := config.LoadFromFile("fixtures/config.json")
	if err != nil {
		t.Fatalf("Expected err != nil, got %s", err)
	}
}

func TestConfiguration(t *testing.T) {
	configUser1 := UserConfiguration{
		Address:      "10.0.0.1",
		ContactEmail: "test@tinker.org",
		Roadwarrior:  false,
		Subnets:      []string{"10.0.0.1/32", "192.168.66.0/24"},
		Username:     "testuser",
	}
	configUser2 := UserConfiguration{
		Address:      "10.0.0.2",
		ContactEmail: "test2@tinker.org",
		Roadwarrior:  true,
		Subnets:      []string{"10.0.0.2/32"},
		Username:     "testuser2",
	}
	fakeConfig := TinkerConfiguration{
		Users: []UserConfiguration{configUser1, configUser2},
	}
	config := TinkerConfiguration{}
	_ = config.LoadFromFile("fixtures/config.json")
	if len(config.Users) != len(fakeConfig.Users) {
		t.Fatalf("Expected len(config.Users) == %d, got %d", len(fakeConfig.Users), len(config.Users))
	}
	for count, fakeUser := range fakeConfig.Users {
		if config.Users[count].Address != fakeUser.Address {
			t.Fatalf("Expected user.Address == %s, got %s", fakeUser.Address, config.Users[count].Address)
		}
		if config.Users[count].ContactEmail != fakeUser.ContactEmail {
			t.Fatalf("Expected user.ContactEmail == %s, got %s", fakeUser.ContactEmail, config.Users[count].ContactEmail)
		}
		if config.Users[count].Address != fakeUser.Address {
			t.Fatalf("Expected user.Roadwarrior == %s, got %s", fakeUser.Roadwarrior, config.Users[count].Roadwarrior)
		}
		if config.Users[count].Subnets[0] != fakeUser.Subnets[0] {
			t.Fatalf("Expected user.Subnets == %s, got %s", fakeUser.Subnets, config.Users[count].Subnets)
		}
		if config.Users[count].Username != fakeUser.Username {
			t.Fatalf("Expected user.Username == %s, got %s", fakeUser.Username, config.Users[count].Username)
		}
	}
}
