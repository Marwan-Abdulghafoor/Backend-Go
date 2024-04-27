package utils

type ResponseDTO struct {
	Success     bool        `json:"success"`
	Message     string      `json:"message"`
	ProfileInfo ProfileInfo `json:"profile"`
}

type ProfileInfo struct {
	Id               int        `json:"id"`
	FirstName        string     `json:"firstName"`
	LastName         string     `json:"lastName"`
	Email            string     `json:"email"`
	Password         string     `json:"password"`
	PhoneNumber      string     `json:"phoneNumber"`
	AddressDTO       AddressDTO `json:"addressDTO"`
	Dob              string     `json:"dob"`
	Sex              string     `json:"sex"`
	BloodType        string     `json:"bloodType"`
	Chronic          string     `json:"chronic"`
	Allergies        string     `json:"allergies"`
	Surgeries        string     `json:"surgeries"`
	Medications      string     `json:"medications"`
	Immunization     string     `json:"immunization"`
	FamilyHistory    string     `json:"familyHistory"`
	PhysicianName    string     `json:"physicianName"`
	PhysicianContact string     `json:"physicianContact"`
	IceInstructions  string     `json:"emergencyInstructions"`
	EmergencyContact string     `json:"emergencyContact"`
}

type AddressDTO struct {
	StreetAddress string `json:"streetAddress"`
	City          string `json:"city"`
	State         string `json:"state"`
	ZipCode       string `json:"zipCode"`
}

type StandardIdRequest struct {
	Id int `json:"id"`
}

type EditRequest struct {
	Id          int         `json:"id"`
	ProfileInfo ProfileInfo `json:"profile"`
}
