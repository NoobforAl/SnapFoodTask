# A Question

Wrong: I think this bad answer Because for each node food adder need one docker container!  
if you think this relay bad you can edit docker-compose file and add new node food adder!

### Folder Structure

Each Files What's for:

A Question/  -- main folder  
├── Api   -- api folder for get all request  
│   ├── controller  -- this folder control all request for see every request how handle see files  
│   │   └── controller.go   
│   ├── main.go  -- main file for run api, in this file logger formated   
│   └── router -- you can see all path api   
│       └── router.go  
├── build.sh -- very simple build tow app   
├── db -- all config database in this folder   
│   └── db.go  
├── postman -- file for request to api  
│   └── SnapFoodPostMan  
├── docker-compose.yml  
├── Dockerfile  
├── go.mod  
├── go.sum  
├── main.go  -- only run this file for debug codes    
├── methods  -- all methods in this file, all query and core method is in this file   
│   └── methods.go  
├── models  -- all models and schema   
│   └── models.go  
├── NodeFoodAdder  
│   ├── core  -- for get food and process for each food   
│   │   └── core.go  
│   └── main.go  - main file for run one node for get foods   
├── README.md  
├── run.sh  -- for run program in docker file   
├── testdv.db -- this file not in github but if you run air this file makes   
├── tests -- all test program in this folder  
│   └── api_test.go  
└── tmp -- this file not in github but if you run air this file makes   
    ├── build-errors.log  
    └── main.exe  

# How Run Program 

first see docker-compose.yml and set your environments.  
and run this command:
> docker-compose -d up


# How Run For Debug and Develop

firs get all dependency:
> go mod tidy

and setup redis server and make .env file.  
for help environments see .env.example.  
install air for better develop:
> go install github.com/cosmtrek/air@latest

and run code:
> air main.go

for test program:  
add test in tests folder and run this command:
> go test ./tests/...
wrong : for run test delete all database file.

for better request to api use postman file in < postman > man.  

# Tree Folder Without Description

A Question/  
├── Api  
│   ├── controller  
│   │   └── controller.go  
│   ├── main.go  
│   └── router  
│       └── router.go  
├── build.sh  
├── db  
│   └── db.go  
├── postman  
│   └── SnapFoodPostMan  
├── docker-compose.yml  
├── Dockerfile  
├── go.mod  
├── go.sum  
├── main.go  
├── methods  
│   └── methods.go  
├── models  
│   └── models.go  
├── NodeFoodAdder  
│   ├── core  
│   │   └── core.go  
│   └── main.go  
├── README.md  
├── run.sh  
├── testdv.db  
├── tests  
│   └── api_test.go  
└── tmp  
    ├── build-errors.log  
    └── main.exe   
