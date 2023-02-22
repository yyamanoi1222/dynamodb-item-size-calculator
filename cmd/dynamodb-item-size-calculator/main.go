package main

import (
  "fmt"
  "bufio"
  "os"
  "encoding/json"
  "flag"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "github.com/yyamanoi1222/dynamodb-item-size-calculator/pkg/calculator"
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

  c := &calculator.Calculator{Item: data}
  c.Calculate()

  fmt.Printf("itemSize: %vbytes\n", c.TotalSize)
  fmt.Printf("RCU: %v\n", c.CapacityUnit.Read)
  fmt.Printf("RCU (consistent): %v\n", c.CapacityUnit.ConsistentRead)
  fmt.Printf("RCU (transaction): %v\n", c.CapacityUnit.TransactionRead)
  fmt.Printf("WCU: %v\n", c.CapacityUnit.Write)
  fmt.Printf("WCU (transaction): %v\n", c.CapacityUnit.TransactionWrite)
}

func parseInput(format string) map[string]interface{} {
  var jsonStr string
  var data map[string]interface{}
  sc := bufio.NewScanner(os.Stdin)

  for sc.Scan() {
    jsonStr += sc.Text()
  }

  if err := sc.Err(); err != nil {
    fmt.Fprintf(os.Stderr, "err: %s \n", err)
    os.Exit(1)
  }

  if err := json.Unmarshal([]byte(jsonStr), &data); err != nil {
      fmt.Fprintf(os.Stderr, "invalid input format \n")
      os.Exit(1)
  }

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
