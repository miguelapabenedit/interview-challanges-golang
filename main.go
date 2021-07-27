package main

import (
	"fmt"

	digitalengine "github.com/miguelapabenedit/interview-challanges-golang/digital-engine"
)

func main() {
	digitalengine.Problem1(1001)
}
func Solution(arr []int32) {
	var min int64
	var max int64
	var sum int64

	for i := 0; i < len(arr); i++ {
		sum = sum + int64(arr[i])

		if int64(arr[i]) < min || min == 0 {
			min = int64(arr[i])
		}
		if int64(arr[i]) > max {
			max = int64(arr[i])
		}
	}

	fmt.Printf("%v %v", sum-max, sum-min)
}

// type Passager struct {
// 	Name             string
// 	BookingReference string
// 	Loyalty          int
// }

// type PassagerGroup struct {
// 	Passagers        []Passager
// 	BookingReference string
// 	Loyalty          int
// }

// func SortPassagersByLoyalty(p []Passager) []Passager {
// 	groupsLoyalty := make(map[string]int)
// 	groups := make(map[string][]Passager)

// 	for _, v := range p {
// 		if br, ok := groupsLoyalty[v.BookingReference]; ok {
// 			if v.Loyalty < br {
// 				groupsLoyalty[v.BookingReference] = v.Loyalty
// 			}
// 		} else {
// 			groupsLoyalty[v.BookingReference] = v.Loyalty
// 		}

// 		if g, ok := groups[v.BookingReference]; ok {
// 			groups[v.BookingReference] = append(g, v)
// 		} else {
// 			groups[v.BookingReference] = []Passager{v}
// 		}
// 	}

// 	for k, v := range groupsLoyalty {
// 		for i := 0; i < len(groupsLoyalty); i++ {

// 		}
// 	}

// 	return nil
// }

// type Person struct {
// 	Name       string
// 	Height     string
// 	Hair_color string
// 	Skin_color string
// 	Eye_color  string
// 	Birth_year string
// 	Gender     string
// 	Homeworld  string
// }

// var persons = []*Person{}

// func MainChannels() {
// 	var wg = sync.WaitGroup{}

// 	consummer := GenerateHandler(&wg)

// 	for i := 0; i <= 5; i++ {
// 		wg.Add(1)
// 		go Worker(consummer, i)
// 	}

// 	wg.Wait()
// 	fmt.Println(persons)
// }

// func GenerateHandler(wg *sync.WaitGroup) chan *Person {
// 	c := make(chan *Person)

// 	go func(wg *sync.WaitGroup) {
// 		for v := range c {
// 			persons = append(persons, v)
// 			wg.Done()
// 		}
// 	}(wg)

// 	return c
// }

// func Worker(c chan *Person, id int) {
// 	fmt.Println("searching", id)
// 	resp, err := http.Get("https://swapi.dev/api/people/" + strconv.Itoa(id) + "/")

// 	if err != nil {
// 		fmt.Println(err)
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)

// 	if resp.StatusCode == 200 {
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		person := Person{}

// 		err = json.Unmarshal(body, &person)

// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println("searching Done", id)
// 		c <- &person
// 	} else {
// 		fmt.Println(resp.StatusCode)
// 		fmt.Printf("body ,request id %v\n", id)
// 	}
// }
