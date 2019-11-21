package mock

import "fmt"

type User struct {
	List []int
}

func (user *User) String() string {
	return fmt.Sprintf("User.List = %v", user.List)
}

func (user *User) Put(url string, form map[string]int) string {
	user.List[0] = form["id"]
	return "ok"
}

func (user *User) MyPost(url string, form map[string]int) string {
	user.List = append(user.List, form["id"], form["id"])
	return "ok"
}

func (user *User) MyGet(url string) string {
	return fmt.Sprintf("%v", user.List)
}
