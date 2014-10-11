
# music-tracker

Music 预售数据跟踪器

## 依赖

* 苹果操作系统
* [Textql][]
* Curl

# 使用

## `fetch`

获取每日竞投数据

```sh
# Fetch 11 Oct 2014 data
./fetch 11-10-2014
```

### `total`

统计

```sh
$ ./total

Bitshares Music Note pre-sale
  20%, each day 5 million Notes
  Start: 2014-10-06 00:00:01 +0000
  End  : 2014-12-04 23:59:59 +0000
  BTC Address: 37X8DHpfiimB7PU5y35rfBcg5Vxj2R6umL

  History:

   DAY        DATE  Notes                  BTC                  Notes/BTC
     1  2014-10-06  5000000    34.77654326000001         143775.07168031277
     2  2014-10-07  5000000    39.45914932000001         126713.32469566779
     3  2014-10-08  5000000    35.20606868000001         142020.96932340582
     4  2014-10-09  5000000   29.959576970000004         166891.54206038173
     5  2014-10-10  5000000   27.908329870000003         179157.97983220554
     6  2014-10-11  5000000            8.9753232           557083.002871696

  End                Total
  2014-12-04           176.2849913
```

### Links

* [English README](README.md)


## Donate, :)

BTC: `1PouzPJj5t3kCf5j7eYu3GXSE1PWuPkHQJ`    
PTS: `PdNWuJJS2DV2fSzuSyuzAwsJhpu8srddGb`    
BTSX: `trytinysmart`    


[textql]: https://github.com/dinedal/textql
