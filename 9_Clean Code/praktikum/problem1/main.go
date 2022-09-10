/*
package main

	type user struct {
		id       int
		username int
		password int
	}

	type userservice struct {
		t []user
	}

	func (u userservice) getallusers() []user {
		return u.t
	}

	func (u userservice) getuserbyid(id int) user {
		for _, r := range u.t {
			if id == r.id {
				return r
			}
		}
		return user{}
	}
*/

package main

type User struct {
	id       int
	username string
	password string
}

type UserService struct {
	t []User
}

func (us UserService) GetAllUsers() []User {
	return us.t
}
func (us UserService) GetUserById(id int) User {
	for _, user := range us.t {
		if id == user.id {
			return user
		}
	}
	return User{}
}
