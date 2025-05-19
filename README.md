# ydb-go-genproto

`ydb-go-genproto` includes code generation from YDB protos

## Re-generation `*.pb.go` code

1. First you must get latest YDB proto files which includes into this project as git submodule:
```bash
cd api
git checkout master
git pull origin master
cd ../
```
or just use a one-liner:
```bash
git submodule update --init --remote api
```
2. Run generation `*.pb.go` code from command:
```bash
make proto
```
This command regenerates `*.pb.go` files.

