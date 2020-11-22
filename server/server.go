package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Degre struct {
	Student string
	Subject string
	Grade   float64
}
type Server struct{}

var subjects = make(map[string]map[string]float64)
var students = make(map[string]map[string]float64)

func (s *Server) Register(degre Degre, reply *bool) error {
	if subjects[degre.Subject] != nil {
		if subjects[degre.Subject][degre.Student] != 0 {
			return errors.New("Esta calificacion ya existe")
		}
		subjects[degre.Subject][degre.Student] = degre.Grade
	} else {
		subjects[degre.Subject] = map[string]float64{degre.Student: degre.Grade}
	}
	if students[degre.Student] != nil {
		students[degre.Student][degre.Subject] = degre.Grade

	} else {
		students[degre.Student] = map[string]float64{degre.Subject: degre.Grade}

	}
	j, err := json.Marshal(subjects)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j))
	return nil
}
func (s *Server) GeneralAverage(name string, reply *float64) error {
	generalAverage := 0.0
	i := 0.0
	for student := range students {
		studentAverage := 0.0
		j := 0.0
		for _, grade := range students[student] {
			studentAverage = studentAverage + grade
			j++
		}
		i++
		generalAverage = generalAverage + studentAverage/j
	}
	*reply = generalAverage / i
	return nil

}
func (s *Server) StudentAverage(student string, reply *float64) error {
	if students[student] == nil {
		return errors.New("Alumno no encontrado")
	} else {
		average := 0.0
		i := 0.0
		for _, grade := range students[student] {
			average = average + grade
			i++
		}
		*reply = average / i

	}
	return nil

}
func (s *Server) SubjectAverage(subject string, reply *float64) error {

	if subjects[subject] == nil {
		return errors.New("Materia no encontrada")
	} else {
		average := 0.0
		i := 0.0
		for _, grade := range subjects[subject] {
			average = average + grade
			i++
		}
		*reply = average / i

	}
	return nil
}

func server() {
	rpc.Register(new(Server))
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)

}
