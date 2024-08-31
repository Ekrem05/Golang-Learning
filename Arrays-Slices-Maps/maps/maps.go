package maps

import "fmt"

func main() {

	var websites = map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites)

	aws:=websites["Amazon Web Services"];
	fmt.Println(aws)

	websites["linkedin"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites,"Google")
	fmt.Println(websites)

}