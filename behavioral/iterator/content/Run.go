package iterator

import "fmt"

func Run() {
	user1 := &User{
		Name: "You",
		Age:  18,
	}
	user2 := &User{
		Name: "Quy",
		Age:  18,
	}
	UserCollection := UserCollection{
		users: []*User{user1, user2},
	}
	iterator := UserCollection.CreateIterator()
	for iterator.HasNext() {
		user := iterator.GetNext()
		fmt.Printf("User is %+v\n", user)
	}

}
