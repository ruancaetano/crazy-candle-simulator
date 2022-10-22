window.onload = function () {
    let socket = new WebSocket("ws://127.0.0.1:8080/ws");
    console.log("Attempting Connection...");

    socket.onopen = () => {
        console.log("Successfully Connected");
    };


    socket.onclose = event => {
        console.log("Socket Closed Connection: ", event);
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error);
    };

    socket.onmessage = message => {
        const data = JSON.parse(message.data)
        console.log(data);
        window.addNewCandleDataPoint(data)
    }
}
