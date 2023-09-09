// The chat scrolls to bottom
const scrollingElement = document.getElementsByClassName("body-panel")[0];
scrollingElement.scrollTop = scrollingElement.scrollHeight;


var chatList = document.getElementById('chat-list');
var roomID = document.getElementById('room-id').innerHTML;
var roomName = document.getElementById('room-name').innerHTML;
var currUserID = document.getElementById('user-id').innerHTML;
var currUsername = document.getElementById('username').innerHTML;

function newChatElement(username, message) {
    var chatElement = document.createElement('li')
    chatElement.classList.add('clearfix');

    if (username === currUsername) {
        let rightChatHTML = `
            <span class="chat-img pull-right">
                <img src="http://placehold.it/50/FA6F57/fff&text=ME" alt="User Avatar"
                    class="img-circle" />
            </span>
            <div class="chat-body clearfix">
                <div class="header">
                    <small class=" text-muted"><span class="glyphicon glyphicon-time"></span>... mins
                        ago</small>
                    <strong class="pull-right primary-font">${username}</strong>
                </div>
                <p>
                    ${message}
                </p>
            </div>
        `;

        chatElement.innerHTML = rightChatHTML;
        chatElement.classList.add('right');
    } else {
        let leftChatHTML = `
            <span class="chat-img pull-left">
                <img src="http://placehold.it/50/55C1E7/fff&text=U" alt="User Avatar"
                    class="img-circle" />
            </span>
            <div class="chat-body clearfix">
                <div class="header">
                    <strong class="primary-font">${username}</strong> <small
                        class="pull-right text-muted">
                        <span class="glyphicon glyphicon-time"></span>... mins ago</small>
                </div>
                <p>
                    ${message}
                </p>
            </div>
        `;

        chatElement.innerHTML = leftChatHTML;
        chatElement.classList.add('left');
    }
    chatList.appendChild(chatElement);

    // Alway scroll down
    const scrollingElement = document.getElementsByClassName("body-panel")[0];
    scrollingElement.scrollTop = scrollingElement.scrollHeight;
}


var sendBtn = document.getElementById('btn-chat');
var messageInput = document.getElementById('message-input');

function sendMessage() {
    var message = messageInput.value;
    if (message == '') {
        return
    }

    socket.send(message);
    messageInput.value = '';
    messageInput.focus();
}

sendBtn.onclick = function () {
    sendMessage();
}

messageInput.onkeydown = function (event) {
    if (event.keyCode == 13) {
        sendMessage();
    }
}

///////////////// WEB SOCKET //////////////////

console.log(roomID, currUserID, currUsername);
socket = new WebSocket(`ws://localhost:8080/api/ws/joinRoom/${roomID}?userId=${currUserID}&username=${currUsername}`);
console.log("Attempting Connection...");

socket.onopen = () => {
    console.log("Successfully Connected");
};

socket.onclose = event => {
    socket.send("Client Closed!")
};

socket.onerror = error => {
    socket.close();
    console.log("Socket Error: ", error);
};

socket.onmessage = (event) => {
    var res = JSON.parse(event.data);
    newChatElement(res['username'], res['content']);
}

window.onunload = function () {
    socket.close();
};