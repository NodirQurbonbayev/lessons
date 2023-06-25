package Hard

import "fmt"
type Users struct{
	Name,Email string
	Friend []Users
}
func (s *Users) AddFriend(friend Users) {
	s.Friend=append(s.Friend,friend)
}
func (s *Users) RemoveFriend(friend Users) {
	var index int 
	for i:=0; i<len(s.Friend); i++{
		if s.Friend[i].Name==friend.Name {
			index=i
		}
		s.Friend=append(s.Friend[:index],s.Friend[index+1:]...)
	}
}
func (s Users) Display() {
	fmt.Println("Name",s.Name)
	fmt.Println("Email",s.Email)
	fmt.Println("Do'stlar")
	for _, v := range s.Friend {
		fmt.Println("\t",v.Name)
	}

}