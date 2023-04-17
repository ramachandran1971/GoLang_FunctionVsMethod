//Function vs. Method
package main
import "fmt"
type Employee struct {
	FirstName 		string
               LastName 		string 
               Email 			string
	Age                        	int

}
// Function
func  EmpInfo(e Employee)string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	
	return "----------------------"
}
// Methods
func (ee Employee)EEmpInfo() {
	fmt.Println(ee.FirstName, ee.LastName)
	fmt.Println(ee.Age)
	fmt.Println(ee.Email)
}
func main() {
       	e := Employee{
		FirstName: "Mark",
		LastName:"Jones",
		Email:     "mark@gmail.com",
		Age:25,
	}
	fmt.Println(EmpInfo(e))   // call - function
            
             // using “new” keyword
	ee := new(Employee)
	ee.FirstName="Ram"
	ee.LastName="Siva"
	ee.Email="ram@gmail.com"
	ee.Age= 45
	
	ee.EEmpInfo()     / call - Methos
}
/*
Output

Mark Jones
25
mark@gmail.com
----------------------
Ram Siva
45
ram@gmail.com  */
