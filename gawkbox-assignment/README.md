# Remote Assignment V3

This project has both react-native and golang applications in separate branches. All you need to do is follow these instructions to run them without a hassle. Although, there is no routes setup on the frontend for OAuth with twitch, there is OAuth setup in the golang application, and you check that while hitting the route. I will also specify what routes I created and how do you have to run this project. But first .... 

## Golang API Server with MySQL

  ## Prerequisites to this project

This project uses `dotenv` configuration for accessing environment variables,
you must create  `variables.env` file and this file must have these properties

````
 TWITCH_CLIENT_ID : XXXXXXXXXXXXXXXXXXXXXXXXXX
 TWITCH_CLIENT_SECRET : XXXXXXXXXXXXXXXXXXXXXX
 MYSQL_DATA_URL : XXXXXXXXXXXXXXXXXXXXXX
 REDIRECT_URI : XXXXXXXXXXXXXXXXXXXXXX
 PRIVATE_IP : XXXXXXXXXXXXXXXXXXXXXX
 MODE : XXXXXXXXXXXXXXXXXXXXXX  // AUTH or anything else!!
 ```

after you set this up in the root directory, you can run
````

go run main.go

```
As the project does not have the auth routes for frontend, you need to specify the mode of the golang application and run it on either [Postman](https://www.getpostman.com/) or your browser


   ## How to install
  
  in the terminal, get the repository this project is hosted on 
  ```
  go get github.com/krishnadubagunta/golang/gawkbox-assignment
  ```
  you can go to this repository and run the main.go provided, you have met the pre-requisites mentioned above. i.e., you should create `variables.env` file in the project directory before running it.
  
  You can now run this command, and make sure you're in the root directory of this project. Which is, `/go/src/github.com/krishnadubagunta/golang/gawkbox-assignment`
  
  ### Setting up the MySQL database
  Please follow the instructions on [MySQL](https://dev.mysql.com/doc/refman/5.7/en/installing.html) to install MySQL on  your local machine, and make sure that you have twitch database in it before running the server
 
  ```
  go run main.go
  ```
  
  you should see this as output, 
  ```
2018/03/18 12:35:50 There initializes the oauth
2018/03/18 12:35:50 DB  :   &{<nil> <nil> 0 0xc42019a780 false 0 {0xc420147ef0} <nil> map[] 0xc4201b8000 0x1626600 0xc4201a6a60 false}
false
Server Booted...
  ```
  if this is the first time you're running this, you will `true` before `Server Booted...` as you don't have `Users` table in your twitch database, this program creates it for you.
  

## React Native Frontend 
