# Help Your Team

## Overview
This project includes both a **backend** (written in Go) and a **frontend** (written in React) that work together to create a task management application.
The main page displays a list of tasks (Ant Design's table component) and allows users to view details of a specific task by clicking on it.

## Clone the Repository
To get started with the project, clone the repository:
```
git clone https://github.com/InesLozanoLopez/help-your-team
```

## Backend
The backend is built using **Go**. It provides an API for managing tasks and their details.
The backend uses a **mock database** (created with **SQLite3**), but you can easily switch to an other database adding .env and DB_PATH variable in the directory.
1. Navigate to the backend directory:
```
cd help-your-team/backend
```

2. Open the folder with your preferred code editor

3. **IMPORTANT**: Make sure to open the **backend** and **frontend** in separate instances of your editor or terminal.

4. Run the following commands to set up the backend:
```
go mod tidy 
go run main.go
```

### Test
The tests in the backend (at the moment there are just service tests) use the same .json file to mock the database as the default **SQLite3** database configuration. This is a temporary database, and it is discarded after the tests are run.
To run all the test, run from the root directory (help-your-team/backend):
```
go test ./...
```


## Frontend
The frontend is built using **React**. It communicates with the backend API to display tasks.

1. Navigate to the frontend directory:
```
cd help-your-team/frontend
```

2. Open the folder with your preferred code editor

3. **IMPORTANT**: Make sure to open the backend and frontend in separate instances of your editor or terminal.

4. Run the following commands to set up the frontend: 
```
npm install 
npm run build
npm run start
```







