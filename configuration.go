package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/user"
	"path"
)

type TinkerConfiguration struct {
	ConfigDir    string
	RepoDir      string
	GitRepo      string `json:"git_repo"`
	SubnetString string `json:"subnet"`
	Subnet       net.IPNet
	Users        []UserConfiguration `json:"users"`
}

type UserConfiguration struct {
	Enabled       bool
	Username      string `json:"username"`
	ContactEmail  string `json:"contact_email"`
	TincAddress   net.IP `json:"tinc_address"`
	PublicAddress net.IP `json:"public_address"`
	Key           string `json:"key"`
}

// Create configuration directory inside $HOME/.config/tinker
func (t *TinkerConfiguration) SetupConfigDir() (err error) {
	usr, err := user.Current()
	if err != nil {
		log.Error("unable to find current user")
		return err
	}
	t.ConfigDir = path.Join(usr.HomeDir, ".config", "tinker")
	if _, err := os.Stat(t.ConfigDir); os.IsNotExist(err) {
		log.Noticef("creating tinker config directory %s", t.ConfigDir)
		if err = os.MkdirAll(t.ConfigDir, os.ModePerm); err != nil {
			log.Errorf("error creating tinker config dir %s", t.ConfigDir)
			return err
		}
	}
	return nil
}

// Setup subnet using net.IPNet struct
func (t *TinkerConfiguration) SetupGlobalSubnet() (err error) {
	_, subnet, err := net.ParseCIDR(t.SubnetString)
	if err != nil {
		log.Errorf("error parsing subnet %s into ip.IPNet type: %s", t.SubnetString, err)
		return err
	} else {
		t.Subnet = *subnet
		return nil
	}
}

// Setup public keys for every user. If the key is found, the user is enabled.
func (t *TinkerConfiguration) SetupUserPublicKeys() {
	for count, user := range t.Users {
		userKey := path.Join(t.RepoDir, fmt.Sprintf("%s.pub", user.Username))
		if _, err := os.Stat(userKey); os.IsNotExist(err) {
			log.Warningf("pub key for user %s not found, disabling user", user.Username)
			t.Users[count].Enabled = false
		} else {
			log.Noticef("pub key for user %s found, enabling user", user.Username)
			t.Users[count].Enabled = true
			t.Users[count].Key = userKey
		}
	}
}

// Parse configuration from filename. The config file reside inside the key git repo.
// Key and Subnet are added to config structures during loading because they cannot be
// unmarshalled (Subnet) or need to be validated (Key).
// filename optional parameter is used to run tests avoiding creation of directories.
func (t *TinkerConfiguration) LoadConfig(filename ...string) (err error) {
	var configFilename string
	if filename[0] == "" {
		// Setup config directory
		if err := t.SetupConfigDir(); err != nil {
			return err
		}

		// Setup up git repo
		t.GitRepo = *flagGitRepo
		t.RepoDir = path.Join(t.ConfigDir, t.GitRepo)
		if _, err := os.Stat(t.RepoDir); os.IsNotExist(err) {
			// TODO: handle this inside git module
			log.Errorf("git repo %s does not exists. please clone it manually", t.RepoDir)
			return err
		}
		configFilename = path.Join(t.RepoDir, "config.json")
	} else {
		configFilename = filename[0]
	}

	// Read json config file
	configFile, err := os.Open(configFilename)
	if err != nil {
		log.Errorf("error opening configuration file %s/%s: %s", t.RepoDir, configFilename, err)
		return err
	}

	// Validate and decode json
	jsonParser := json.NewDecoder(configFile)
	if err = jsonParser.Decode(t); err != nil {
		log.Errorf("error loading data from configuration file %s: %s", filename, err)
		return err
	}

	// Setting up subnet
	if err = t.SetupGlobalSubnet(); err != nil {
		return err
	}

	// Setting up public keys keys
	t.SetupUserPublicKeys()
	return nil
}
