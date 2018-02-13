window.addEventListener("load", function(evt) {
    var output = document.getElementById("rootcontainer");

    var draw = function(stringSet) {

        window.console.log(JSON.stringify(stringSet))
        output.innerHTML = "";

        var ledStrip = document.createElement("div");
        ledStrip.setAttribute("class", "");
        var title = document.createTextNode("Led Strip");
        ledStrip.appendChild(title);




        for (var element of stringSet.set) {
            var led = document.createElement("div");
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
        window.console.log(JSON.stringify(data));
        draw(data);
    });
});


// if (ws) {
//     return false;
// }
// ws = new WebSocket("ws://" + window.location.host + "/ws");
// ws.onopen = function(evt) {
//     print("connection established");
// }
// ws.onclose = function(evt) {
//     print("connection closed");
//     ws = null;

// }
// ws.onmessage = function(evt) {


//     var obj = JSON.parse(evt.data)
//     draw(obj)


// }
// ws.onerror = function(evt) {
//     print("ERROR: " + evt);
// }
// return false;