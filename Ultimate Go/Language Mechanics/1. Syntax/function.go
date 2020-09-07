package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type user struct {
	ID   int
	Name string
}

type updateStats struct {
	Modified int
	Duration float64
	Success  bool
	Message  string
}

func main() {
	// Retrieve the user profile
	u, err := retrieveUser("CHOI")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", *u)

	// Update user name.
	// if scope 밖에서는 아무것도 필요 없기 때문에, compact syntax를 사용 할 수 있다.
	if _, err := updateUser(u); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Updated user record for ID", u.ID)
}

func retrieveUser(name string) (*user, error) {
	r, err := getUser(name)
	if err != nil {
		return nil, err
	}

	var u user

	// 값을 안전한 call stack의 아래 쪽에 공유 하여, Unmarshal functino이 document를 읽고 초기화 할 수 있게 해주었다.
	err = json.Unmarshal([]byte(r), &u)

	/*
		다시 call stack으로 공유한다. 이 라인 때문에, 우리는 할당이 된 것을 알 수 있다.
		이전 스텝의 값은 스택이 아니라 힙 위에 있다.
	*/
	return &u, err
}

func getUser(name string) (string, error) {
	response := `{"ID": 101, "Name": "CHOI"}`
	return response, nil
}

func updateUser(u *user) (*updateStats, error) {
	// response simulates a JSON response.
	response := `{"Modified":1, "Duration":0.005, "Success" : true,
"Message": "updated"}`

	var us updateStats
	if err := json.Unmarshal([]byte(response), &us); err != nil {
		return nil, err
	}

	if us.Success != true {
		return nil, errors.New(us.Message)
	}

	return &us, nil
}
