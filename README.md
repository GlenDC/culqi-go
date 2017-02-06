# culqi-go v1.2

culqi-go es una biblioteca que permite integrar la pasarela de pagos CULQI

```
$ go get -u github.com/glendc/culqi-go/api/culqi
```

## OpenAPI

culqi-go se genera usando [go-swagger][].

La especificación oficial de culqi se puede encontrar en
https://culqi.api-docs.io/v1.2

### Cómo actualizar las especificaciones

1. Guarde las especificaciones actualizadas como `./swagger/swagger.json`;
2. Genere el código: `go generate`;

[go-swagger]: https://github.com/go-swagger/go-swagger