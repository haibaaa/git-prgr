package main

import ( 
  //formatted i/o operations
  "fmt"

  //adding flags to program  
  "flag"

  //string processing
  "strings"

  //r/w from buffer
  "bufio"

  //traversing the folders
  "os"
  "path/filePath"
  "io"
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
      folders = scanGitFolders(folderis, path)
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

/*
* parse existing to a slice
* add new items to the slice(without duplicates)
* rewrite the file
*/
func addNewSliceElementsToFile(filePath string, newRepos []string) {
  existingRepos := parseFileLinesToSlice(filePath)
  repos := joinSlices(newRepos, existingRepos)
  dumpStringSliceToFile(repos, filePath)
}

/*
* takes file content and parses line by line to a slice of strings
*/
func parseFileLinesToSlice(filePath string) []string {
  f := openFile(filePath)
  defer f.Close()

  var lines []string
  scanner := buffio.NewScanner(f)
  for scanner.Scan(){
    lines := append(lines, scanner.Text())
  }
  if err := scanner.Err(); err != nil {
    if err != io.EOF {
      panic(err)
    }
  }

  return lines
}

/*
*pretty self explanatory
*/
func openFile(filePath string)  {
  f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0755)
  if err != nil {
    if os.IsNotExist(err) {
      //file does not exist
      _, err = os.Create(filePath)
      if err != nil {
       panic(err)
      }
    } else {
      panic(err)
    }
  }
  return f  
}


/*
*adds new unique slice 
*/
func joinSlices(new []string, existing []string) []string {
  for _, i := range new {
    if !sliceContains(existing, i) {
      existing = append(existing, i)
    }
  }
  return existing
}

/*
*returns true if it finds a duplicate
*/
func sliceContains(slice []string, value string) bool {
  for _,v := range slice {
    if v == value {
      return true
    }
  }
  return false
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
