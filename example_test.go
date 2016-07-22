package fullstory_test

import (
	"fmt"
	"log"

	"github.com/nishanths/fullstory"
)

func Example_usage() {
	cfg := fullstory.Config{"your API token"}
	client := fullstory.NewClient(cfg)

	s, err := client.Sessions(15, "foo", "hikingfan@gmail.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
