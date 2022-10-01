package main

import (
  "fmt"
  "bufio"
  "os"
  "encoding/json"
  "flag"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {
  validFormat := []string{"djson","json"}
  var format string
  flag.StringVar(&format, "f", "djson", "input json format")
  flag.Parse()

  if !contains(validFormat, format) {
      fmt.Fprintf(os.Stderr, "invalid format erro \n")
      os.Exit(1)
  }

  data := parseInput(format)

  c := &Calculator{item: data}
  c.Calculate()

  fmt.Printf("itemSize: %vbytes", c.totalSize)
}

func parseInput(format string) map[string]interface{} {
  var jsonStr string
  var data map[string]interface{}
  sc := bufio.NewScanner(os.Stdin)

  for sc.Scan() {
    jsonStr += sc.Text()
  }

  json.Unmarshal([]byte(jsonStr), &data)

  if format != "djson" {
    r, err := dynamodbattribute.MarshalMap(data)
    if err != nil {
      fmt.Fprintf(os.Stderr, "marshal error occured \n")
      os.Exit(1)
    }
    s, _ := json.Marshal(r)
    json.Unmarshal([]byte(s), &data)
  }

  return data
}

func contains(s []string, e string) bool {
  for _, a := range s {
    if a == e {
      return true
    }
  }
  return false
}
