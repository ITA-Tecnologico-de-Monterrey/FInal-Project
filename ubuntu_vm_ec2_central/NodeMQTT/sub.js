const mqtt = require('mqtt')
const mysql = require('mysql')

const db = mysql.createConnection({
    host: "localhost",
    port: 3306,
    user: 'root',
    password: 'Iot123456789',
    database: 'events'
})

db.connect((err) => {
        if (err) {
            console.error('CONNECT FAILED', err.code);
        }
        else {
            console.error('CONNECTED'); 
	}
})

const sub = mqtt.connect('mqtt://localhost:1884')

sub.on('connect', () => {
    console.log('Connected to mqtt broker')
    sub.subscribe('events', function(err, granted){ 
    if(err){
	console.log(err)
	return
    }
    console.log('subscribed!')
    })
})

sub.on('message', (topic, message) => {
    console.log("message read")
    let m = JSON.parse(message)
    console.log(m)
    db.query(
        'insert into events set ?',
        {client: m.client, lat: m.lat, lon: m.lon, color: m.color, msg: m.msg, time: m.time},
        (err, rows) => {
            if(!err) console.log('data saved!')
            if(err) console.log(err)
        }
    )
})

