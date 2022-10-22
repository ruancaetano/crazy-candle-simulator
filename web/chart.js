let dataPoints = []
let chart = new CanvasJS.Chart("chartContainer", {
    animationEnabled: true,
    theme: "light2", // "light1", "light2", "dark1", "dark2"
    exportEnabled: true,
    title: {
        text: "Real time crazy coin price"
    },
    subtitles: [{
        text: "Second Averages"
    }],
    axisX: {
        interval: 1,
        valueFormatString: "hh:mm:ss"
    },
    axisY: {
        prefix: "C$",
        title: "Price"
    },
    data: [{
        type: "candlestick",
        yValueFormatString: "C$##0.00",
        dataPoints: dataPoints
    }]
})

function addNewCandleDataPoint(candle) {
    dataPoints.push({
        x: new Date(candle.timestamp),
        y: [candle.opening, candle.highest, candle.lowest, candle.closing]
    })

    if (dataPoints.length >= 60) {
        dataPoints.shift()
    }

    chart.render()
}

(async () => {
    result = await fetch("http://127.0.0.1:8080/candles")
        .then(response => response.json())
        .then(candles => {
            return candles.map(candle => ({
                x: new Date(candle.timestamp),
                y: [candle.opening, candle.highest, candle.lowest, candle.closing]
            }))
        }).catch(() => [])

    console.log(result.slice(0, 1))
    dataPoints.push(...result)
})()

