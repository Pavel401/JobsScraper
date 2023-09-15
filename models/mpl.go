package models

// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    mPL, err := UnmarshalMPL(bytes)
//    bytes, err = mPL.Marshal()

import "encoding/json"

func UnmarshalMPL(data []byte) (MPL, error) {
	var r MPL
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *MPL) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type MPL struct {
	Status  string  `json:"status"`
	Message Message `json:"message"`
}

type Message struct {
	Jobscount int64  `json:"jobscount"`
	Jobs      []Data `json:"jobs"`
}

type Data struct {
	ID                     string                `json:"id"`
	DesignationDisplayName string                `json:"designation_display_name"`
	ExperienceFromNum      int64                 `json:"experience_from_num"`
	ExperienceToNum        int64                 `json:"experience_to_num"`
	CreatedOn              string                `json:"created_on"`
	OfficelocationShowArr  OfficelocationShowArr `json:"officelocation_show_arr"`
	OfficelocationArr      []OfficelocationArr   `json:"officelocation_arr"`
	JobPostingOn           int64                 `json:"job_posting_on"`
	Department             string                `json:"department"`
	DepartmentID           string                `json:"department_id"`
	EmpType                EmpType               `json:"emp_type"`
	EmpTypeID              EmpTypeID             `json:"emp_type_id"`
	FunctionalArea         string                `json:"functional_area"`
	FunctionalAreaID       string                `json:"functional_area_id"`
	Title                  string                `json:"title"`
	ToolTipLocations       OfficelocationShowArr `json:"tool_tip_locations"`
	Timezone               Timezone              `json:"timezone"`
}

type EmpType string

const (
	Fte    EmpType = "FTE"
	Intern EmpType = "Intern"
)

type EmpTypeID string

const (
	The5Ecbc7079278A EmpTypeID = "5ecbc7079278a"
	The5Ecbc70Eb2E96 EmpTypeID = "5ecbc70eb2e96"
)

type OfficelocationArr string

const (
	A63Ad69B880B6D   OfficelocationArr = "a63ad69b880b6d"
	The5Ecbc9Fa9E511 OfficelocationArr = "5ecbc9fa9e511"
	The6098D53F7Fea0 OfficelocationArr = "6098d53f7fea0"
	The609A31E3045F1 OfficelocationArr = "609a31e3045f1"
)

type OfficelocationShowArr string

const (
	BangaloreKarnatakaIndia OfficelocationShowArr = "Bangalore, Karnataka, India "
	BengaluruKarnatakaIndia OfficelocationShowArr = "Bengaluru, Karnataka, India "
	PuneMaharashtraIndia    OfficelocationShowArr = "Pune, Maharashtra, India "
)

type Timezone string

const (
	AsiaJakarta Timezone = "Asia/Jakarta"
)
