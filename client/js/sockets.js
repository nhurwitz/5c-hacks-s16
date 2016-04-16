var exampleSocket = new WebSocket("ws://localhost:8000/ws");

exampleSocket.onopen = function(event) {
    console.log('here');
}

exampleSocket.onmessage = function(event) {
    var serverEvent = JSON.parse(event.data);

    switch (serverEvent.eventType) {
        case 'Spawn':
            return;
        case 'Die':
            // nothing here
            return;
        case 'World':
            processResponse(serverEvent.world);
            return;
        case 'Leave':
            // nothing here
            return;
        case 'Welcome':
            myID = serverEvent.snakeID;
            CELL_WIDTH = GRID_WIDTH / serverEvent.world["sideLength"];
            initGrid();
            return;
    }
}

document.addEventListener('keydown', function(e) {
    e.preventDefault();
    if (e.keyCode in keys) {
        var msg = {
            actionType: "Direction",
            snakeID: myID,
            direction: keys[e.keyCode]
        };
        exampleSocket.send(JSON.stringify(msg));

    }
}, false);
