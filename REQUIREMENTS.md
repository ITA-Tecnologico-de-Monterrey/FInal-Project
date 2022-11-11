### Technical requirements

As a developer, I want that the messages from the brigde brokers(drones) to be stored in a database and processed so that the rescue teams can visualize them.

**Acceptance criteria**: The messages will be converted into SQL format and stored in a SQL database. With Freeboard.io or Grafana the information from the database can be printed in a helpful dashboard that can help identify and prioritize which actions are required to keep safe the disaster victims

As a developer, I want to be able to simulate the situation to better prepare my app for a real disaster and test my features.

**Acceptance criteria**: Mininet-Wifi will be used to emulate the basic network infrastructure. Victims and drones will be simulated as nodes in the network. Python will be used for easier node development and integration with Mininet.

As a developer, I want to show a dashboard with geo-information about the user nodes as the classification (KPI), calculated based on the severity sent. The formula to calculate the KPI is shown, and the color code is presented in the dashboard. 

**Acceptance criteria**: Each node sends its state using MQTT to the bridge broker. The conditions can below, middle, or high; where high, the individual is in danger, and low, he is safe. The dashboard needs to compile this data and, using the formula, define the risk situation and colorize the node representation in the map.

### Non-technical requirements

As a victim, I want to inform to a crisis management center about my position and the severity of my situation.

**Acceptance criteria**: An MQTT broker protocol will be used to allow the victims' forward messages from their smartphones to the crisis management center.

As a victim, I want to be able to send the message from anywhere, disregarding my signal, because I don't know where an accident might happen.

**Acceptance criteria**: Drones will be flying across the disaster region, so everywhere is going to have signal eventually, and that's when the message is gonna be collected.

As a developer, I want to provide communication support to the victims of natural disasters.

**Acceptance criteria**: A four-stage IoT architecture (sensors and actuators, internet gateways and data adquisition systems, edge IT, data center cloud) will be implemented, in which drones will be deployed to intercept messages to other drones until the message arrives at the CMC.

As a developer, I want to create an app with a simple and intuitive design.

**Acceptance criteria**. An app will be developed using a Javascript library called React Native and The Paho Javascript Client library that uses WebSockets to connect to an MQTT broker. The user only needs to set the broker address, the broker port, select his severity and write a message. The app will connect to an MQTT broker.