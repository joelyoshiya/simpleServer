// A simple client that will call a hackerrank mock api and run a specific operation the data
// the endpoint has info about users and articles they've written
// the client will call the endpoint and determine which users have a submission count greater than a given threshold
// the client will then return the user id and the number of submissions for each user that meets the threshold

package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/joelyoshiya/go_rest_api_no_frameworks/dataStructs"
)

// define structures in memory on the server

// type Body struct {
// 	Page       int     `json:"page"`
// 	PerPage    int     `json:"per_page"`
// 	Total      int     `json:"total"`
// 	TotalPages int     `json:"total_pages"`
// 	Data       Authors `json:"data"`
// }

// type Author struct {
// 	ID              int       `json:"id"`
// 	Username        string    `json:"username"`
// 	About           string    `json:"about"`
// 	Submitted       int       `json:"submitted"`
// 	UpdatedAt       time.Time `json:"updated_at"`
// 	SubmissionCount int       `json:"submission_count"`
// 	CommentCount    int       `json:"comment_count"`
// 	CreatedAt       int       `json:"created_at"`
// }

// type Authors []Author

// define the endpoint - note, this is a paginated endpoint
const endpoint = "https://jsonmock.hackerrank.com/api/article_users" // can use query params to filter pages
// define a threshold for the number of submissions
const threshold = 10

// get the data from the endpoint, then print to stdout
func getAndPrint() error {
	currPage := 1
	totalPages := 1
	for currPage <= totalPages {
		// get the data from the endpoint
		resp, err := http.Get(endpoint + "?page=" + strconv.Itoa(currPage))
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		// read the response body into a byte array
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		// unmarsall the data into a Body type
		body := dataStructs.Body{}
		bp := &body
		json.Unmarshal(b, &bp)
		if err != nil {
			return err
		}
		if bp == nil {
			return errors.New("body is nil")
		}

		// set the total pages if this is the first page
		if currPage == 1 {
			totalPages = body.TotalPages
		}

		// iterate through the data and print the data that meets the threshold
		for _, author := range body.Data {
			if author.SubmissionCount > threshold {
				println(author.Username, author.SubmissionCount)
			}
		}
		// increment the page
		currPage++
	}
	return nil
}

func main() {
	err := getAndPrint()
	if err != nil {
		log.Fatal(err)
	}
}
