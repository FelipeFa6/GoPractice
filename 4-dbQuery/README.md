## running

add the .env credentials, in this case the .env file is uploaded for convenience purposes

```bash
make install
make
```

- make install wil start the Dockerfile container
with the postgres image.

- `make` runs the golang code in main.go

## building

```bash
make build
```

## uninstalling

```bash
make clean
```

- this will stop any container created with the
`Makefile`

- also remove the image

- if exists, it will remove the binary generated
on build
