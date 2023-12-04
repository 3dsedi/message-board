# message board
Welcome to the Message Board project! This application is my assignment (<a href="https://github.com/3dsedi">Sedigheh Ghazinezam</a>) to join (<a href="https://github.com/othneildrew/Best-README-Template">Chaintraced</a>  (Chaintraced company is a pioneer in supply chain traceability) 


## application features
allows users to:
 * Authentication: very simple authentication has been provided (not secure)
 * create an message
 * view messages posted by others
 * reply to messages
 * manage their own messages (deleting them). 
 

## project structure
this is a full-stack assignment and it contains two sections front-end and back-end
backend is implemented in `Go` and the frontend utilizes `React`. 

### Unit tests
although I did my best, my test cases (in both front and back) were not sufficient and unfortunately, I was out of time. for the backend are written using the `testify` framework, and `Jest` is used for the frontend unit tests.


## project setup and run:
1.by `docker-compose` help
please in the project root run `docker-compose up -d` command: this will run both backend and frontend.
please create a user and start enjoying the application

2.without `docker-compose`:
- Backend Setup:
please go to the backend dir and run the below commands
```sh
make build
make run
```
it will bring up our backend on `localhost:8080` address
- frontend Setup:
please go to frontend dir and run the below commands
```sh
npm install
npm start
```
it will bring up the frontend on `localhost:3000` address

## gratitud 
it is the first time I implemented something by the Golang. although I had many challenges to resolve, it was a pleasure time for me and extended my knowledge

thanks for giving me  such an opportunity to learn a state-of-the-art language


