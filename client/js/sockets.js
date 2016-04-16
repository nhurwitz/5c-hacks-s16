var exampleSocket = new WebSocket("ws://localhost:8000/ws");

exampleSocket.onopen = function(event) {
  console.log('here');
}

var myID = null;

response = {
    "sideLength": 10,
    "pendingPoints": [{
        "x": 7,
        "y": 5,
        "z": 3
    }, {
        "x": 2,
        "y": 2,
        "z": 2
    }],
    "snakes": {
        "abc123thisIsAnID": {
            "id": "abc123thisIsAnID",
            "head": {
                "x": 8,
                "y": 8,
                "z": 7
            },
            "tail": [{
                "x": 8,
                "y": 7,
                "z": 7
            }, {
                "x": 8,
                "y": 6,
                "z": 7
            }],
            "direction": "down"
        },
        "abc124thisIsAnID": {
            "id": "abc123thisIdsAnID",
            "head": {
                "x": 2,
                "y": 2,
                "z": 1
            },
            "tail": [{
                "x": 2,
                "y": 3,
                "z": 1
            }, {
                "x": 2,
                "y": 3,
                "z": 2
            }],
            "direction": "down"
        }
    }
}

exampleSocket.onmessage = function(event) {
  var serverEvent = JSON.parse(event.data);

  switch (serverEvent.eventType) {
    case 'Spawn':
      // nothing here
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
      return;
  }
}
