package target

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type CompanyInfo struct {
	Company        string `json:"Company"`
	AshbySlug      string `json:"AshbySlug,omitempty"`
	LeverSlug      string `json:"LeverSlug,omitempty"`
	Enabled        bool   `json:"Enabled"`
	FrequencyHours int    `json:"FrequencyHours,omitempty"`
}

type CompaniesConfig struct {
	Lever []CompanyInfo `json:"lever"`
	Ashby []CompanyInfo `json:"ashby"`
}

var DefaultCompanies []CompanyInfo
var LeverCompanies []CompanyInfo

func init() {
	loadCompaniesFromFile()
}

func loadCompaniesFromFile() {
	paths := []string{"companies.json", "./companies.json", "/app/companies.json", filepath.Join(os.Getenv("PWD"), "companies.json")}

	var data []byte
	var err error

	for _, p := range paths {
		data, err = os.ReadFile(p)
		if err == nil {
			break
		}
	}

	if err != nil {
		loadCompaniesFromEnv()
		return
	}

	var config CompaniesConfig
	if err := json.Unmarshal(data, &config); err != nil {
		loadCompaniesFromEnv()
		return
	}

	DefaultCompanies = config.Ashby
	LeverCompanies = config.Lever
}

func loadCompaniesFromEnv() {
	DefaultCompanies = []CompanyInfo{}
	LeverCompanies = []CompanyInfo{}

	envCompanies := os.Getenv("ASHBY_COMPANIES")
	if envCompanies != "" {
		var parsed []CompanyInfo
		if err := json.Unmarshal([]byte(envCompanies), &parsed); err == nil {
			DefaultCompanies = parsed
			return
		}
	}

	commaSep := os.Getenv("ASHBY_COMPANIES_COMMA")
	if commaSep != "" {
		parts := strings.Split(commaSep, ",")
		for _, p := range parts {
			kv := strings.Split(strings.TrimSpace(p), ":")
			if len(kv) == 2 {
				DefaultCompanies = append(DefaultCompanies, CompanyInfo{
					Company: kv[0], AshbySlug: kv[1], Enabled: true,
				})
			}
		}
	}
}

func GetEnabledCompanies() []CompanyInfo {
	companies := []CompanyInfo{}

	envCompanies := os.Getenv("ASHBY_COMPANIES")
	if envCompanies != "" {
		var parsed []CompanyInfo
		if err := json.Unmarshal([]byte(envCompanies), &parsed); err == nil {
			companies = append(companies, parsed...)
		}
	}

	commaSep := os.Getenv("ASHBY_COMPANIES_COMMA")
	if commaSep != "" {
		parts := strings.Split(commaSep, ",")
		for _, p := range parts {
			kv := strings.Split(strings.TrimSpace(p), ":")
			if len(kv) == 2 {
				companies = append(companies, CompanyInfo{
					Company: kv[0], AshbySlug: kv[1], Enabled: true,
				})
			}
		}
	}

	if len(companies) == 0 {
		for _, c := range DefaultCompanies {
			if c.Enabled {
				companies = append(companies, c)
			}
		}
	}

	return companies
}

func GetEnabledLeverCompanies() []CompanyInfo {
	companies := []CompanyInfo{}

	envCompanies := os.Getenv("LEVER_COMPANIES")
	if envCompanies != "" {
		var parsed []CompanyInfo
		if err := json.Unmarshal([]byte(envCompanies), &parsed); err == nil {
			companies = append(companies, parsed...)
		}
	}

	commaSep := os.Getenv("LEVER_COMPANIES_COMMA")
	if commaSep != "" {
		parts := strings.Split(commaSep, ",")
		for _, p := range parts {
			kv := strings.Split(strings.TrimSpace(p), ":")
			if len(kv) == 2 {
				companies = append(companies, CompanyInfo{
					Company: kv[0], LeverSlug: kv[1], Enabled: true,
				})
			}
		}
	}

	if len(companies) == 0 {
		for _, c := range LeverCompanies {
			if c.Enabled {
				companies = append(companies, c)
			}
		}
	}

	return companies
}

func GetAllCompanies() []CompanyInfo {
	return DefaultCompanies
}

func GetAllLeverCompanies() []CompanyInfo {
	return LeverCompanies
}

func GetDueCompanies(lastScraped map[string]*string, allCompanies []CompanyInfo) []CompanyInfo {
	var due []CompanyInfo
	for _, c := range allCompanies {
		if !c.Enabled {
			continue
		}
		last := lastScraped[c.AshbySlug]
		if last == nil {
			due = append(due, c)
			continue
		}
		if shouldScrape(*last, c.FrequencyHours) {
			due = append(due, c)
		}
	}
	return due
}

func shouldScrape(lastScraped string, frequencyHours int) bool {
	if lastScraped == "" {
		return true
	}
	t, err := time.Parse(time.RFC3339, lastScraped)
	if err != nil {
		return true
	}
	return time.Since(t) > time.Duration(frequencyHours)*time.Hour
}
