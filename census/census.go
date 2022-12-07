package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	residentInfos := &Resident{
		Name: name,
		Age: age,
		Address: address,
	}
	return residentInfos
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	rName := r.Name
	rAddress := r.Address

	switch {
	case rName == "": 
		return false
	case rAddress == nil:
		return false
	case rAddress["street"] == "":
		return false
	}
	return true
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Name = ""
	r.Age = 0
	r.Address = nil
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	countResidents := 0 
	for i := 0; i < len(residents); i++ {
		resident := residents[i]
		allFieldsFilled := resident.Name != "" && resident.Age != 0 && resident.Address != nil

		if allFieldsFilled {
			countResidents ++
		}
	}
	return countResidents
}
