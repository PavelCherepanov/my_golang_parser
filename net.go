package main

import (
    "net/http"
    "os"
    "io/ioutil"
  //  "io"
    "fmt"
    "regexp"

)


func main() {
    resp, err := http.Get("https://example.com")
    HandleError(err)

    defer resp.Body.Close()
    //io.Copy(os.Stdout, resp.Body)

    body, err := ioutil.ReadAll(resp.Body)
    HandleError(err)

    f, err := os.Create("html.file")
    HandleError(err)

    defer f.Close()
    err = ioutil.WriteFile("html.file", body, 0777)
    HandleError(err)

    re := regexp.MustCompile(`<h1>(.*?)</h1>`)
    fmt.Printf("%q\n", re.FindAllString(string(body), -1))
    // matched, err := regexp.MatchString(`<h1>([\s]+?)<h1>`, string(body))
    // fmt.Println(matched)
}

func SavePageInFile(){
  f, err := os.Create("html.file")
  HandleError()

  defer f.Close()
  err = ioutil.WriteFile("html.file", body, 0777)
  HandleError(err)
}

func HandleError(err error){
  if err != nil {
      panic(err.Error())
  }
}
