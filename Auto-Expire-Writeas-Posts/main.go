package main

import (
	"fmt"
	"go.code.as/writeas.v2"
	"time"
)

func main() {

	c := writeas.NewClient()

	u, err := c.LogIn("blahblah", "blahblah")

	access_token := u.AccessToken

	c.SetToken(access_token)

	p, err := c.CreatePost(&writeas.PostParams{
		Title:   "Set up lunch with Paul",
		Content: "Remind yourself to respond to Paul about lunch tomorrow. Chipotle?!"})

	if err != nil {
		fmt.Println("Ooops! %v\n", err)
	} else {
		fmt.Println("Here is the post: https://write.as/" + p.ID)
	}

	fmt.Println("It will expire in 5 minutes!")

	token := p.Token

	time.Sleep(5 * time.Minute)

	d := c.DeletePost(p.ID, token)

	if d != nil {
		fmt.Printf("%v", d)
	} else {
		fmt.Println("You deleted it!")
	}

}

//	p, err := c.CreatePost(&writeas.PostParams{
//		Title:   "Check this out!",
//		Content: "I am writing in go!",
//	})
//	if err != nil {
// Perhaps show err.Error()
//	}
//  fmt.Println(p)
// Save token for later, since it won't ever be returned again
//	token := p.Token

//}
