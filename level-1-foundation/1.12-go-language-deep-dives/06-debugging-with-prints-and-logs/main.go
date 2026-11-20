package main

import "fmt"

type Service struct {
	Host string
	Port int
}

// InspectVariable menghasilkan representasi %+v dan %#v dari nilai yang diberikan.
func InspectVariable(val any) (fieldFormat, syntaxFormat string) {
	// TODO: Gunakan fmt.Sprintf dengan %+v dan %#v
	return "", ""
}

func main() {
	s := Service{Host: "127.0.0.1", Port: 9000}
	f, _ := InspectVariable(s)
	fmt.Printf("Field format: %s\n", f)
}
