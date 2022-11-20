from paho.mqtt import client as mqtt_client
import json
import random
import time
import sys

broker = 'ec2-52-67-152-119.sa-east-1.compute.amazonaws.com'
port = 1883
topic = "python/mqtt"
client_id = f'python-mqtt-{random.randint(0, 1000)}'


def connect_mqtt():
    def on_connect(client, userdata, flags, rc):
        if rc == 0:
            print("Connected to MQTT Broker!")
        else:
            print("Failed to connect, return code %d\n", rc)
    # Set Connecting Client ID
    client = mqtt_client.Client(client_id)
    client.on_connect = on_connect
    client.connect(broker, port)
    return client


def keep_publishing(client, argv):
    _, x, y, z, severity, msg = argv
    msg_dict = {
        "x": float(x),
        "y": float(y),
        "z": float(z),
        "severity": severity,
        "msg": msg,
    }
    while True:
        time.sleep(1)
        msg_dict['time'] = round(time.time() * 1000)
        msg_json = json.dumps(msg_dict)
        msg = msg_json
        result = client.publish(topic, msg)
        # result: [0, 1]
        status = result[0]
        if status == 0:
            print(f"Send `{msg}` to topic `{topic}`")
        else:
            print(f"Failed to send message to topic {topic}")


if __name__ == '__main__':

    # file.py X Y Z Severity Message
    argv = sys.argv
    assert(len(argv) == 6)

    client = connect_mqtt()
    client.loop_start()
    keep_publishing(client, argv)
