window.addEventListener("load", function(evt) {
    var output = document.getElementById("rootcontainer");

    var draw = function(stringSet) {

        output.innerHTML = "";

        var ledStrip = document.createElement("div");
        ledStrip.setAttribute("class", "jumbotron");
        var titlebox = document.createElement("h2")
        var title = document.createTextNode("Led Strip");
        titlebox.appendChild(title)
        ledStrip.appendChild(titlebox);


        for (var element of stringSet.set) {
            var led = document.createElement("button");
            if (element.selected) {
                led.setAttribute("class", "btn btn-warning")
            } else {
                led.setAttribute("class", "btn")
            }



            var port = document.createTextNode(element.port);
            led.appendChild(port);
            if (element.value) {

                var ledOnImg = document.createElement("img");
                ledOnImg.setAttribute("src", "/img/ledon.png");
                ledOnImg.setAttribute("alt", "ON");
                ledOnImg.setAttribute("width", "42")
                led.appendChild(ledOnImg);
            } else {

                var ledOffImg = document.createElement("img");
                ledOffImg.setAttribute("src", "/img/ledoff.png");
                ledOffImg.setAttribute("alt", "OFF");
                ledOffImg.setAttribute("width", "42")
                led.appendChild(ledOffImg);
            }
            ledStrip.appendChild(led);
        };

        output.appendChild(ledStrip);

    };

    $.get("all", function(data) {
        draw(data);
    });


    var ws;
    if (ws) {
        return false;
    }
    ws = new WebSocket("ws://" + window.location.host + "/ws");
    ws.onopen = function(evt) {

    }
    ws.onclose = function(evt) {

        ws = null;

    }
    ws.onmessage = function(evt) {


        var obj = JSON.parse(evt.data)
        draw(obj)


    }
    ws.onerror = function(evt) {

    }




    return false;



});

function selectNext() {
    $.post("next", function(data) {

    });
};

function addPort() {

    portElement = document.getElementById("port")

    port = portElement.value;
    portElement.value = ""

    $.post("add/" + port, function(data) {

    });
};

function remove() {
    $.post("remove", function(data) {

    });
};



function selectPrevious() {
    $.post("previous", function(data) {

    });
};

function moveRight() {
    $.post("move/right", function(data) {

    });
};

function moveLeft() {
    $.post("move/left", function(data) {

    });
};

function switchSelected() {
    $.post("switch/selected", function(data) {

    });
};

function switchAll() {
    $.post("switch/all", function(data) {

    });
};