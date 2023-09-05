package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    amazon, err := UnmarshalAmazon(bytes)
//    bytes, err = amazon.Marshal()

import "encoding/json"

func UnmarshalAmazon(data []byte) (Amazon, error) {
	var r Amazon
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Amazon) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Amazon struct {
	Error   interface{} `json:"error"`
	Hits    int64       `json:"hits"`
	Facets  Facets      `json:"facets"`
	Content Content     `json:"content"`
	Jobs    []Jobs      `json:"jobs"`
}

type Contents struct {
	Sidebar       Facets `json:"sidebar"`
	SearchResults Facets `json:"search_results"`
}

type Facets struct {
}

type Jobs struct {
	BasicQualifications     string             `json:"basic_qualifications"`
	BusinessCategory        string             `json:"business_category"`
	City                    *string            `json:"city"`
	CompanyName             string             `json:"company_name"`
	CountryCode             CountryCode        `json:"country_code"`
	Description             string             `json:"description"`
	DepartmentCostCenter    interface{}        `json:"department_cost_center"`
	DescriptionShort        string             `json:"description_short"`
	DisplayDistance         interface{}        `json:"display_distance"`
	ID                      string             `json:"id"`
	IDIcims                 string             `json:"id_icims"`
	IsIntern                interface{}        `json:"is_intern"`
	IsManager               interface{}        `json:"is_manager"`
	JobCategory             string             `json:"job_category"`
	JobFamily               string             `json:"job_family"`
	JobFunctionID           interface{}        `json:"job_function_id"`
	JobPath                 string             `json:"job_path"`
	JobScheduleType         JobScheduleType    `json:"job_schedule_type"`
	Location                string             `json:"location"`
	NormalizedLocation      *string            `json:"normalized_location"`
	OptionalSearchLabels    []string           `json:"optional_search_labels"`
	PostedDate              string             `json:"posted_date"`
	PreferredQualifications *string            `json:"preferred_qualifications"`
	PrimarySearchLabel      *string            `json:"primary_search_label"`
	SourceSystem            SourceSystem       `json:"source_system"`
	State                   *string            `json:"state"`
	Title                   string             `json:"title"`
	UniversityJob           interface{}        `json:"university_job"`
	UpdatedTime             string             `json:"updated_time"`
	URLNextStep             string             `json:"url_next_step"`
	Team                    map[string]*string `json:"team"`
}

type CountryCode string

const (
	Ind CountryCode = "IND"
)

type JobScheduleType string

const (
	FullTimes JobScheduleType = "full-time"
)

type SourceSystem string

const (
	Hvh        SourceSystem = "HVH"
	JobCreator SourceSystem = "JobCreator"
)
