package main

import (
	"log"
	"os"
	"strings"
)

// scan given a path crawls it and its subfolders
// searching for Git repositories
func scan(folder string) []string {
    repositories := recursiveScanFolder(folder)
    return repositories
}

// scanGitFolders returns a list of subfolders of `folder` ending with `.git`.
// Returns the base folder of the repo, the .git folder parent.
// Recursively searches in the subfolders by passing an existing `folders` slice.
func scanGitFolders(folders []string, folder string) []string {
    // trim the last `/`
    folder = strings.TrimSuffix(folder, "/")

    f, err := os.Open(folder)
    if err != nil {
        log.Fatal(err)
    }
    files, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        log.Fatal(err)
    }

    var path string

    for _, file := range files {
        if file.IsDir() {
            path = folder + "/" + file.Name()
            if file.Name() == ".git" {
                path = strings.TrimSuffix(path, "/.git")
                // fmt.Println(path)
                folders = append(folders, path)
                continue
            }
            if file.Name() == "vendor" || file.Name() == "node_modules" {
                continue
            }
            folders = scanGitFolders(folders, path)
        }
    }

    return folders
}

// recursiveScanFolder starts the recursive search of git repositories
// living in the `folder` subtree
func recursiveScanFolder(folder string) []string {
    return scanGitFolders(make([]string, 0), folder)
}

