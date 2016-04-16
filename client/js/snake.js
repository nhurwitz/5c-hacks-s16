var keys = {
    38: 'Down', // up key
    40: 'Up', // down key
    39: 'East', // -> key
    37: 'West', // <- key
    87: 'North', // W key
    83: 'South', // S key
};
var CAMERA_KEY_CODE = 32;

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
var FOOD_COLOR = 'red'
var MARGIN_TOP = 125;
var CELL_WIDTH;

var cubeGeo, sphereGeo, squareGeo, circleGeo, boundingGrid, myID;

var container = document.getElementById('container');
var width = container.clientWidth;
var height = container.clientHeight;

var scene = new THREE.Scene();
var camera = new THREE.PerspectiveCamera(25, width / height, 1, 30000);

var renderer = new THREE.WebGLRenderer({
    antialias: true,
});
renderer.setClearColor(BACKGROUND_COLOR);
renderer.setSize(width, height);
container.appendChild(renderer.domElement);

var directionalLight = new THREE.DirectionalLight(LIGHT_COLOR);
directionalLight.position.set(3 * GRID_WIDTH, 2 * CAMERA_Y, CAMERA_Z).normalize();
scene.add(directionalLight);

camera.position.set(GRID_WIDTH / 2, CAMERA_Y, CAMERA_Z);
camera.lookAt(new THREE.Vector3());
plane = 'xy';

document.addEventListener('keydown', function(e) {
    if (e.keyCode == CAMERA_KEY_CODE) {
        if (plane == 'xy') camYZ();
        else camXY();
    }
}, false);

animate();

function processResponse(response) {
    scene.remove(boundingGrid);

    CELL_WIDTH = GRID_WIDTH / response["sideLength"];

    cubeGeo = new THREE.BoxGeometry(CELL_WIDTH, CELL_WIDTH, CELL_WIDTH);
    sphereGeo = new THREE.SphereGeometry(CELL_WIDTH / 2);

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
        addFood(response["pendingPoints"], FOOD_COLOR);
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


function addFood(food, color) {
    var sphereMaterial = new THREE.MeshLambertMaterial({
        color: color
    });
    food.forEach(function(position) {
        var sphere = new THREE.Mesh(sphereGeo, sphereMaterial);
        sphere.position.set(CELL_WIDTH / 2 + position.x * CELL_WIDTH,
            CELL_WIDTH / 2 + position.y * CELL_WIDTH,
            CELL_WIDTH / 2 + position.z * CELL_WIDTH);
        boundingGrid.add(sphere);
    })
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
        color: color,
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
