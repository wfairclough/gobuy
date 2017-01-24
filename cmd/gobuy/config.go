package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/wfairclough/gobuy"
)

// Global Flags available to all commands
var (
	flagConfig     = flag.String("config", path.Join(os.Getenv("HOME"), ".gobuy.json"), "A config yaml file for default flag options")
	flagShopDomain = flag.String("domain", "", "Your Shopify domain. Example: example.myshopify.com") // Required Flag or in Config
	flagApiKey     = flag.String("api-key", "", "The API Key for the Mobile App Sales Channel")       // Required Flag or in Config
	flagAppId      = flag.String("app-id", "", "The APP ID for the Mobile App Sales Channel")         // Required Flag or in Config
	flagAppName    = flag.String("app", "", "The name of this app. Example: gobuy")
)

// config defines the defealt configuration required by the gobuy command
type config struct {
	ShopDomain string `json:"shop_domain"`
	ApiKey     string `json:"api_key"`
	AppId      string `json:"app_id"`
	AppName    string `json:"app_name"`
}

// isValid checks if the configuration if valid for use by the subcommands
func (c config) isValid() bool {
	return c.ShopDomain != "" &&
		c.ApiKey != "" &&
		c.AppId != "" &&
		c.AppName != ""
}

// clientFromConfig creates a BuyClient from the configuration provided by cli
func clientFromConfig() (*gobuy.BuyClient, error) {
	if _, err := os.Stat(*flagConfig); os.IsNotExist(err) {
		// config file does not exist
		return gobuy.Client(*flagShopDomain, *flagAppName, *flagApiKey, *flagAppId), nil
	}
	file, err := os.Open(*flagConfig)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	cfgContent, err := cleanUpConfigData(file)
	if err != nil {
		return nil, err
	}
	cfg := &config{}
	err = json.Unmarshal([]byte(cfgContent), &cfg)
	if err != nil {
		return nil, err
	}
	if flagShopDomain != nil && *flagShopDomain != "" {
		cfg.ShopDomain = *flagShopDomain
	}
	if flagAppName != nil && *flagAppName != "" {
		cfg.AppName = *flagAppName
	}
	if flagApiKey != nil && *flagApiKey != "" {
		cfg.ApiKey = *flagApiKey
	}
	if flagAppId != nil && *flagAppId != "" {
		cfg.AppId = *flagAppId
	}
	if !cfg.isValid() {
		return nil, errors.New("Invalid configuration cannot run commands. Check the configuration file: " + *flagConfig)
	}
	return gobuy.Client(cfg.ShopDomain, cfg.AppName, cfg.ApiKey, cfg.AppId), nil
}

// Clean up any comment lines that are in the json
func cleanUpConfigData(r io.Reader) (string, error) {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return "", err
	}

	re := regexp.MustCompile(".*//.*")
	res := re.ReplaceAllString(string(b), "")
	return res, nil
}
