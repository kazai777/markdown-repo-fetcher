package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/cheggaaa/pb/v3"
)

const (
	githubAPI   = "https://api.github.com/repos/"
	outputFile  = "documentation.md"
	accessToken = "yourtoken" // Replace with your github access token
)

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"`
}

func main() {
	var owner, repo string

	fmt.Print("Please enter the repo owner : ")
	fmt.Scanln(&owner)

	fmt.Print("Please enter the repo name : ")
	fmt.Scanln(&repo)

	fmt.Println("Fetching file list...")
	files := getFiles("", []File{}, owner, repo)
	if len(files) == 0 {
		fmt.Println("No files found in the repository or error fetching files.")
		return
	}

	mdFiles := filterMdFiles(files)
	if len(mdFiles) == 0 {
		fmt.Println("No markdown files found in the repository.")
		return
	}

	var content strings.Builder

	count := len(mdFiles)
	fmt.Printf("Found %d markdown files.\n", count)
	bar := pb.StartNew(count)

	for _, file := range mdFiles {
		data := getFileContent(file.Path, owner, repo)
		if data != nil {
			content.WriteString(fmt.Sprintf("# %s\n\n", file.Path))
			content.Write(data)
			content.WriteString("\n\n")
		}

		bar.Increment()
	}

	bar.Finish()

	err := os.WriteFile(outputFile, []byte(content.String()), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Documentation created successfully:", outputFile)
	}
}

func getFiles(path string, allFiles []File, owner, repo string) []File {
	url := fmt.Sprintf("%s%s/%s/contents/%s", githubAPI, owner, repo, path)
	fmt.Printf("Fetching files from %s\n", url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching files:", err)
		return allFiles
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error: Received status code %d\n", resp.StatusCode)
		return allFiles
	}

	var files []File
	err = json.NewDecoder(resp.Body).Decode(&files)
	if err != nil {
		fmt.Println("Error decoding response:", err)
		return allFiles
	}

	for _, file := range files {
		if file.Type == "dir" {
			allFiles = getFiles(file.Path, allFiles, owner, repo)
		} else {
			allFiles = append(allFiles, file)
		}
	}
	return allFiles
}

func filterMdFiles(files []File) []File {
	var mdFiles []File
	for _, file := range files {
		if strings.HasSuffix(file.Name, ".md") {
			mdFiles = append(mdFiles, file)
		}
	}
	return mdFiles
}

func getFileContent(path string, owner, repo string) []byte {
	url := fmt.Sprintf("%s%s/%s/contents/%s", githubAPI, owner, repo, path)
	fmt.Printf("Fetching content from %s\n", url)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "token "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error fetching file content:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("Error: Received status code %d for file %s\n", resp.StatusCode, path)
		return nil
	}

	var fileContent struct {
		Content  string `json:"content"`
		Encoding string `json:"encoding"`
	}

	err = json.NewDecoder(resp.Body).Decode(&fileContent)
	if err != nil {
		fmt.Println("Error decoding file content:", err)
		return nil
	}

	var decodedContent []byte
	if fileContent.Encoding == "base64" {
		decodedContent, err = base64.StdEncoding.DecodeString(fileContent.Content)
		if err != nil {
			fmt.Println("Error decoding base64 content:", err)
			return nil
		}
	} else {
		decodedContent = []byte(fileContent.Content)
	}

	return decodedContent
}
