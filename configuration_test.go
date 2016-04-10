package main

import (
	"net"
	"testing"
)

func TestLoadConfigurationInvalidFile(t *testing.T) {
	config := TinkerConfiguration{}
	err := config.LoadConfig("fixtures/config.json.notexist")
	if err == nil {
		t.Fatalf("Expected err =! nil, got %s", err)
	}
}

func TestLoadConfigurationInvalidJSON(t *testing.T) {
	config := TinkerConfiguration{}
	err := config.LoadConfig("fixtures/config.json.invalid")
	if err == nil {
		t.Fatalf("Expected err =! nil, got %s", err)
	}
}

func TestLoadConfigurationOk(t *testing.T) {
	config := TinkerConfiguration{}
	err := config.LoadConfig("fixtures/config.json")
	if err != nil {
		t.Fatalf("Expected err == nil, got %s", err)
	}
}

func TestConfiguration(t *testing.T) {
	configUser1 := UserConfiguration{
		Username:      "test",
		ContactEmail:  "test@tinker.org",
		TincAddress:   net.ParseIP("176.18.16.1"),
		PublicAddress: net.ParseIP("54.10.10.43"),
	}
	configUser2 := UserConfiguration{
		Username:      "test2",
		ContactEmail:  "test2@tinker.org",
		TincAddress:   net.ParseIP("176.18.16.2"),
		PublicAddress: net.ParseIP("54.51.110.21"),
	}
	_, subnet, _ := net.ParseCIDR("176.18.16.0/22")
	fakeConfig := TinkerConfiguration{
		Subnet: *subnet,
		Users:  []UserConfiguration{configUser1, configUser2},
	}
	config := TinkerConfiguration{}
	_ = config.LoadConfig("fixtures/config.json")
	if config.Subnet.String() != fakeConfig.Subnet.String() {
		t.Fatalf("Expected config.Subnet) == %s, got %s", fakeConfig.Subnet.String(), config.Subnet.String())
	}
	if len(config.Users) != len(fakeConfig.Users) {
		t.Fatalf("Expected len(config.Users) == %d, got %d", len(fakeConfig.Users), len(config.Users))
	}
	for count, fakeUser := range fakeConfig.Users {
		if config.Users[count].Username != fakeUser.Username {
			t.Fatalf("Expected user.Username == %s, got %s", fakeUser.Username, config.Users[count].Username)
		}
		if config.Users[count].ContactEmail != fakeUser.ContactEmail {
			t.Fatalf("Expected user.ContactEmail == %s, got %s", fakeUser.ContactEmail, config.Users[count].ContactEmail)
		}
		if config.Users[count].TincAddress.String() != fakeUser.TincAddress.String() {
			t.Fatalf("Expected user.TincAddress == %s, got %s", fakeUser.TincAddress.String(), config.Users[count].TincAddress.String())
		}
		if config.Users[count].PublicAddress.String() != fakeUser.PublicAddress.String() {
			t.Fatalf("Expected user.PublicAddress == %s, got %s", fakeUser.PublicAddress.String(), config.Users[count].PublicAddress.String())
		}
	}
}
