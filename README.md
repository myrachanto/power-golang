# golang-echo-backend-e-store-part1

 Am in no way  belittling php but a comming from php golang is quite impressive. An expression a like to associate golang with is like "a nuclear weapon", I mean amno way a an expert but man with the little i have learned so far golang is quite impressive.
Golang is fast, efficient, profficeint and effective. may be I am biased but hey what the heck I am biased! 

Any way lets get to what we have here

## The archtecture of this application -backend application mind you!
 
 
 ## frontend<===routes<=====contollers <===== services ====> respository ===> databases
 
I have used echo famework which can be easilly substituted by any other framework like gin -sinces it only rounds the routes and the controller

on the repository i have used gorm, which also can be easily substituted with other orms such as xorm. Remember this does not alianate the use of other databases like the use of nosql databases infact they could be embed into the repository easily

##the scope of golang echo backend e-store-part1

- It has an Authorization system embed into it- with **jwt tokens**
- It has the crud opreations product
- It has the crud opreations customer
- It has the crud opreations category of products
- It has the crud opreations major category of products
- It has the crud opreations sub category of products
- It Utilizes the gorm(orm) platform to interact with the database-relational database which as i said earlier it is easily subtitutable with other orm like xorm
- It Utilizes the echo framework to its controllers and routes- easily subsituted 
-
## the **substituted** is sort of becoming a ryme in this readme file right? well there is a reason for that!
and it is simply to avoid headached with third party libraries - well i can atest to that -- you really do not want to have such an experince especially when you are on a deadline

## how to run this application is just simple -another specks of a golang programmer
-I assume you have go configured in your computer
-Pull or clone this project from github
-as for the database creation- create the .env file on the same directory as main.go file and copy paste the code below

### .env file
    EncryptionKey: Application
    DbHost: localhost
    DbPort: 5432 
    DbType: mysql
    DbName: store
    DbUsername: root
    DbPassword: 
    PORT= :7000

    DbType1: postgress

    DbType2: Mongo
    Mongohost: mongodb://localhost:27017
    MongodbName: micro1
    Collection: users
    
 this code provides enviromental variables for the application to run- easily configuration-side benefits
  
-on this directory open cmd and run "go run ." and viola the application should be running courtesy to go modules-(they will get the required packages from the internet -assuming you are connected to the internet)
-if you have go from version 12 above run "go modules init" to introduce go mmodules
-to test the aplication use postman or any other applaction
-as for routes check the routes folder
-as for the composition of tables check the modules folder


  
