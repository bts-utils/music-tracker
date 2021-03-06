# music-tracker

Music 预售数据跟踪器。    
数据来源于 [agsexplorer.com][].

## 使用

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


## 支持多个平台, Linux， Mac OS X，Windows

你可以从下面地址下载需要的版本，记得校验。

https://bintray.com/bts-utils/music-tracker/music-tracker/view/files


## 提示

如果你在[Blockchain.info][]上查看最新状态，由于数据使用了[Blockchain.info][]历史接口，    
这些数据上的时间使用的交易时间，而不是块产生时间，这些交易结束的时候，有可能会进入第二天的产生的块中，    
所以在做统计、比例时可能跟官方有些出入。


## 链接

* [English README](README.md)


## 赞助, :)

* BTC: `1PouzPJj5t3kCf5j7eYu3GXSE1PWuPkHQJ`    
* PTS: `PdNWuJJS2DV2fSzuSyuzAwsJhpu8srddGb`    
* BTSX: `trytinysmart`


## License

MIT


[agsexplorer.com]: http://www1.agsexplorer.com/
[Blockchain.info]: https://blockchain.info/address/37X8DHpfiimB7PU5y35rfBcg5Vxj2R6umL
