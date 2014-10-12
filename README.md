# music-tracker

BitShares Music Note Pre-sale Tracker.    
The source data via [agsexplorer.com][].


## Usage

```sh
$ ./music-tracker -h

NAME:
   music-tracker - BitShares Music Note pre-sale tracker

USAGE:
   music-tracker [global options] command [command options] [arguments...]

VERSION:
   0.0.0.1011

COMMANDS:
   fetch    fetch or update each day data
   summary  show bitshares music pre-sale summary
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h   show help
   --version, -v  print the version
```

```
$ ./music-tracker fetch

Day   Date          BTC         Notes/BTC
1     2014-10-06    34.786598   143733.515422
2     2014-10-07    39.459149   126713.324696
3     2014-10-08    32.596869   153388.966563
4     2014-10-09    33.257446   150342.271546
5     2014-10-10    27.134661   184266.170015
6     2014-10-11    28.818317   173500.762096
7     2014-10-12    0.372000    13440860.215054

Total BTC     AVG BTC       AVG Notes
196.425040    28.060720     178185.021697

Sparkline:
██▇▇▆▆▁
```


## Cross Platform

You can download multi platform binary if you need.

https://bintray.com/bts-utils/music-tracker/music-tracker/view/files


## Notes

If you checked the stauts on [Blockchain.info][], that the data source is from blockchain.info's history data API.    
The timestamp it uses is the time transaction is made instead of the time the block is produced which includes the transactions.    
For those transactions made close to the end of the day may be included in the block that is produced the next day.    
Therefore the totals and the ratio could be slightly different from the official calculation.


### Links

* [中文帮助](README_CN.md)


## Donates, :)

* BTC: `1PouzPJj5t3kCf5j7eYu3GXSE1PWuPkHQJ`    
* PTS: `PdNWuJJS2DV2fSzuSyuzAwsJhpu8srddGb`    
* BTSX: `trytinysmart`


## License

MIT


[agsexplorer.com]: http://www1.agsexplorer.com/
[Blockchain.info]: https://blockchain.info/address/37X8DHpfiimB7PU5y35rfBcg5Vxj2R6umL