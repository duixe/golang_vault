package main

import "fmt"

type ContactInfo struct {
	email string
	zipCode int
}

type Person struct {
	firstName string
	lastName string
	contact ContactInfo
}

func (p *Person) updateName (newName string) {
	(*p).firstName = newName;
}

func (p Person) print () {
	fmt.Printf("person is: %+v", p);
}

func main()  {
	jude := Person {
		firstName: "Emmanuel",
		lastName: "Akomaning",
		contact: ContactInfo {
			email: "emma@gmail.com",
			zipCode: 2345,
		},
	};

	refrencedJude := &jude;
	refrencedJude.updateName("julian");
	refrencedJude.print();

	fmt.Println("person: ", jude);
}