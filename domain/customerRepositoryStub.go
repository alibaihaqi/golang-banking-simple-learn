package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Name: "Hengki", City: "Jakarta", Zipcode: "12810", DateOfBirth: "2000-01-01", Status: "1"},
		{Name: "Jacky", City: "Bandung", Zipcode: "11012", DateOfBirth: "2000-01-01", Status: "1"},
	}

	return CustomerRepositoryStub{customers}
}
