# music-tracker

BitShares Music Note Pre-sale Tracker.    
The data source via [agsexplorer][]


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
   fetch  fetch or update each day data
   show   show one day status
   list   list all days status
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
6     2014-10-11    17.542415   285023.466683

Total         AVG
18.477714     3079618965.000000

Sparkline:
▇█▆▆▄▁
```

[agsexplorer]: http://www1.agsexplorer.com/