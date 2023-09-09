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
                    console.log(res['error']);
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