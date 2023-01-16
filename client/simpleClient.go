// A simple client that will call a hackerrank mock api and run a specific operation the data
// the endpoint has info about users and articles they've written
// the client will call the endpoint and determine which users have a submission count greater than a given threshold
// the client will then return the user id and the number of submissions for each user that meets the threshold

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// define structures in memory on the server

type Body struct {
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
	Total      int     `json:"total"`
	TotalPages int     `json:"total_pages"`
	Data       Authors `json:"data"`
}

type Author struct {
	ID              int       `json:"id"`
	Username        string    `json:"username"`
	About           string    `json:"about"`
	Submitted       int       `json:"submitted"`
	UpdatedAt       time.Time `json:"updated_at"`
	SubmissionCount int       `json:"submission_count"`
	CommentCount    int       `json:"comment_count"`
	CreatedAt       int       `json:"created_at"`
}

type Authors []Author

// define the endpoint - note, this is a paginated endpoint
const endpoint = "https://jsonmock.hackerrank.com/api/article_users" // can use query params to filter pages
// define a threshold for the number of submissions
const threshold = 10

// manual ummarshalling for body
func (b *Body) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "page" {
			b.Page, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "per_page" {
			b.PerPage, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "total" {
			b.Total, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "total_pages" {
			b.TotalPages, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "data" {
			// define an Authors type
			// for each data entry, unmarshal the json into an Author type
			// append the Author to the Authors type

			// first, unmarshal the json into a map of strings
			var rawAuthors map[string]string
			err := json.Unmarshal([]byte(v), &rawAuthors)
			if err != nil {
				return err
			}
			// now, iterate over the map and unmarshal each entry into an Author type
			for _, v := range rawAuthors {
				var a Author
				var ap = &a
				err := ap.UnmarshalJSON([]byte(v))
				if err != nil {
					return err
				}
				b.Data = append(b.Data, *ap)
			}
		}
	}
	return nil
}

// manual unmashalling of json: https://ukiahsmith.com/blog/go-marshal-and-unmarshal-json-with-time-and-url-data/
func (a *Author) UnmarshalJSON(j []byte) error {
	var rawStrings map[string]string

	err := json.Unmarshal(j, &rawStrings)
	if err != nil {
		return err
	}

	for k, v := range rawStrings {
		if strings.ToLower(k) == "id" {
			a.ID, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "username" {
			a.Username = v
		}
		if strings.ToLower(k) == "about" {
			a.About = v
		}
		if strings.ToLower(k) == "submitted" {
			a.Submitted, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "submission_count" {
			a.SubmissionCount, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "updated_at" {
			t, err := time.Parse(time.RFC3339, v) // confirmed via this site: https://ijmacd.github.io/rfc3339-iso8601/#:~:text=RFC%203339%20is%20case%2Dinsensitive,the%20standard%20allows%20arbitrary%20precision.
			if err != nil {
				return err
			}
			a.UpdatedAt = t
		}
		if strings.ToLower(k) == "comment_count" {
			a.CommentCount, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
		if strings.ToLower(k) == "created_at" {
			a.CreatedAt, err = strconv.Atoi(v)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

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
		var body Body
		var bp = &body
		bp.UnmarshalJSON(b) // this is a custom unmarshal function (triggers the custom unmarshal functions for the Author and Authors types)
		if err != nil {
			return err
		}
		if bp == nil {
			return errors.New("body is nil")
		}
		// print page, per_page, total, total_pages
		println("page:", body.Page)
		println("per_page:", body.PerPage)
		println("total:", body.Total)
		println("total_pages:", body.TotalPages)
		// set the total pages
		totalPages = body.TotalPages
		// iterate through the data and print the data that meets the threshold
		for _, author := range body.Data {
			// if author.SubmissionCount > threshold {
			// 	println(author.ID, author.SubmissionCount)
			// }
			// print to stdout
			fmt.Println(author.ID, author.SubmissionCount)
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
