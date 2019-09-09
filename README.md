# wiki-dump-parser

Parses [wiki-dump xml files](https://dumps.wikimedia.org/enwiki/) and indexes nodes as a directed graph to [big-data graph DB](https://github.com/dgoldstein1/graphApi).

[![Maintainability]()]()
[![Test Coverage]()]()
[![CircleCI]()]()

## Build it

#### Binary

```sh
go get -u github.com/dgoldstein1/wiki-dump-parser
```

#### Docker
```sh
docker build . -t dgoldstein1/wikiDumpParser
```

## Run it

```sh
dc up -d
```

or with dependencies running locally

```sh
export GRAPH_DB_ENDPOINT="http://localhost:5000" # endpoint of graph database
export TWO_WAY_KV_ENDPOINT="http://localhost:5001" # endpoint of k:v <-> v:k lookup metadata db
export PARALLELISM=20 # number of parallel threads to run
export METRICS_PORT=8002 # port where prom metrics are served
wiki-dump-parser parse enwiki-20190620-pages-articles1.xml-p10p30302 
```


## Development

#### Local Development

- Install [inotifywait](https://linux.die.net/man/1/inotifywait)
```sh
./watch_dev_changes.sh
```

#### Testing

```sh
go test $(go list ./... | grep -v /vendor/)
```

#### Benchmarks

| Dump Size | Execution Time | Number of Nodes | Number of Edges | Nodes Added / Sec |
|-----------|----------------|-----------------|-----------------|-------------------|
| 619mb     | 4m52.936s      | 1280817         | 2648926         | 4386.35           |


## Authors

* **David Goldstein** - [DavidCharlesGoldstein.com](http://www.davidcharlesgoldstein.com/?github-wiki-dump-parser) - [Decipher Technology Studios](http://deciphernow.com/)

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
