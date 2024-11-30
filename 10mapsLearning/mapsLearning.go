package main

import (
	"errors"
	"fmt"
)

/*
Maps are similar to JavaScript objects, Python dictionaries, and Ruby hashes. Maps are a data structure that provides key->value mapping.
The zero value of a map is nil.

ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21 // overwrites 24

ages = map[string]int{
	"John": 37,
	"Mary": 21,
}

*/
type user struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error){
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}
	for i:=0; i<len(names); i++{
		userMap[names[i]] = user{name: names[i], phoneNumber: phoneNumbers[i]}
	}
	return userMap, nil
}

func printUserMap(names []string, phoneNumbers []int){
	x,y := getUserMap(names, phoneNumbers)
	if(y != nil){
		fmt.Println(y)
	}else{
		fmt.Println(x)
	}	
}

/*
MUTATIONS

INSERT AN ELEMENT
m[key] = elem

GET AN ELEMENT
elem = m[key]

DELETE AN ELEMENT
delete(m, key)

CHECK IF A KEY EXISTS
elem, ok := m[key]

NOTE ON PASSING MAPS
Like slices, maps are also passed by reference into functions. This means that when a 
map is passed into a function we write, we can make changes to the original, we don't have a copy.

*/
type user1 struct{
	name                 string
	number               int
	scheduledForDeletion bool
}

func deleteIfNecessary(users map[string]user1, name string) (deleted bool,err error){
	existingUser, ok := users[name]
	if !ok {
		return false, nil
	}
	if existingUser.scheduledForDeletion {
        delete(users, name)
        return true, nil
    }
	return false, nil
}

func printDeleteIfNecessary(users map[string]user1, name string){
	x, _:= deleteIfNecessary(users, name)
	fmt.Println(x)
}

/*
KEY TYPES

Any type can be used as the value in a map, but keys are more restrictive.

As mentioned earlier, map keys may be of any type that is comparable. The language spec defines this precisely, but in short, 
comparable types are boolean, numeric, string, pointer, channel, and interface types, and structs or arrays that contain only 
those types. Notably absent from the list are slices, maps, and functions; these types cannot be compared using ==, 
and may not be used as map keys.
*/

func getCounts(userIds []string) map[string]int{
	counts := make(map[string]int)
	for _, userId := range userIds{
		count := counts[userId]
		count++
		counts[userId] = count
	}
	return counts
}

func printGetCounts(userIds []string){
	x := getCounts(userIds)
	fmt.Println(x)
}

//NESTED
func getNameCounts(names []string) map[rune]map[string]int{
	counts := make(map[rune]map[string]int)
	for _, name := range names{
		if name == ""{
			continue
		}
		firstChar := rune(name[0])
		_, ok := counts[firstChar]
		if !ok {
			counts[firstChar] = make(map[string]int)
		}
		counts[firstChar][name]++
	}

	return counts
}

func printNameCounts(names []string){
	fmt.Println("#######printNameCounts#######")
	x := getNameCounts(names)
	fmt.Println(x)
	fmt.Println("#######printNameCounts#######")
}

func main() {
	printUserMap( []string{"Eren", "Armin", "Mikasa"}, []int{14355550987, 98765550987, 18265554567})
	printDeleteIfNecessary(map[string]user1{"Erwin": {"Erwin", 14355550987, true}}, "Erwin")

	printGetCounts([]string{"cersei", "jaime", "cersei"})

	printNameCounts([]string{
		"Grant",
		"Eduardo",
		"Peter",
		"Matthew",
		"Matthew",
		"Matthew",
		"Peter",
		"Peter",
		"Henry",
		"Parker",
		"Parker",
		"Parker",
		"Collin",
		"Hayden",
		"George",
		"Bradley",
		"Mitchell",
		"Devon",
		"Ricardo",
		"Shawn",
		"Taylor",
		"Nicolas",
		"Gregory",
		"Francisco",
		"Liam",
		"Kaleb",})
}