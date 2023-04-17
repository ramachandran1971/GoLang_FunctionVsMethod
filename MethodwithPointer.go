package main
import "fmt"
type Employee struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
}


func (emp  Employee)EmpInfo() {
    emp.FirstName="RamachandranSivarajan"
	fmt.Println(emp.FirstName, emp.LastName)
	fmt.Println(emp.Age)
	fmt.Println(emp.Email)
}
func (emp  *Employee)EEmpInfo() {
    emp.FirstName="RamachandranSivarajan"
	fmt.Println(emp.FirstName, emp.LastName)
	fmt.Println(emp.Age)
	fmt.Println(emp.Email)
}
func main() {

	ee := new(Employee)
	    ee.FirstName="Ram"
	    ee.LastName="Siva"
	    ee.Email="ram@gmail.com"
    	ee.Age= 45
    fmt.Println("==================================================")
    fmt.Println("Without pointer - No Change")	
    fmt.Println("==================================================")
	fmt.Println(ee)
	        ee.EmpInfo()
	fmt.Println(ee)
	fmt.Println("==================================================")
	fmt.Println("With pointer - Change")	
	fmt.Println("==================================================")
	fmt.Println(ee)
	        ee.EEmpInfo()
	fmt.Println(ee)
}

/*
output
==================================================
Without pointer - No Change
==================================================
&{Ram Siva ram@gmail.com 45}
RamachandranSivarajan Siva
45
ram@gmail.com
&{Ram Siva ram@gmail.com 45}
==================================================
With pointer - Change
==================================================
&{Ram Siva ram@gmail.com 45}
RamachandranSivarajan Siva
45
ram@gmail.com
&{RamachandranSivarajan Siva ram@gmail.com 45}

	 
*/