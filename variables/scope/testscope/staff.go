package testscope

var OverpaidLimit = 70000

type Employee struct {
	FirstName string
	LastName  string
	Salary    int
	FullTime  bool
}

type Office struct {
	AllStaff []Employee
}

func (o *Office) Show() []Employee {
	return o.AllStaff
}

func (o *Office) OverPaid() []Employee {
	var overPaid []Employee

	for _, x := range o.AllStaff {
		if x.Salary > OverpaidLimit {
			overPaid = append(overPaid, x)
		}
	}

	return overPaid
}
