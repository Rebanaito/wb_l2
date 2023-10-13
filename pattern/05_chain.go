package main

import "fmt"

type team interface {
	interview(*intern)
	setNext(team)
}

type recruiter struct {
	next team
}

func (r recruiter) interview(i *intern) {
	if i.passedAll {
		fmt.Println(i.name, "discussed the contract details with the recruiter and got a job!")
	} else {
		if i.nice {
			fmt.Println(i.name, "was nice and got invited for the technical interview")
			r.next.interview(i)
		} else {
			fmt.Println(i.name, "was not nice and didn't get invited for the technical interview")
		}
	}
}

func (r *recruiter) setNext(next team) {
	r.next = next
}

type teamlead struct {
	next team
}

func (l teamlead) interview(i *intern) {
	if i.smart {
		fmt.Println(i.name, "succesfully passed the technical interview")
		l.next.interview(i)
	} else {
		fmt.Println(i.name, "could not pass the technical interview")
	}
}

func (l *teamlead) setNext(next team) {
	l.next = next
}

type securityDept struct {
	next team
}

func (s securityDept) interview(i *intern) {
	if i.lawAbiding {
		fmt.Println(i.name, "has passed the background check")
		i.passedAll = true
		s.next.interview(i)
	} else {
		fmt.Println(i.name, "has failed the background check")
	}
}

func (s *securityDept) setNext(next team) {
	s.next = next
}

type intern struct {
	name       string
	lawAbiding bool
	smart      bool
	nice       bool
	passedAll  bool
}

// func main() {
// 	hr := &recruiter{}
// 	lead := &teamlead{}
// 	hr.setNext(lead)
// 	security := &securityDept{}
// 	lead.setNext(security)
// 	security.setNext(hr)
// 	goodGuy := &intern{name: "John", lawAbiding: true, smart: true, nice: true}
// 	sketchyGuy := &intern{name: "Michael", lawAbiding: false, smart: true, nice: true}
// 	weakGuy := &intern{name: "Phillip", lawAbiding: true, smart: false, nice: true}
// 	rudeGuy := &intern{name: "Sean", lawAbiding: true, smart: true, nice: false}
// 	fmt.Println("Sean applies for the job")
// 	hr.interview(rudeGuy)
// 	fmt.Println()
// 	fmt.Println("Phillip applies for the job")
// 	hr.interview(weakGuy)
// 	fmt.Println()
// 	fmt.Println("Michael applies for the job")
// 	hr.interview(sketchyGuy)
// 	fmt.Println()
// 	fmt.Println("John applies for the job")
// 	hr.interview(goodGuy)
// }
