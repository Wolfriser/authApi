## **Authorization API** 

### **Instructions**

Procedure for the user:
<br>

1. Clone the project from the repository
2. Go to the project and run the command: `go run .`
3. After server launch use `curl -X POST -H "" -d '{"email": "your_email", "password": "your_password"}' http://localhost:8080/register` to register yourself in system.
4. Then to login use `curl -url -X POST -H "" -d '{"email": "your_email", "password": "your_password"}' http://localhost:8080/login`
5. After that you can change your data with `curl -X PUT -H "" -b "session_token" -d '{"email": "", "password": "", "birthday": "", etc...}' http://localhost:8080/update`
6. Token is returned in response and saved in cookie
7. You can log in as an admin with `curl -url -X POST -H "" -d '{"email": "admin", "password": "0000"}' http://localhost:8080/login`
8. As an Admin you have access to `http://localhost:8080/CreateFilm`

### **Autors**

[@Wolfriser](https://github.com/Wolfriser/)
