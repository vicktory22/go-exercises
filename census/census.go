package census

type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{
		Name:    name,
		Age:     age,
		Address: address,
	}
}

func (r *Resident) HasRequiredInfo() bool {
	hasName := r.Name != ""
	hasStreet := r.Address["street"] != ""

	return hasName && hasStreet
}

func (r *Resident) Delete() {
	*r = *NewResident("", 0, map[string]string(nil))
}

func Count(residents []*Resident) int {
	residentCount := 0

	for _, resident := range residents {
		if !resident.HasRequiredInfo() {
			continue
		}

		residentCount++
	}

	return residentCount
}
