
# Taquihnos du Macaco - GSL Final Course - ITA & Tecnológico de Monterrey




## Table of Contents

* [Authors](#Authors)
* [About the project](#About-the-project)
* [Motivation](#Motivation-of-the-project)
* [User Stories](#User-Stories)

## Authors

- [@EricBastos](https://www.github.com/EricBastos)
- [@adri-agsz](https://www.github.com/adri-agsz)
- [@Aram32mm](https://github.com/Aram32mm)
- [@InakiVa](https://github.com/InakiVA)


## About the project

The project is intented to provide an efficient and secure environment to send a distress message when 
a natural disaster occurs(in this case, an earthquake). Due to the fact that during these situations 
telecommunication  links are usually down, it is important 
to ensure another means of communication. 

The proposed solution for this problem is an app that will 
allow the victims to establish communication with a crisis management
center to inform them about their position and the severity of their situation.
During these events, drones will be deployed so that the victims' information
can be intercepted. The following image was extracted from the RETO description, which
is a visual representations of this process and a proposed design for the mobile app:

![IotSolution]([https://experiencia21.tec.mx/courses/307528/files/114737145/preview](https://github.com/ITA-Tecnologico-de-Monterrey/FInal-Project/blob/main/IoT%20Architecture.png))

This will be implemented through an 4-stage IoT architecture: 
* Sensors and actuators
* Internet gateways, data adquisition systems
* Edge IT
* Data Center Cloud

The drones will provide the two middle layers through an MQTT broker 
that will allow the users to forward messages from their smartphones to the
crisis management center, whis is hosted in the cloud.



## Project's motivation

Every year, thousands of people are victims of natural disasters(tornados, earthquakes,
 floods, hurricanes, etc.). Some of these events, due to climate change, are expected to
increase in thier frequency and intensity in the following years. Being able to respond addecuately 
to these events is crucial in order to protect a citizen's health and general wellbeing.

![Natural Disasters](https://assets.weforum.org/editor/GNQPZvZnxOWwrC6_xZAcY94z0FOvSWjuh3o6i6VweJE.PNG)

It is very important that, as software engineers, we can come up with innovative solutions from
our area of expertise in order to reduce people's vulnerability during these situations. The proposed
solution, we think, could be an effective solution for the management of crisis, specially in those places in which,
for socio-economical factors, are more prone to be affected by these disasters.
## User Stories

### Technical requirements

### Non-technical requirements
