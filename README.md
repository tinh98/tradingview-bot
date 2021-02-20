<p align="center">
<img src="https://img.shields.io/badge/Tradingview--bot-v1.0.0-blue" />
<img src="https://img.shields.io/badge/License-Apache--2.0-green"/>
</p>
Tradingview bot is a telegram bot to help gather chart from tradingview and send to group.

<p align="center">
<img src="https://img.shields.io/badge/MacOS-Supported-orange"/>
<img src="https://img.shields.io/badge/Ubuntu-Supported-orange"/>  
</p>

## Table of Contents

- [Demo](#demo)
- [Usage](#usage)
  - [Trigger by $SYMBOL](#trigger-by-symbol)
  - [Trigger by bot commands](#trigger-by-bot-commands)
- [Require](#require)
- [Build](#build)
- [Notice](#notice)

## Demo
This project is also work in telegram as bot [@tradingview_awesome_bot](https://t.me/tradingview_awesome_bot). Please feel free to try it.

## Usage

### Trigger by $SYMBOL
Send stock symbol to bot and bot will return chart.

Example for query Apple
```
$AAPL
```

### Trigger by bot commands
- chart1d - 查询股票图表(1天)
- chart1m - 查询股票图表(1个月)
- chart3m - 查询股票图表(3个月)
- chart1y - 查询股票图表(1年)

[ chart1d | chart1m | chart3m | chart1y ] SYMBOL

Example for query Apple
```
/chart1d AAPL
```

## Require
- [Golang](https://golang.org/dl/)
- [capture-website-cli](https://github.com/sindresorhus/capture-website-cli)
- [npm](https://github.com/nodesource/distributions/blob/master/README.md)

## Build
```
make
./tradingview-bot-linux
```

## Notice
- You must make sure `img` directory is existed. All the chart got from tradingview will temporarily stored in `img` for 20 seconds and then be deleted.
- Make sure you have your bot token in variable `token`