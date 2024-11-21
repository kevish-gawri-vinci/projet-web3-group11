# What is this folder for ? 
### Conventionally, for declaring structs of each potential request.
### For example signing up will have a request containing the username and the password (crypted). We will create a struct for each field / attribute.
### type AddUserRequest struct {
    Username string `form:"name" json:"name" binding:"required"`
    Password string `form:"password" json:"password" binding:"required"`
}

### This will convert the data we receive to usable objects, making it easier to code.
#### Structs of this folder will mainly be used in Handlers, when managing the incoming requests. Similar to DTO (Data to Object)

## So have a request struct for each request or a response with a body containing multiple attributes
### Eg. /auth/user-role does not require one as response is just 