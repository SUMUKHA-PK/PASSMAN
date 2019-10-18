# PASSMAN

A password manager is essential to maintain lots of different passwords meeting the security needs of all the different hosts.
Here is a simple solution to the same problem at zero cost.

THIS IS PASSMAN! It includes a client using REDIS server on your local system and a locally hostable server.

## Functions :

* Register a new user - allows any user to register for the PASSMAN service.
* View saved passwords - all locally saved passwords of the vault are displayed. 
* Add password - enables the user to add a password into the vault locally.
* Remove password - enables the user to remove any password locally.
* Change master password - enables user to change the master password that encrypts the vault.
* Sync data with server - enables the user to sync the local passwords with the locally hoster server on port `6666`.
* View data on server - enables the user to view the state of the vault in the server.
* Remove data from server - removes all data from the server if the user opts to.

## Features :

* Locally hostable on a network for home, school or university.
* Separate client for maintaining data if network is unavailable.
* Can access your data over the network with just the username and master password.
* Uses REDIS server for persistent storage,thus is crash tolerant. Restarting the server gets back all the data safely.
* Easy to use CLI commands.
* Uses Go modules! Up to date with tech!
* Setup of the client in local system is easy, single script does the whole initial setup.
* HACKABLE code, fork and make your own modifications!
* No trust involved with the cloud party as you can host it in your own network :)
* Enables password sync between multiple devices due to centralised behaviour.
* Excellent syncing ability - Passwords updated from different machines are synced through the server when prompted.

## Documentation : 

* Well commented code, easy to read and write modifications.
* GoDoc link : 

## Usage : 

* Client setup :
    - Redis server setup :
        1. Pull the docker image by `docker pull redis`
        2. Run the docker image on your local system by `docker run -d -p 6397:6397 redis` 
       
    - Run PASSMAN : 
        1. Clone the repo by `git clone https://github.com/SUMUKHA-PK/Password-Manager`
        2. Naviagate to the client folder by `cd ~/path-to-repo/client`
        3. Provide permissions by `chmod +x build.sh`
        4. Build and the client by `./build.sh`

* Server setup :
      Similar redis server setup and run `./build.sh` in /server
      
* Using docker for server:
    - Pull the docker container:
    	`docker pull peekay46/passman_server_v1.1`		
    - Run the docker image on your machine with any port you can:
    	`docker run -d peekay46/passman_server_v1.1 -p port:6666` 
    
## Contributions :

* If you have an excellent idea on how to make this better, reach out to me by opening an issue! We can work on it together :)

* I also would love to hear "your own twist" to this. If you implemented a better/creative solution based on this, I'd love to hear from you! (sumukhapk46@gmail.com or just open an issue) 

## LICENSE :

MIT LICENSE provided.

## Screenshot(s):

![alt text](https://github.com/SUMUKHA-PK/PASSMAN/blob/master/images/reg.png)

## Demonstration : 

```
+-----+ +----+ +----- +----- |\    /| +----+ |\    |
|     |	|    | |      |      | \  / | |    | | \   |
+-----+	|----| +----+ +----+ |  \/  | |----| |  \  |
|       |    |      |      | |      | |    | |   \ |
|       |    | -----+ -----+ |      | |    | |    \|

Welcome to PASSMAN, your locally hosted Password Manager!




What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
3

PASSMAN Password addition sequence.

Enter your email id: sumukhaPK@gmail.com
Hello sumukhaPK@gmail.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: a8a1252db47e17f3f1bc21b1cf296c186796aa42720df2f8844a807d017c5398

Enter the host of the password: amazon

Enter the password of the host: 

Re-enter password to confirm: 


Password addition complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
2

PASSMAN Vault view sequence.

Enter your email id: sumukhaPK@gmail.com
Hello sumukhaPK@gmail.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: a8a1252db47e17f3f1bc21b1cf296c186796aa42720df2f8844a807d017c5398

Your vault currently looks like this : 
{"amazon":{"HostPwd":"dummyaz","TimeStamp":"2019-10-06T19:41:49.147492183+05:30"}}


Vault access complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
3 

PASSMAN Password addition sequence.

Enter your email id: sumukhaPK@gmail.com
Hello sumukhaPK@gmail.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: a8a1252db47e17f3f1bc21b1cf296c186796aa42720df2f8844a807d017c5398

Enter the host of the password: fb

Enter the password of the host: 

Re-enter password to confirm: 


Password addition complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
2

PASSMAN Vault view sequence.

Enter your email id: sumukhaPK@gmail.com
Hello sumukhaPK@gmail.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: a8a1252db47e17f3f1bc21b1cf296c186796aa42720df2f8844a807d017c5398

Your vault currently looks like this : 
{"amazon":{"HostPwd":"dummyaz","TimeStamp":"2019-10-06T19:41:49.147492183+05:30"},"fb":{"HostPwd":"dummyFB","TimeStamp":"2019-10-06T19:42:42.886956591+05:30"}}


Vault access complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
4

PASSMAN Password removal sequence.

Enter your email id: sumukhaPK@gmail.com
Hello sumukhaPK@gmail.com!
Please enter your master password: 
Re-enter password to confirm: 
Enter the host you want to remove: amazon


Password removal complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
2

PASSMAN Vault view sequence.

Enter your email id: sumukhaPK@gmail.com
Hello sumukhaPK@gmail.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: a8a1252db47e17f3f1bc21b1cf296c186796aa42720df2f8844a807d017c5398

Your vault currently looks like this : 
{"fb":{"HostPwd":"dummyFB","TimeStamp":"2019-10-06T19:42:42.886956591+05:30"}}


Vault access complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Remove Password.
5. Change master password.
6. Sync data with server.
7. View data on server.
8. Remove data from server.
9. Exit PASSMAN :(
		
9
Exiting PASSMAN. Bye :)

```
