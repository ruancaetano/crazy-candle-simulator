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
    toolTip: {
        content: "Date: {x}<br /><strong>Price:</strong><br />Open: {y[0]}, Close: {y[3]}<br />High: {y[1]}, Low: {y[2]}"
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

    if (dataPoints.length === 50) {
        dataPoints.shift()
    }

    chart.render()
}