package calculator

import (
  "encoding/base64"
  "os"
  "math"
  "strings"
)

const list_overhead = 1
const list_base_size = 3

type Calculator struct {
  TotalSize int
  CapacityUnit CapacityUnit
  Item map[string]interface{}
}

type CapacityUnit struct {
  Read int
  ConsistentRead float64
  TransactionRead int
  Write int
  TransactionWrite int
}

func (c *Calculator) CalculateCapacityUnit() {
  c.CapacityUnit  = CapacityUnit{
    Read: int(math.Ceil(float64(c.TotalSize) / float64(4096))),
    ConsistentRead: (math.Ceil(float64(c.TotalSize) / float64(4096))) * 0.5,
    TransactionRead: int(math.Ceil(float64(c.TotalSize) / float64(4096))) * 2,
    Write: int(math.Ceil(float64(c.TotalSize) / float64(1024))),
    TransactionWrite: int(math.Ceil(float64(c.TotalSize) / float64(1024))) * 2,
  }
}

func (c *Calculator) Calculate() {
  for k, v := range c.Item {
    attr := v.(map[string]interface{})

    c.TotalSize += len(k)
    c.calculateAttr(attr)
  }
  c.CalculateCapacityUnit()
}

func (c * Calculator) calculateAttr(attr map[string]interface{}) {
  if val, ok := attr["N"]; ok && val != nil {
    c.calculateNum(val.(string))
    return
  }

  if val, ok := attr["S"]; ok && val != nil  {
    c.calculateStr(val.(string))
    return
  }

  if val, ok := attr["B"]; ok && val != nil {
    c.calculateBin(val.(string))
    return
  }

  if val, ok := attr["L"]; ok && val != nil {
    c.calculateList(val.([]interface{}))
    return
  }

  if val, ok := attr["M"]; ok  && val != nil {
    c.calculateMap(val.(map[string]interface{}))
    return
  }

  if val, ok := attr["BOOL"]; ok  && val != nil {
    c.calculateBool()
    return
  }

  if val, ok := attr["NULL"]; ok && val != nil {
    c.calculateNull()
    return
  }

  if val, ok := attr["SS"]; ok && val != nil {
    c.calculateStrSet(val.([]interface{}))
    return
  }

  if val, ok := attr["NS"]; ok && val != nil {
    c.calculateNumSet(val.([]interface{}))
    return
  }

  if val, ok := attr["BS"]; ok && val != nil {
    c.calculateBinSet(val.([]interface{}))
    return
  }

  panic("invalid attribute type")
}

func (c *Calculator) calculateNum(v string) {
  n := v
  b := 1

  if strings.Contains(n, "-") {
    n = strings.TrimLeft(n, "-")
    b+=1
  }

  i := strings.Index(n, ".")
  if i != -1 {
    b+=1
    l := strings.TrimLeft(n[:i], "0")
    r := strings.TrimRight(n[i+1:], "0")
    n = l + r
  } else {
    n = strings.TrimLeft(n, "0")
  }

  b += int(math.Ceil(float64(len(n) / 2)))
  if b > 21 {
    b = 21
  }
  c.TotalSize += b
}

func (c *Calculator) calculateStr(v string) {
  c.TotalSize += len(v)
}

func (c *Calculator) calculateBin(v string) {
  dec, err := base64.StdEncoding.DecodeString(v)
  if err != nil {
    os.Exit(1)
  }
  c.TotalSize += len(dec)
}

func (c *Calculator) calculateList(v []interface{}) {
  c.TotalSize += list_base_size
  for _, value := range v {
    c.calculateAttr(value.(map[string]interface{}))
    c.TotalSize += list_overhead
  }
}

func (c *Calculator) calculateMap(v map[string]interface{}) {
  c.TotalSize += list_base_size
  for k, v := range v {
    attr := v.(map[string]interface{})

    c.TotalSize += len(k)
    c.calculateAttr(attr)
    c.TotalSize += list_overhead
  }
}

func (c *Calculator) calculateBool() {
  c.TotalSize += 1
}

func (c *Calculator) calculateNull() {
  c.TotalSize += 1
}

func (c *Calculator) calculateStrSet(v []interface{}) {
  for _, s := range v {
    c.calculateStr(s.(string))
  }
}

func (c *Calculator) calculateNumSet(v []interface{}) {
  for _, s := range v {
    c.calculateNum(s.(string))
  }
}

func (c *Calculator) calculateBinSet(v []interface{}) {
  for _, s := range v {
    c.calculateBin(s.(string))
  }
}
