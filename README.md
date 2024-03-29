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
from file  
```
$ aws dynamodb get-item --table-name <table name> --key <key> | jq  '.Item' > item.json
$ dynamodb-item-size-calculator -f djson item.json
itemSize: 70bytes
RCU: 1
RCU (consistent): 0.5
RCU (transaction): 2
WCU: 1
WCU (transaction): 2

```

from stdin  
```
$ aws dynamodb get-item --table-name <table name> --key <key> | jq  '.Item' | dynamodb-item-size-calculator -f djson
itemSize: 70bytes
RCU: 1
RCU (consistent): 0.5
RCU (transaction): 2
WCU: 1
WCU (transaction): 2

```

## Option

<code>-f</code> option specifies the json format to be input json or djson(dynamodb json) (default is djson)
