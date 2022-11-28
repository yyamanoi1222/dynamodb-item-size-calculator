# dynamodb-item-size-calculator
CLI tool for calculating dynamodb item size and capacity unit.

Calculation result is based on this.
https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/CapacityUnitCalculations.html

# Install
### Using Go
```
$ go install github.com/yyamanoi1222/dynamodb-item-size-calculator/cmd/dynamodb-item-size-calculator@latest
```

### From release
https://github.com/yyamanoi1222/dynamodb-item-size-calculator/releases  
Download the appropriate binary

# Usage
Item size can be calculated by passing dynamodb json as stdin

```
$ aws dynamodb get-item --table-name <table name> --key <key> | jq  '.Item' | dynamodb-item-size-calculator -f djson
```

## Option

<code>-f</code> option specifies the json format to be input json or djson (default is djson)
