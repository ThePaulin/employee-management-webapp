# employee-management-webapp
ThePaulin/employee-management-app
## Description:

This app allows an admin(manager) account to create new users(employees) and to assign
 to them work shifts and to manage the workstation for every worker. At the same time
it allows the employee to know when and where the will be working. The app also includes
a real-time chat component to allow employees to communicate with the manager in order
to have their schedule changed or to make any other request.

figma designs: https://www.figma.com/file/U4rXmdM21zcxu69qu7U1hs/Employee-Management-App?node-id=0%3A1
or embedded : <iframe style="border: 1px solid rgba(0, 0, 0, 0.1);" width="800" height="450" src="https://www.figma.com/embed?embed_host=share&url=https%3A%2F%2Fwww.figma.com%2Ffile%2FU4rXmdM21zcxu69qu7U1hs%2FEmployee-Management-App%3Fnode-id%3D0%253A1" allowfullscreen></iframe>

## Clean Architecture
This project follows the following clean architecture principles:
- Independent of frameworks. The architecture doesn't depend on the existence of some library.
- Testable. The business rules can be tested without the UI, DB, Web Server, or any external element.
- Independent of UI. The UI can change easily, without changing the rest of the system.
- Independent of Database. You can swap out SQL for MongoDB or something else. Your business rules are not bound to the database.
- Independent of any external agency. The business rules simply don't know anything at all about the outside world.

### 4 Domain Layers
1) Entities (Models) Layer -> an object with methods, a set of ds and functions.
2) Data (Repository) Layer
3) UseCase (Service) Layer -> implements all of the use cases of the system
4) Presentation (Delivery) Layer


### System
![alt text](https://github.com/ThePaulin/employee-management-webapp/blob/main/45477.jpg?raw=true)
![alt text](https://github.com/ThePaulin/employee-management-webapp/blob/main/31676.jpg?raw=true)
