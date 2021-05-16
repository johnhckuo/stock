package test

import (
	"log"
	"net"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func init() {

	err := godotenv.Load("test.env")
	if err != nil {
		log.Fatal("Error loading test.env file")
	}

	timeout := time.Second
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
	if err != nil {
		log.Fatalf("Connecting error: %v", err.Error())
	}
	if conn != nil {
		defer conn.Close()
		log.Println("Opened", net.JoinHostPort(host, port))
	}
}

func BenchmarkCreateUser(b *testing.B) {

	resp, err := http.Post("http://127.0.0.1:8080/users/john", "application/json", nil)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	assert.Equal(b, 201, resp.StatusCode, "should be 201")
	/*
		defer resp.Body.Close()
		//Read the response body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		sb := string(body)
		log.Printf(sb)
	*/
}

func BenchmarkAddTimeSlots(b *testing.B) {
	b.Log("test")
}

func BenchmarkGetTimeSlots(b *testing.B) {
	b.Log("hi")
}
func BenchmarkDeleteTimeSlot(b *testing.B) {
	b.Log("hi")
}
