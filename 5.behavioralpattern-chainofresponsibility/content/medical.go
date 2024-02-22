package chainofresponsibility

import "fmt"

type Medical struct {
	next Department
}

func (m *Medical) Execute(p *Patient) {
	if p.isMedicineProvided {
		fmt.Println("Patient already provided doctor")
		m.next.Execute(p)
		return
	}
	fmt.Println("We are providing medicine to patient")
	p.isMedicineProvided = true
	m.next.Execute(p)
}

func (m *Medical) SetNext(next Department) {
	m.next = next
}
