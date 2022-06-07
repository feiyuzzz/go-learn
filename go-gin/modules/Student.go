package entity

import "fmt"

type Student struct {
	Id   string
	Name string
	Age  int
	Sex  int
}

func ExistedStudent(Id int) (bool, error) {
	return true, nil
}

func GetStudent(id int) (*Student, error) {
	var stud Student
	return &stud, nil
}

func AddStudent(data map[string]interface{}) error {
	stud := Student{
		Id:   data["id"].(string),
		Name: data["name"].(string),
		Age:  data["age"].(int),
		Sex:  data["sex"].(int),
	}
	fmt.Printf("stud: %v\n", stud)
	return nil
}

func UpdateStudent(id int, data interface{}) error {
	return nil
}

func DeleteStudent(id int) error {
	return nil
}
