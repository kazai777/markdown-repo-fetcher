# MarkdownRepoFetcher

## Description

MarkdownRepoFetcher is a Go program designed to fetch all markdown files from a specified GitHub repository. This tool is particularly useful for developers who want to gather documentation files from a project to create a comprehensive documentation source. By collecting all markdown files, you can create your own documentation or use it as a knowledge base for training custom AI models, such as GPTs.

## Features

- Fetch all markdown files from a specified GitHub repository.

- Automatically generate a single documentation file combining all markdown files.

- Progress bar to indicate the fetching progress.

- Command-line prompts for interactive user input.

## Requirements

- Go 1.16 or higher

- A GitHub personal access token with repository access

## Installation

1. Clone the repository:

```sh
git clone https://github.com/yourusername/markdown-repo-fetcher.git

cd markdown-repo-fetcher
```

2. Replace the placeholder for the GitHub access token in the `main.go` file with your actual GitHub token:

```go
const (
    accessToken  =  "yourtoken"  // Replace with your GitHub access token
)
```

3. Build the program:

```sh
go build -o markdown-repo-fetcher
```

## Usage

1. Run the program:

```sh
./markdown-repo-fetcher
```

2. Enter the GitHub repository owner when prompted:

```
Please enter the repo owner : <owner>
```

3. Enter the GitHub repository name when prompted:

```
Please enter the repo name : <repo>
```

4. The program will fetch all markdown files from the specified repository and create a single documentation file named `documentation.md`.
 

## Contribution

Feel  free  to  open  issues  or  submit  pull  requests  if  you  have  suggestions  or  improvements.