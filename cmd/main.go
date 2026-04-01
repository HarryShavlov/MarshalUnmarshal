package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"-"`
	active   bool
}

func main() {
	user := User{
		ID:       1,
		Name:     "Гофер",
		Email:    "gopher@gophermate.com",
		Password: "I4mG0ph3R",
		active:   true,
	}

	// преобразуйте user в JSON формат
	out, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("ошибка при сериализации в json: %s\n", err.Error())
		return
	}

	var newUser User

	// десериаизуйте данные из JSON формата в переменную newUser
	// ...
	if err = json.Unmarshal(out, &newUser); err != nil {
		fmt.Printf("ошибка при десериализации в json: %s\n", err.Error())
		return
	}

	fmt.Println(newUser)
}
