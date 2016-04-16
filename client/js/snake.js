var keys = {
    38: 'up', // up key
    40: 'down', // down key
    39: 'right', // -> key
    37: 'left', // <- key
    87: 'in', // W key
    83: 'out' // S key
};
var GRID_WIDTH = 450;
var CAMERA_Y = 300;
var CAMERA_Z = 2000;
var ANIMATION_TIME = 200;
var GRID_COLOR = "#444036";
var BACKGROUND_COLOR = "#DFEBED";
var LIGHT_COLOR = "white";
var MY_TAIL_COLOR = '#3498DB';
var MY_HEAD_COLOR = '#FF9A00';
var OTHER_TAIL_COLOR = '#95A5A6';
var OTHER_HEAD_COLOR = '#ABB7B7';
var MARGIN_TOP = 125;
var CELL_WIDTH;

var cubeGeo, squareGeo, boundingGrid, myID;

var container = document.getElementById('container');
var width = container.clientWidth;
var height = container.clientHeight;

var scene = new THREE.Scene();
var camera = new THREE.PerspectiveCamera(25, width / height, 1, 30000);

var renderer = new THREE.WebGLRenderer({
    antialias: true
});
renderer.setClearColor(BACKGROUND_COLOR,1);
renderer.setSize(width, height);
container.appendChild(renderer.domElement);

var directionalLight = new THREE.DirectionalLight(LIGHT_COLOR);
directionalLight.position.set(3 * GRID_WIDTH, 2 * CAMERA_Y, CAMERA_Z).normalize();
scene.add(directionalLight);

camera.position.set(GRID_WIDTH / 2, CAMERA_Y, CAMERA_Z);
camera.lookAt(new THREE.Vector3());
plane = 'xy';

myID = 'abc123thisIsAnID';
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

document.addEventListener('keydown', function(e) {
    if (keys[e.keyCode] == 'in' || keys[e.keyCode] == 'out') {
        if (plane == 'xy') camYZ();
        else camXY();
    }
}, false);

animate();

processResponse(response);

function processResponse(response) {
    CELL_WIDTH = GRID_WIDTH / response["sideLength"];

    cubeGeo = new THREE.BoxGeometry(CELL_WIDTH, CELL_WIDTH, CELL_WIDTH);

    var squareShape = new THREE.Shape();
    squareShape.moveTo(0, 0);
    squareShape.lineTo(0, CELL_WIDTH);
    squareShape.lineTo(CELL_WIDTH, CELL_WIDTH);
    squareShape.lineTo(CELL_WIDTH, 0);
    squareShape.lineTo(0, 0);
    squareGeo = new THREE.ShapeGeometry(squareShape);

    boundingGrid = new THREE.Object3D();

    var gridXY = createAGrid();
    boundingGrid.add(gridXY);

    var gridYZ = createAGrid();
    gridYZ.rotation.x = Math.PI / 2;
    boundingGrid.add(gridYZ);

    var gridXZ = createAGrid();
    gridXZ.rotation.y = -Math.PI / 2;
    boundingGrid.add(gridXZ);

    boundingGrid.translateY(-MARGIN_TOP);

    scene.add(boundingGrid);

    for (var key in response["snakes"]) {
        snake = response["snakes"][key];
        if (snake["id"] == myID) {
          tailColor = MY_TAIL_COLOR;
          headColor = MY_HEAD_COLOR;
        }
        else {
          tailColor = OTHER_TAIL_COLOR;
          headColor = OTHER_HEAD_COLOR;
        }
        addHead(snake["head"], headColor);
        addTail(snake["tail"], tailColor);
    }
}

function camXY() {
    plane = 'xy';
    animation = new TWEEN.Tween(camera.position).to({
        x: GRID_WIDTH / 2,
        y: CAMERA_Y,
        z: CAMERA_Z
    }, ANIMATION_TIME).start();
    animation.onUpdate(onCamUpdate);
}

function camYZ() {
    plane = 'yz';
    animation = new TWEEN.Tween(camera.position).to({
        x: CAMERA_Z,
        y: GRID_WIDTH / 2,
        z: CAMERA_Y
    }, ANIMATION_TIME).start();
    animation.onUpdate(onCamUpdate);
}

function onCamUpdate() {
    camera.lookAt(new THREE.Vector3(0, 0, 0));
}

function addTail(tail, color) {
    var cubeMaterial = new THREE.MeshLambertMaterial({
        color: color
    });
    tail.forEach(function(position) {
        var cube = new THREE.Mesh(cubeGeo, cubeMaterial);
        cube.position.set(CELL_WIDTH / 2 + position.x * CELL_WIDTH,
            CELL_WIDTH / 2 + position.y * CELL_WIDTH,
            CELL_WIDTH / 2 + position.z * CELL_WIDTH);
        boundingGrid.add(cube);
    })
}

function addHead(position, color) {
    var cubeMaterial = new THREE.MeshLambertMaterial({
        color: color
    });
    var cube = new THREE.Mesh(cubeGeo, cubeMaterial);
    cube.position.set(CELL_WIDTH / 2 + position.x * CELL_WIDTH,
        CELL_WIDTH / 2 + position.y * CELL_WIDTH,
        CELL_WIDTH / 2 + position.z * CELL_WIDTH);
    boundingGrid.add(cube);

    var squareXY = new THREE.Mesh(squareGeo, new THREE.MeshLambertMaterial({
        color: color
    }));
    squareXY.position.set(position.x * CELL_WIDTH,
        position.y * CELL_WIDTH,
        0);
    boundingGrid.add(squareXY);

    var squareXZ = new THREE.Mesh(squareGeo, new THREE.MeshLambertMaterial({
        color: color
    }));
    squareXZ.position.set(position.x * CELL_WIDTH,
        0,
        (position.z + 1) * CELL_WIDTH);
    squareXZ.rotation.x = -Math.PI / 2;
    boundingGrid.add(squareXZ);

    var squareYZ = new THREE.Mesh(squareGeo, new THREE.MeshLambertMaterial({
        color: color
    }));
    squareYZ.position.set(0,
        position.y * CELL_WIDTH,
        (position.z + 1) * CELL_WIDTH);
    squareYZ.rotation.y = Math.PI / 2;
    boundingGrid.add(squareYZ);
}

function createAGrid() {
    var material = new THREE.LineBasicMaterial({
        color: GRID_COLOR,
        opacity: 0.2
    });

    var gridObject = new THREE.Object3D(),
        gridGeo = new THREE.Geometry();

    for (var i = 0; i <= GRID_WIDTH; i += CELL_WIDTH) {
        gridGeo.vertices.push(new THREE.Vector3(0, i, 0));
        gridGeo.vertices.push(new THREE.Vector3(GRID_WIDTH, i, 0));
        gridGeo.vertices.push(new THREE.Vector3(i, 0, 0));
        gridGeo.vertices.push(new THREE.Vector3(i, GRID_WIDTH, 0));
    }


    var line = new THREE.Line(gridGeo, material, THREE.LinePieces);
    gridObject.add(line);

    return gridObject;
}


function animate() {

    requestAnimationFrame(animate);

    render();

}

function render() {
    TWEEN.update();
    renderer.render(scene, camera);
}
render();
