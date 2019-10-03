# Password-Manager
A docker based, locally hostable, password manager for organisations! 
Written in Go!
Based on computerphile video on password managers.


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