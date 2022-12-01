const mqtt = require('mqtt')

const pub = mqtt.connect('mqtt://localhost:1883')

const topicName = 'events'

let payload

pub.on("connect", function (connack) {
    console.log("client connected", connack);

    setInterval(() => {
	let id = 'android-mqtt-' + Math.floor((Math.random() * 1001));
        let la = 19.27 + (Math.random() * 0.03);
        let lo = -99.122 + (Math.random() * .02);
        let co = Math.ceil(Math.random() * 3);
	let msg = 'help me please!';
	let date = Date.now();

        payload = {
	    "client": id,
            "lat": la,
            "lon": lo,
	    "color": co,
	    "msg": msg,
	    "time": date
        }
        pub.publish(topicName, JSON.stringify(payload))

        console.log("Message published: " + payload.client + ", " + payload.lat + ", " + payload.lon + ", " + payload.color + ", " + payload.msg + ", " + payload.time)
    }, 3000);
})

pub.on("error", function (err) {
    console.log("Error: " + err)
    if (err.code == "ENOTFOUND") {
        console.log("Network error, make sure you have an active internet connection")
    }
})

pub.on("close", function () {
    console.log("Connection closed by client")
})

pub.on("reconnect", function () {
    console.log("client trying a reconnection")
})

pub.on("offline", function () {
    console.log("client is currently offline")
})
