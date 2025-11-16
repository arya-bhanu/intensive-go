package main

import "fmt"

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
	Indexes   map[int]int
}

func (m *Manager) AddEmployee(e Employee) {
	if m.Indexes == nil {
		m.Indexes = make(map[int]int)
	}
	_, ok := m.Indexes[e.ID]
	if ok {
		fmt.Println(fmt.Errorf("duplicate ID of %v", e.ID))
		return
	}
	m.Employees = append(m.Employees, e)
	m.Indexes[e.ID] = len(m.Employees) - 1
}

func (m *Manager) RemoveEmployee(id int) {
	_, ok := m.Indexes[id]
	if !ok {
		fmt.Printf("data not found for id: %v\n", id)
		return
	}
	delete(m.Indexes, id)
	var newEmployees []Employee
	var newMapIndex = make(map[int]int)
	i := 0
	for _, val := range m.Indexes {
		newEmployees = append(newEmployees, m.Employees[val])
		newMapIndex[m.Employees[val].ID] = i
		i += 1
	}
	m.Employees = newEmployees
	m.Indexes = newMapIndex
}

func (m *Manager) GetAverageSalary() float64 {
	totalSalary := 0.0
	employeesCount := len(m.Employees)
	for _, m := range m.Employees {
		totalSalary += m.Salary
	}
	return totalSalary / float64(employeesCount)
}

func (m *Manager) FindEmployeeByID(id int) *Employee {
	return &m.Employees[m.Indexes[id]]
}
