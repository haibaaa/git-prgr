package main

import ( 
  //formatted i/o operations
  "fmt"

  //adding flags to program  
  "flag"

  //string processing
  "strings"

  //traversing the folders
  "os"
  "path/filePath"
  "io/ioutil"
  "log"
)

/*
scan crawls the given path and it's subfolders
for git repositories
*/
func scan(path string) {
  fmt.Printf("found: \n\n")
  repositories := recursiveScanFolder(folder)
  filePath := getDotFilePath()
  addNewSliceElementsToFile(filePath, repositories)
  fmt.Printf("\n\nadded successfully\n\n")
}

/* generates a graph for the git contributions */
func stats(email string) {
  print("stats")
}

/*
returns a list of subfolders of the given folder ending with ".git"
recursively searches each folder of the repo for more .git folders
*/
func scanGitFolders(folders []string, folder string) []string {
  /* trim the last / */
  folder = strings.TrimSuffix(folder, "/")
 
  /* creates a folder descriptor */
  f,err := os.Open(folder)
  if err != nil {
    log.Fatal(err)
  }

  /* f.Readdir(-1) reads all files in the folder and logs an
eror on failure */
  files, err := f.Readdir(-1)
  if err != nil {
    log.Fatal(err)
  }

  var path string 

  for _, file := range files {
    if file.IsDir() {
      path = folder + "/" file.Name()
      if file.Name() == "./git" {
        path = strings.TrimSuffix(path, "/.git")
        fmt.Println(path)
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

/*
  calls 
*/
func recursiveScanFolder(folder string) []string {
  return scanGitFolders(make([]string, 0), folder)
}

func getDotFilePath() string {
  usr, err := user.Current()
  if err != nil {
    log.Fatal(err)
  }

  dotFile := usr.Homedir + "./gogitlocalstats"

  return dotFile
}

func main() {
  var folder string
  var email string
  flag.StringVar(&folder, "add", "", "add a folder to scan for repositories")
  flag.Parse()

  if folder != "" {
    scan(folder)
    return
  }

  stats(email)
}
