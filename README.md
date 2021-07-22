# compound-sdk-go
compound sdk

## Installing

```
go get github.com/fox-one/compound-sdk-go
```

## Using

* Initialize endpoint
  
```
compound.Endpoint = "https://compound-test-api.fox.one"
```
* Request supply action url
  
```
compound.RequestSupply(...)
```

* Request borrow action url

```
compound.RequestBorrow(...)
```

* Request Liquidate action url
  
```
compound.RequestLiquidate(...)
```

More details of the API using, please read the [example](./example) 

[LICENSE](./LICENSE)