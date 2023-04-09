package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Firts string
	Age   int
}

func main() {
	u1 := user{
		Firts: "James",
		Age:   32,
	}

	u2 := user{
		Firts: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		Firts: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}
	fmt.Println(users)

	outt, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(outt))
	}

	// your code goes here
}

