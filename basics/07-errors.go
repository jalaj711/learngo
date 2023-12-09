package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Dividing by zero error")
	}
	return a / b, nil
}

type CustomError struct {
	user_id int
}

func (err CustomError) Error() string {
	return fmt.Sprintf("The customer with id %d is not allowed to perform this action", err.user_id)
}

func sendMessage(userid int) (success bool, err error) {
	if userid%3 == 0 {
		return false, CustomError{user_id: userid}
	}
	return true, nil
}

func main() {
	v, e := divide(4, 0)
	if e != nil {
		fmt.Printf("Error while dividing: %s\n", e.Error())
	} else {

		fmt.Printf("4/0 = %d\n", v)
	}
	v2, e2 := sendMessage(3)
	if e2 != nil {
		fmt.Printf("Error while sending message: %s\n", e2.Error())
	} else if v2 {
		fmt.Printf("Message sent succesfully")
	} else {
		fmt.Printf("Message not sent succesfully")
	}
}
