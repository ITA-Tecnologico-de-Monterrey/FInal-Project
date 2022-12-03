
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
    <a href="https://www.youtube.com/@taquinhosdumacaco">->Team YouTube Channel<-</a>
    <br />
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
    <li><a href="#Set-Up-Guide">How To...</a></li>
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


* Set up the Mosquitto Brokers

First, you should set up both the Mosquitto Bridge and the Mosquitto Central. To do so, using the cloud service of your choice, create two different virtual machines: the first one, the bridge, and the second one, the central. Install Mosquitto Eclipse in both machines, which can be done with the command:
```bash
sudo apt install -y mosquitto
```

For the bridge, use the [configuration file](https://github.com/ITA-Tecnologico-de-Monterrey/Final-Project/tree/main/ubuntu_vm_ec2_bridge/Mosquitto) inside ubuntu_vm_ec2_bridge/Mosquitto, and for the central, use the [configuration file](https://github.com/ITA-Tecnologico-de-Monterrey/Final-Project/tree/main/ubuntu_vm_ec2_central/Mosquitto) inside ubuntu_vm_ec2_central/Mosquitto. Inside the bridge config file, make sure to change the bridge address to the address of the machine hosting your central. You can finally start the service with the command:
```bash
mosquitto -c mosquitto.conf
```

Note: after installing Mosquitto, it automatically starts an instance on background. You might want to kill it first, because the ports might conflict, but you can also change the ports inside the conf files.`


* Set up a Publisher, a Broker and a Subscriber using Mosca JS

Using Mosca JS is pretty simple. First, make sure to have node 14.0.0, feel free to use either nvm or other tool to install it. Once you have the latest version of node and npm installed, open the folder were you have your mosca js scripts and your dependencies listen on a json file and simply run:
```bash
npm i
```
And then, just execute your code using:
```bash
node <name_code_file>
```
Note: If you get an error in the validator.js file, go to that file, edit it with nano and simply comment the line which is thrwoing the error.


* Set up the React-Native Android App

Simply follow any react-native set up link, like this [one](https://reactnative.dev/docs/environment-setup) which includes installing Node.js, the react-native lib and setting up Android Studio. After setting the environment up, simply run the following command inside the app terminal to start metro:
```bash
npx react-native start
```
Open your app on Android Studio and start runing your emulator. And in another terminal select the android file within the app and run:
```bash
npx react-native run-android
```

You'll want to configure the bridge machine address inside app.js to make sure it connects to your bridge vm.


* Set up the MYSQL Database

Once you have your virtual machines set up, go to the central one, open it in the terminal. Run the following commands to update the vm:
```bash
sudo apt update
```
```bash
sudo apt upgrade
```
Now that you have your vm ready, go ahead and install mysql:
```bash
sudo apt install mysql-server
```
Now that you have mysql, you can go ahead and enter the system by running the following command:
```bash
sudo mysql -uroot
```
Feel free to add a password to the root, or even to create a new user. To learn more about this, visit this [page](https://phoenixnap.com/kb/how-to-create-new-mysql-user-account-grant-privileges)
Now that you have the server running, feel free to create a database and a tabla for the project. If you would like to import the database of the project here in the github. Simply run:
```bash
sudo mysql -uroot -p <db_name> < events.sql
```
Just make sure to clone this [repo](https://github.com/ITA-Tecnologico-de-Monterrey/eventsDB) and that your present working directory is inside this repo.


* Set up the Grafana Dashboard

Finally, for the dashboard install grafana following these [steps](https://grafana.com/grafana/download) on ubuntu. And once you've got everything installed. Start it using:
```bash
sudo systemctl start grafana-server.service
```
Note: run the commands that appear in the installation to set up grafana to start the server automatically everytime the vm starts.

Now to make a tunnel, use nginx, install it with:
```bash
sudo apt install nginx
```
Now to make your grafana server show in a browser when you enter your ip, go to the directory:
```bash
 cd /etc/nginx/sites-available
```
Run
```bash
 sudo touch grafana
```
And write out the grafana file with this [configuration](https://github.com/ITA-Tecnologico-de-Monterrey/eventsDB/blob/main/grafana). Now that you have your grafana file, to create a simbolic link simply run: 
```bash
 sudo ln -s /etc/nginx/sites-available/grafana /etc/nginx/sites-enabled/
```
Once you have finished that, go to the sites-enabled directory with:
```bash
cd /etc/nginx/sites-enabled
```
Delete the default file:
```bash
udo rm default 
```
And restart the server:
```bash
sudo systemctl restart nginx
```
And your'e all set, now when if you enter your central vm machine's ip on a browser, it will rediret you to a grafana site where the user and password is both admin.

For more questions about the set up and compilation, feel free to watch thge youtube videos provided by our page or to contact us via e-mail

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

