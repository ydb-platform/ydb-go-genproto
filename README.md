# ydb-go-genproto

`ydb-go-genproto` includes code generation from YDB protos

## Re-generation `*.pb.go` code

1. First you must get latest YDB proto files which includes into this project as git submodule:
```
cd api
git checkout master
git pull origin master
cd ../
```
2. Run generation `*.pb.go` code from command:
```
make proto
```
This command re-generate `*.pb.go` files

