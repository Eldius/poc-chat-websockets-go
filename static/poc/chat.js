
(() => {
const socket = new WebSocket(`ws://${document.location.host}/echo`);
const name = Math.random().toString().substr(2, 8);
var intervalID;

socket.onopen = function (e) {
    console.log("[open] Connection established");
    console.log("Sending to server");
    //socket.send("My name is John");
    let counter = 0;
    intervalID = setInterval(() => {
        socket.send(JSON.stringify({ "Msg": `Test message ${counter++}`, "Name": name }));
    }, 2000);
};

socket.onmessage = function (event) {
    console.log(`[message] Data received from server: ${event.data}`);
    let msg = JSON.parse(event.data);
    let textLine = `[${msg.Name}] ${msg.Msg}`
    document.querySelector("#chat").textContent = document.querySelector("#chat").textContent + "\n" + textLine + "\n";
    document.getElementById("chat").scrollTop = document.getElementById("chat").scrollHeight 
};

socket.onclose = function (event) {
    clearInterval(intervalID);
    if (event.wasClean) {
        console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
    } else {
        // e.g. server process killed or network down
        // event.code is usually 1006 in this case
        console.log('[close] Connection died');
    }
};

socket.onerror = function (error) {
    console.log(`[error] ${error.message}`);
};

})();
