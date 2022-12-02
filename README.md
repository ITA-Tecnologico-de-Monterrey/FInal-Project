
<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/ITA-Tecnologico-de-Monterrey">
    <img src="images/logo.png" alt="Logo" width="350" height="200">
  </a>

  <h3 align="center" style="font-size:150px";>
  GSL Final Course - ITA & Tecnológico de Monterrey</h3>

  <p align="center">
    Taquinhos du Macaco Team!
    <br />
    <a href="https://github.com/ITA-Tecnologico-de-Monterrey/Final-Project"><strong>Explore the project »</strong></a>
    <br />
    <br />
  </p>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#Authors">Authors</a></li>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    <li>
      <a href="#About-the-project">About The Project</a>
    </li>
    <li><a href="#Motivation-of-the-project">Project's Motivation</a></li>
    <li><a href="#Requirements">Requirements</a></li>
  </ol>
</details>


## Authors

- [@adri-agsz](https://www.github.com/adri-agsz)
- [@Aram32mm](https://github.com/Aram32mm)
- [@DanielHBtec](https://github.com/DanielHBtec)
- [@EnzoVargasM](https://www.github.com/EnzoVargasM)
- [@EricBastos](https://www.github.com/EricBastos)
- [@FedeZenten](https://www.github.com/FedeChocuh)
- [@InakiVa](https://github.com/InakiVA)
- [@JFranciscoGR03](https://github.com/JFranciscoGR03)
- [@ViniciusDePadua](https://www.github.com/viniciusdepadua)


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

![IotSolution](/images/architecture.png)

This will be implemented through an 4-stage IoT architecture: 
* Sensors and actuators
* Internet gateways, data adquisition systems
* Edge IT
* Data Center Cloud

The drones will provide the two middle layers through an MQTT broker 
that will allow the users to forward messages from their smartphones to the
crisis management center, whis is hosted in the cloud.

### Built With

This project will be developed using the following technologies, each one of them carefully choosen to fulfill the requirements that the project demands.

* [Mininet-Wifi](https://mininet-wifi.github.io/)
* [![Android Device][android.com]][android-url]
* [![Eclipse Mosquitto][eclipse.com]][eclipse-url]
* [Grafana](https://grafana.com/oss/grafana/)
* [![python][python.org]][Python-url]
* [![MySQL][mysql.com]][mysql-url]

<p align="right">(<a href="https://github.com/ITA-Tecnologico-de-Monterrey/gsl-iot-2022/blob/main/SOLUTION_DESCRIPTION.md">Read more</a>)</p>



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

## Requirements

Check [Requirements.md](https://github.com/ITA-Tecnologico-de-Monterrey/Final-Project/blob/main/REQUIREMENTS.md) to read the user stories.

## How To...

* Mosquitto

First, you should setup both the Mosquitto Bridge and the Mosquitto Central. To do so, using the cloud service of your choice, create two different virtual machines: the first one, the bridge, and the second one, the central. Install Mosquitto Eclipse in both machines (which can be done with the command sudo apt install -y mosquitto).

For the bridge, use the configuration file inside ubuntu_vm_ec2_bridge/Mosquitto, and for the central, use the configuration file inside ubuntu_vm_ec2_central/Mosquitto. Inside the bridge config file, make sure to change the bridge address to the address of the machine hosting your central. You can finally start the service with the command mosquitto -c mosquitto.conf.

Note: after installing Mosquitto, it automatically starts an instance on background. You might want to kill it first, because the ports might conflict, but you can also change the ports inside the conf files.`

* App

Simply follow any react-native set up link, which includes installing Node.js, the react-native lib and setting up Android Studio. After setting the environment up, simply run npm run android to start the Android app. You'll want to configure the bridge machine address inside app.js to make sure it connects to your bridge vm.

* Set up a Publisher, a Broker and a Subscriber using Mosca JS

* Set up the React-Native Android App

* Set up the MYSQL Database

* Set up the Grafana Dashboard


<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[python.org]: https://img.shields.io/badge/Python-14354C?style=for-the-badge&logo=python&logoColor=white
[python-url]: https://www.python.org/
[mysql.com]: https://img.shields.io/badge/MySQL-00000F?style=for-the-badge&logo=mysql&logoColor=white
[mysql-url]: https://www.mysql.com/
[eclipse.com]: https://img.shields.io/badge/Eclipse-2C2255?style=for-the-badge&logo=eclipse&logoColor=white
[eclipse-url]: https://mosquitto.org/
[android.com]: https://img.shields.io/badge/Android-3DDC84?style=for-the-badge&logo=android&logoColor=white
[android-url]: https://developer.android.com/studio

