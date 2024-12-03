package aoc

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// FetchInput fetches the user input from the AOC website.
func FetchInput(day int) error {
	outputFile := fmt.Sprintf("day%02d/input", day)
	_, err := os.Stat(outputFile)
	if err == nil {
		return nil
	}

	session := os.Getenv("AOC_SESSION")
	if session == "" {
		return fmt.Errorf("Missing AOC_SESSION env variable")
	}

	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)

	fmt.Printf("Fetching input for day %d\n", day)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get input for day %d: invalid status code %d",
			day, resp.StatusCode)
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err

}
