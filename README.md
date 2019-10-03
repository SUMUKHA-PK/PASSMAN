# PASSMAN

A password manager is essential to maintain lots of different passwords meeting the security needs of all the different hosts.
Here is a simple solution to the same problem at zero cost.

THIS IS PASSMAN!

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
    
## Contributions :

If you have an excellent idea on how to make this better, reach out to me by opening an issue! We can work on it together :)

## LICENSE :

MIT LICENSE provided.

## Demonstration : 


```
some_guyy@localhost:~/pk/git/Password-Manager/client$ ./build.sh 

|------ |----| ------ ------ |\    /| |----| |\    |
|     |	|    | |      |      | \  / | |    | | \   |
|------	|----| |----- |----- |  \/  | |----| |  \  |
|       |    |      |      | |      | |    | |   \ |
|       |    | -----| -----| |      | |    | |    \|

Welcome to PASSMAN, your locally hosted Password Manager!



What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

1

PASSMAN Registration sequence.
Enter your email id: qwe@123.com
Hello qwe@123.com!
Please enter your master password: 
Re-enter password to confirm: 

Generating vault key....
Your vault password is: fdb2663fa1d52232105e5ef92e63c62dd537c5c7627272af4002fa5a6808c3ce

Your auth password is: 34ff80c70d0b3aa6b1d93cc2a959fa748403f74d2ad667bd002edc6fce59b989

Registration complete!

What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

3

PASSMAN Password addition sequence.
Enter your email id: qwe@123.com
Hello qwe@123.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: fdb2663fa1d52232105e5ef92e63c62dd537c5c7627272af4002fa5a6808c3ce

map[qwe@123.com:qwe@123.com]
Enter the host of the password: amazon

Enter the password of the host: 

Re-enter password to confirm: 

Password addition complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

2

PASSMAN Vault view sequence.
Enter your email id: qwe@123.com
Hello qwe@123.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: fdb2663fa1d52232105e5ef92e63c62dd537c5c7627272af4002fa5a6808c3ce

Your vault currently looks like this : 
{"amazon":"qwe"}



What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

4

PASSMAN Master Password changing sequence.
Enter your email id: qwe@123.com
Hello qwe@123.com!
Please enter your master password: 
Re-enter password to confirm: 

Enter the new master password: 

Re-enter password to confirm: 
Your vault password is: 6ffe1f7998e9dc59c319907ab51e4730717ce2f76c7920ea3837166d9bf5011d

Password change complete!

What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

3

PASSMAN Password addition sequence.
Enter your email id: qwe@123.com
Hello qwe@123.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: 6ffe1f7998e9dc59c319907ab51e4730717ce2f76c7920ea3837166d9bf5011d

map[amazon:qwe]
Enter the host of the password: walmart

Enter the password of the host: 

Re-enter password to confirm: 

Password addition complete!


What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

4

PASSMAN Master Password changing sequence.
Enter your email id: qwe@123.com
Hello qwe@123.com!
Please enter your master password: 
Re-enter password to confirm: 
Passwords dont match!

What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

2

PASSMAN Vault view sequence.
Enter your email id: qwe@123.com
Hello qwe@123.com!
Please enter your master password: 
Re-enter password to confirm: 

Your vault password is: 6ffe1f7998e9dc59c319907ab51e4730717ce2f76c7920ea3837166d9bf5011d

Your vault currently looks like this : 
{"amazon":"qwe","walmart":"qwe"}



What do you want to do?

1. Register on PASSMAN.
2. View saved passwords.
3. Add password.
4. Change master password.
5. Exit PASSMAN :(

5
Exiting PASSMAN. Bye :)

```
