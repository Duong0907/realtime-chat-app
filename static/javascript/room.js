//////////////////// HANDLER ERROR //////////////////////////////

const notifications = document.querySelector(".notifications")

// Object containing details for different types of toasts
var toastDetails = {
    timer: 5000,
    success: {
        icon: 'fa-circle-check',
        text: 'Success: This is a success toast.',
    },
    error: {
        icon: 'fa-circle-xmark',
        text: 'Error: This is an error toast.',
    },
    warning: {
        icon: 'fa-triangle-exclamation',
        text: 'Warning: This is a warning toast.',
    },
    info: {
        icon: 'fa-circle-info',
        text: 'Info: This is an information toast.',
    }
}

const removeToast = (toast) => {
    toast.classList.add("hide");
    if (toast.timeoutId) clearTimeout(toast.timeoutId); // Clearing the timeout for the toast
    setTimeout(() => toast.remove(), 500); // Removing the toast after 500ms
}

const createToast = (id) => {
    // Getting the icon and text for the toast based on the id passed
    const { icon, text } = toastDetails[id];
    const toast = document.createElement("li"); // Creating a new 'li' element for the toast
    toast.className = `toast ${id}`; // Setting the classes for the toast
    // Setting the inner HTML for the toast
    toast.innerHTML = `<div class="column">
                         <i class="fa-solid ${icon}"></i>
                         <span>${text}</span>
                      </div>
                      <i class="fa-solid fa-xmark" onclick="removeToast(this.parentElement)"></i>`;
    notifications.appendChild(toast); // Append the toast to the notification ul
    // Setting a timeout to remove the toast after the specified duration
    toast.timeoutId = setTimeout(() => removeToast(toast), toastDetails.timer);
}

///////////////////// CREATE A ROOM //////////////////////////////

var roomInput = document.getElementById('room-input');
const minID = 1000;
const maxID = 9999;

function createRoom() {
    var id = Math.floor(Math.random() * (maxID - minID) + minID);
    var options = {
        method: 'POST',
        body: JSON.stringify({
            "id": id.toString(),
            "name": roomInput.value
        })
    }

    return fetch('/api/ws/createRoom', options)
        .then(res => res.json())
}



function renderRoom(id, name) {
    var html = `<span class="room-id" style="display:none">${id}</span>
                <span class = "room-name">${name}</span>
                <i class="fa fa-angle-right"></i>`
    var newRoomNode = document.createElement('div')
    newRoomNode.innerHTML = html;
    newRoomNode.classList.add('items-body-content');
    newRoomNode.onclick = joinRoomHandler;
    document.getElementsByClassName('items-body')[1].appendChild(newRoomNode)
}

roomInput.onkeydown = function (event) {
    // When pressing button 'Enter'
    if (event.keyCode == 13) {
        
        createRoom()
            .then(res => {
                roomInput.value = '';
                roomInput.focus();
                
                if (!res['error']) {
                    renderRoom(res['id'], res['name']);
                } else {
                    toastDetails.error.text = res['error'];
                    createToast('error');
                }
            })

    }
}

//////////////////////////// JOIN A ROOM //////////////////////////////

function joinRoomHandler(event) {
    var item = event.target;
    var roomId;
    var roomName;

    if (item.classList.contains('items-body-content')) {
        roomId = item.querySelector('.room-id').innerHTML;
        roomName = item.querySelector('.room-name').innerHTML;
    } else {
        roomName = item.innerHTML;
        roomId = item.previousSibling.previousSibling.innerHTML;
    }

    window.location.href = `/chat?room_name=${roomName}&room_id=${roomId}`;
}

var roomItems = document.getElementsByClassName('items-body-content');

for (let i = 0; i < roomItems.length; i++) {
    console.log(roomItems[i]);
    roomItems[i].onclick = joinRoomHandler;
}

//////////////////// LOG OUT ///////////////////////////
var logoutBtn = document.getElementById('logout-btn');
logoutBtn.onclick = function() {
    fetch("/api/user/logout");
    window.location.href = "/login";
}