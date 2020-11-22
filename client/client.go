package main

import (
	"fmt"
	"net/rpc"
)

type Degre struct {
	Student string
	Subject string
	Grade   float64
}

func client() {
	c, err := rpc.Dial("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		fmt.Println("1.- Agregar calififcacion")
		fmt.Println("2.- Obtener Promedio de un alumno")
		fmt.Println("3.- Obtener el promedio de todos los alumnos")
		fmt.Println("4.- Obtener promedio por materia")
		var opc int
		fmt.Scanln(&opc)
		switch opc {
		case 1:
			var student string
			var subject string
			var grade float64
			fmt.Println("Nombre del alumno: ")
			fmt.Scanln(&student)
			fmt.Println("Nombre de la materia: ")
			fmt.Scanln(&subject)
			fmt.Println("Calificacion: ")
			fmt.Scanln(&grade)
			var reply bool
			message := &Degre{student, subject, grade}
			err = c.Call("Server.Register", message, &reply)
			if err != nil {
				fmt.Println(err)
			}
			break
		case 2:
			var student string
			var average float64
			fmt.Println("Nombre del alumno: ")
			fmt.Scanln(&student)
			err = c.Call("Server.StudentAverage", student, &average)
			if err != nil {
				fmt.Println(err)
				break
			}
			fmt.Println("El promedio del alumno:", student, " es: ", average)
			break
		case 3:
			var average float64
			err = c.Call("Server.GeneralAverage", "", &average)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("El promedio de todos los alumnos es: ", average)
			break
		case 4:
			var subject string
			var average float64
			fmt.Println("Nombre de la  materia: ")
			fmt.Scanln(&subject)
			err = c.Call("Server.SubjectAverage", subject, &average)
			fmt.Println("El promedio de la materia: ", subject, " es: ", average)
		}

	}
}
func main() {

	client()

}
