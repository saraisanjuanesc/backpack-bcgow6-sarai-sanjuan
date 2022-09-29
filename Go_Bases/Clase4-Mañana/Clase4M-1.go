package main

import "fmt"

type ErrorSalary struct{
}
func (e *ErrorSalary)Error() string{
	return "error: el salario ingresado no alcanza el mínimo imponible"
}

func calcula(salary int)(err error){
	if salary < 150_000{
		err = &ErrorSalary{}
	}
	return
}

func main(){
	var salary int = 15000
	err := calcula(salary)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}else{
		fmt.Println("Debe Pagar impuesto")
	}
}