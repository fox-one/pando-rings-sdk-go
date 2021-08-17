# Pando Rings SDK

## Installing

```
go get github.com/fox-one/pando-rings-sdk-go
```

## Using

* Initialize endpoint
  
```
rings.Endpoint = "https://compound-test-api.fox.one"
```
* Request supply action url
  
```
rings.RequestSupply(...)
```

* Request borrow action url

```
rings.RequestBorrow(...)
```

* Request Liquidate action url
  
```
rings.RequestLiquidate(...)
```

More details of the API using, please read the [example](./example) 

[LICENSE](./LICENSE)