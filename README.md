# Remote Assignment V3

This project has both react-native and golang applications in separate branches. All you need to do is follow these instructions to run them without a hassle. Although, there is no routes setup on the frontend for OAuth with twitch, there is OAuth setup in the golang application, and you check that while hitting the route. I will also specify what routes I created and how do you have to run this project. But first .... 

## Golang API Server with MySQL

   ### Prerequisites for this project

This project uses `dotenv` configuration for accessing environment variables,
you must create  `variables.env` file and this file must have these properties

```
 TWITCH_CLIENT_ID : XXXXXXXXXXXXXXXXXXXXXXXXXX
 TWITCH_CLIENT_SECRET : XXXXXXXXXXXXXXXXXXXXXX
 MYSQL_DATA_URL : XXXXXXXXXXXXXXXXXXXXXX
 REDIRECT_URI : XXXXXXXXXXXXXXXXXXXXXX
 PRIVATE_IP : XXXXXXXXXXXXXXXXXXXXXX
 MODE : XXXXXXXXXXXXXXXXXXXXXX  // AUTH or anything else!!
 ```

after you set this up in the root directory, you can run

```

go run main.go

```
As the project does not have the auth routes for frontend, you need to specify the mode of the golang application and run it on either [Postman](https://www.getpostman.com/) or your browser.

By root directory, I mean the directory where main.go is present, which is 

```
$GOPATH/src/github.com/krishnadubagunta/golang/gawkbox-assignment/
```


   ### How to install
  
  in the terminal, run this line to get the project in your `go/src/github` 
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
  

## React Native 

This is present in the react-native branch of [this](https://github.com/krishnadubagunta/golang) repository.

This project uses [CRNA CLI](https://facebook.github.io/react-native/blog/2017/03/13/introducing-create-react-native-app.html) to create the application, and [react-native-elements](https://react-native-training.github.io/react-native-elements/docs/overview.html) for the UI. 

All you need to do here is , again, setup the environment variables, which in this case, just one. 

```
HOST = xxxxxxxxx
```

this variable is given because, if `MODE=AUTH` in the Golang API, then there is no way we can use `localhost` for android and iOS, instead we have to use something called [localtunnel](https://localtunnel.github.io/www/) to expose our api server to public. now this id that localtunnel gives, can be used in `HOST= `, and everything runs perfectly. 

if `MODE=notAuth` then we can directly use our IP address as `HOST=10.0.0.82:8080` with the port number, which is 8080. 

if android gives you any problem, which it gives, you have to reverse-map the port numbers of the device to the machine's port, like 
```
adb reverse tcp:8080 tcp:8080
```

## Screenshots of the application on React Native
Live Streams List:
![alt text](https://farm1.staticflickr.com/793/39989128795_8b29ee506c_b.jpg "Streams")

User screen after selecting a stream from the list:
![alt text](https://farm5.staticflickr.com/4784/40882549401_57928de527_b.jpg "User Page to View to Live Stream")

Searching users from the database:
![alt text](https://farm1.staticflickr.com/797/39989128115_fed0427913_b.jpg "Live Search from SQL Database")

## Note : [Very Very Important] !

Live search from SQL database cannot happen until you signin or signup using OAuth for this application. i.e., for this feature to return any users, you need to use `MODE=AUTH` atleast once in the Golang API.

# Thank you for Reading ME.
