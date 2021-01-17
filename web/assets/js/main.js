var userNameInput;
var roomInput;
var messageInput;
var sendButton;

let socket = new WebSocket("ws://localhost:8080/ws/client");

socket.onopen = function (e) {
    console.log("[open] Connection established");
};

socket.onmessage = function (event) {
    console.log(`[message] Data received from server: ${event.data}`);
};

socket.onclose = function (event) {
    if (event.wasClean) {
        console.log(`[close] Connection closed cleanly, code=${event.code} reason=${event.reason}`);
    } else {
        console.log('[close] Connection died');
    }
    socket.close();
};

socket.onerror = function (error) {
    console.log(`[error] ${error.message}`);
    socket.close();
};

$(document).ready(function () {
    userNameInput = $("div.message-control input[name='userName']");
    roomInput = $("div.message-control input[name='room']");
    messageInput = $("div.message-control input[name='message']");
    sendButton = $("div.message-control input[name='send']");


    sendButton.click(function () {
        if (socket.readyState === socket.OPEN) {
            socket.send(getMessageToWebsocket());
        }
    });
})

function getMessageToWebsocket() {
    var messageObj = {
        username: userNameInput.val(),
        room: roomInput.val(),
        message: messageInput.val(),
        date: Date()
    };

    return JSON.stringify(messageObj);
}