# Package dependencies

will print all internal package dependencies in a tree form

Usage: `package-dependencies` or `package-dependencies -m cmd/web` in case your main package is located in `cmd/web`

Also see `package-dependencies -h` for help:

```
  -m string
    	Path to files containing the main package (if not in root, it is likely you have multiple)
  -p string
    	Base path to the directory (default ".")
```

## example output
```
main
 |- lib
 |- resque
 |- clientinfo
 |- config
 |   |- lib
 |- helper
```
