# gostatmsa

Concurrency for large msa manipulation

## Version

Initial draft to handle large protein multiple sequence alignment

## Dependencies

The tiny [go-msa](https://github.com/glaunay/go-msa) package 

## Basic usage

```bash
go run github.com/statmsa LARGE_FILE.ALN --name="cathepsin" --word="IPFVEAYIVS"
```

Where,

* LARGE_ALN is a clustal formated file
* **--name** is an optional substring to look for in sequence headers
* **--word** is an optional substring to look for in sequence content, stripped of gap symbols
