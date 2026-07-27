package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
    // Check for employee with the same ID
    // for _, emp := range m.Employees{
    //     if emp.ID == e.ID{
    //         return
    //     }
    // }
	m.Employees = append(m.Employees, e)
	return
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
    for i, e := range m.Employees{
        //If employee with ID exists it will be removed
        if e.ID == id{
            m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
            return
        }
    }
    return
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
    sum := 0.0
    amount := 0
    for _, e := range m.Employees{
        //If employee with ID exists it will be removed
        sum += e.Salary
        amount += 1
    }
    if amount != 0{
        return sum / float64(amount)
    }
    return 0
    
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
    for i := range m.Employees {
        if m.Employees[i].ID == id {
            return &m.Employees[i]
        }
    }
    return nil
}

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
