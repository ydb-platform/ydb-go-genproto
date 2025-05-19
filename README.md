# ydb-go-genproto

`ydb-go-genproto` includes code generation from YDB protos

## Re-generation `*.pb.go` code

1. First you must get latest YDB proto files which includes into this project as git submodule:

If you have just cloned `ydb-go-genproto` and `api` folder is empty, just use this one-liner:

```bash
git submodule update --init --remote api
```

Otherwise checkout the latest master:

```bash
cd api
git checkout master
git pull origin master
cd ../
```

2. Run generation `*.pb.go` code from command:

```bash
make proto
```

This command regenerates `*.pb.go` files.

