package main

import (
	"fmt"
	"go.code.as/writeas.v2"
	"time"
)

func main() {
	
	// Following Write.as API protocol for authentication
	// Create a new client, log in, and retrieve the access token

	c := writeas.NewClient()

	u, err := c.LogIn("blahblah", "blahblah")

	access_token := u.AccessToken

	c.SetToken(access_token)
	
	// Now we create the post

	p, err := c.CreatePost(&writeas.PostParams{
		Title:   "Set up lunch with Paul",
		Content: "Remind yourself to respond to Paul about lunch tomorrow. Chipotle?!"})

	if err != nil {
		fmt.Println("Ooops! %v\n", err)
	} else {
		fmt.Println("Here is the post: https://write.as/" + p.ID + "\nIt will expire in 5 minutes!")
		
	// This is where the expiration magic (or lack thereof) comes in
	// Let the program sleep for an allotted time 
		
	time.Sleep(5 * time.Minute)
			
	//Once that expires, the program runs through to delete the post
		
	token := p.Token

	d := c.DeletePost(p.ID, token)

	if d != nil {
		fmt.Printf("%v", d)
	} else {
		fmt.Println("You deleted it!")
	}

}

