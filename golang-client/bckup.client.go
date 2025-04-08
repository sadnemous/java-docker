package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Define a struct to represent a movie
type Movie struct {
	Title  string  `json:"title"`
	Year   int     `json:"year"`
	Genre  string  `json:"genre"`
	Rating float64 `json:"rating"`
}

// Define a struct to represent a list of movies
type MovieList struct {
	Movies []Movie `json:"movies"`
}

// Define a struct to represent a single movie
type MovieDetails struct {
	Title    string  `json:"title"`
	Year     int     `json:"year"`
	Genre    string  `json:"genre"`
	Rating   float64 `json:"rating"`
	Synopsis string  `json:"synopsis"`
}

// Client is a simple HTTP client that can be used to make requests to a server.
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

// NewClient creates a new Client with the given base URL and timeout.
func NewClient(baseURL string, timeout time.Duration) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get makes a GET request to the given endpoint and returns the response.
func (c *Client) Get(ctx context.Context, endpoint string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	return resp, nil
}

// main function to demonstrate the usage of the Client
func main() {
	client := NewClient("http://localhost:8080", 10*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, "/movies")
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response status: %s\n", resp.Status)
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Expected status OK, got %s", resp.Status)
	}
	// Further processing of the response can be done here
	// For example, reading the response body, unmarshalling JSON, etc.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	} else {
		fmt.Printf("Response body: %s\n", string(body))
	}
	// // Handle the response as needed
	// // For example, unmarshalling JSON into a struct, etc.
	// var movies []Movie
	// err = json.Unmarshal(body, &movies)
	// if err != nil {
	// 	log.Fatalf("Error unmarshalling response body: %v", err)
	// } else {
	// 	fmt.Printf("Movies: %+v\n", movies)
	// }
	// // Handle the movies slice as needed
	// // For example, printing the movie titles, etc.
	// for _, movie := range movies {
	// 	fmt.Printf("Movie title: %s\n", movie.Title)
	// }
}
