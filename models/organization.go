package models

import "gopkg.in/mgo.v2/bson"

type Organization struct {
	OrganizationID bson.ObjectId `json:"organizationID" bson:"_id"`
	OrgTitle       string        `json:"orgtitle"`
	OrgWebsite     string        `json:"orgwebsite"`
	StreetAddress  string        `json: "street_address"`
	City           string        `json: "city"`
	State          string        `json: "state"`
	ZipCode        string        `json: "zip_code"`
	Phone          string        `json: "phone"`
	Email          string        `json: "email"`
	ActivityDesc   string        `json: "activity_desc"`
	Lat            float64       `json: "lat"`
	Long           float64       `json: "long"`
	HasShadow      bool          `json: "has_shadow"`
	HasCost        bool          `json: "has_cost"`
	HasTransport   bool          `json: "has_transport"`
	Under18        bool          `json: "under_18"`
	CareerEmp      []string      `json: "career_emp"`
	GradeLevels    []int         `json: "grade_levels"`
}

type UpdateOrganization struct {
	OrgTitle      string   `json:"org_title"`
	OrgWebsite    string   `json:"org_website"`
	StreetAddress string   `json: "street_address"`
	City          string   `json: "city"`
	State         string   `json: "state"`
	ZipCode       string   `json: "zip_code"`
	Phone         string   `json: "phone"`
	Email         string   `json: "email"`
	ActivityDesc  string   `json: "activity_desc"`
	Lat           float64  `json: "lat"`
	Long          float64  `json: "long"`
	HasShadow     bool     `json: "has_shadow"`
	HasCost       bool     `json: "has_cost"`
	HasTransport  bool     `json: "has_transport"`
	Under18       bool     `json: "under_18"`
	CareerEmp     []string `json: "career_emp"`
	GradeLevels   []int    `json: "grade_levels"`
}

// [{"OrgID": 1, "OrgTitle": "Pre-Health Professions Advising Program",
// "OrgWebsite": "www.uidaho.edu/pre-health", "StreetAddress": "875 Perimeter Drive",
// "City": "Moscow", "State": "Idaho", "ZipCode": "83844-2436", "Phone": "(208) 885-5809",
//  "Email": "pre-health@uidaho.edu",
//  "ActivityDesc": "The Pre-Health Advising Program at the University of Idaho serves as a resource
//  for students and alumni, from all majors, who are exploring graduate and professional programs in healthcare.
//    Services provided include assistance with career exploration, prerequisite course sequencing, advice for
//     resume building and entrance exam preparation, and support with the application and interview process. "
// 	, "Lat": 46.727471200000004, "Long": -117.02390190000001,
// 	"HasShadow": false, "HasCost": false, "HasTransport": false, "Under18": false,
// 	"CareerEmp": ["Medicine", "Dentistry", "Pharmacy", "Public Health", "Generic Health Sciences",
// 	"Allied Health", "STEM"],
// 	"GradeLevels": [9, 10, 11, 12]},
