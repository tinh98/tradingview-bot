<p align="center">
<img src="https://img.shields.io/badge/Tradingview--bot-v1.0.0-blue" />
<img src="https://img.shields.io/badge/License-Apache--2.0-green"/>
</p>
Tradingview bot is a telegram bot to help gather chart from tradingview and send to group.

<p align="center">
<img src="https://img.shields.io/badge/MacOS-Supported-orange"/>
<img src="https://img.shields.io/badge/Ubuntu-Supported-orange"/>  
</p>

## Demo
This project is also work in telegram as bot [@tradingview_awesome_bot](https://t.me/tradingview_awesome_bot). Feel free to try it.

## Notice
- You must make sure `img` directory is existed. All the chart got from tradingview will temporarily stored in `img` for 20 seconds and then be deleted.
- Make sure you have your bot token in variable `token`

## Compile & Run
```
make
tradingview-bot-linux
```

## Usage

#### Trigger by $SYMBOL
Send stock symbol to bot and bot will return chart.

Example 
```
$AAPL
```

#### Commands
- chart1d - 查询股票图表(1天)
- chart1m - 查询股票图表(1个月)
- chart3m - 查询股票图表(3个月)
- chart1y - 查询股票图表(1年)

[ chart1d | chart1m | chart3m | chart1y ] SYMBOL

Example
```
/chart1d AAPL
```