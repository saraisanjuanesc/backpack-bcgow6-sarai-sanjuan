package main

import "fmt"

type user struct {
	Name     string
	LastName string
	Age      int
	Email    string
	Password string
}

func updateName(us *user, newname string, newlastname string) {
	us.Name = newname
	us.LastName = newlastname
}

func updateAge(us *user, newage int) {
	us.Age = newage
}

func updateEmail(us *user, newemail string) {
	us.Email = newemail
}

func updatePassword(us *user, newpassword string) {
	us.Password = newpassword
}
func main() {
	var user1 user = user{
		Name:     "Martín",
		LastName: "Hernández",
		Age:      25,
		Email:    "martin.hz@gmail.com",
		Password: "658271!hjs",
	}

	fmt.Println("Los datos Originales son: ")
	fmt.Printf("Nombre: %s\nApellido: %s\nEdad: %d\nEmail: %s\nPassword: %s\n", user1.Name, user1.LastName, user1.Age, user1.Email, user1.Password)

	updateName(&user1, "José Juan", "Campos")
	fmt.Println("Se actualizó el Nombre y el apellido")

	updateAge(&user1, 24)
	fmt.Println("Se actualizó la Edad")

	updateEmail(&user1, "jose.campos@gmail.com")
	fmt.Println("Se actualizó la Edad")

	updatePassword(&user1, "das329jas!ojf@")
	fmt.Println("Se actualizó la Contraseña")

	fmt.Println("Los datos Nuevos son: ")
	fmt.Printf("Nombre: %s\nApellido: %s\nEdad: %d\nEmail: %s\nPassword: %s\n", user1.Name, user1.LastName, user1.Age, user1.Email, user1.Password)

}
