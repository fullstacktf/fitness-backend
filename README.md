# YouLift Back End
<!-- TOC depthFrom:2 orderedList:false -->

- [ER Diagram](#er-diagram)
- [API definition](#api-definition)
    - [**How it read this documentation**](#how-it-read-this-documentation)
    - [**V1**](#v1)
        - [**User**](#user)
        - [**Assigned routine**](#assigned-routine)
        - [**Routine Specific Exercises**](#routine-specific-exercises)
        - [**Permissions**](#permissions)
        - [**Roles**](#roles)
        - [**Base exercises**](#base-exercises)
        - [**Base Routines**](#base-routines)
        - [**Routine Category**](#routine-category)
        - [**Exercise Category**](#exercise-category)
        - [**Muscle**](#muscle)
        - [**Special calls**](#special-calls)
        - [**Reset password**](#reset-password)
        - [**User stats**](#user-stats)
        - [**User weight history**](#user-weight-history)
        - [**Exercise history**](#exercise-history)

<!-- /TOC -->
## ER Diagram
![ER](docs/ER/ER.png)


## API definition
### **How it read this documentation**
Every element (or CRUD) follows the same principles:

+ Has 5 routes
+ The "get one" and "delete" routes are followed by the id of the element to get/delete

This is an example with the "users" element:

**Users**

Route: user/

Request body:

```
{
  "DNI": String,
  "Name": String,
  "Surname": String,
  "Email": String,
  "Phone": String,
  "BirthDate": String,
  "Address": String,
  "UserRole": Int
}
```

<input type="checkbox" checked> Filtered

This means that the "user" calls are as follows:

+ Create (POST): user/
+ Get one (GET): user/:id
+ Get filtered (GET): user/
+ Update (PUT): user/
+ Delete (DELETE): user/:id

Also, the create, update and get filtered need a request body as specified.

The filtering of the "Get filtered" call can be done with any of the fields of the request body, for example, we would filter users with the admin role performing a GET request to user/ with this request body:

```
{
  "UserRole": 1
}
```

The filtering is only available on calls wich have a filtering tick mark like this one: <input type="checkbox" checked> Filtered

<br><br><br>

### **V1**

Every url defined in the V1 version prepends ```/v1/```

#### **User**

Route: ```user/```

Request body:

```
{
  "DNI": String,
  "Name": String,
  "Surname": String,
  "Email": String,
  "Phone": String,
  "BirthDate": String,
  "Address": String,
  "UserRole": Int
}
```

<input type="checkbox" checked> Filtered

#### **Assigned routine**

Route: ```assignedRoutine/```

Request body:

```
{
  "UserID": Int,
  "Name": String,
  "Description": String,
  "BaseRoutineID": Int
}
```

<input type="checkbox" checked> Filtered

#### **Routine Specific Exercises**

Route: ```routineSpecificExercises/```

Request body:

```
{ 
  "BaseExercisesID": Int,
  "AssignedRoutinesID": Int,
  "Series": Int,
  "Repetitions": Int
}
```

The create method requires only the "BaseExercisesID" and "AssignedRoutinesID" fields.

<input type="checkbox" checked> Filtered

#### **Permissions**

Route: ```permission/```

Request body:
```
{ 
    "Description": String
}
```

<input type="checkbox" checked> Filtered

#### **Roles**

Route: ```role/```

Request body:

```
{ 
    "Description": String,
    "Permissions": []Permissions
}
```

<input type="checkbox" checked> Filtered

#### **Base exercises**

Route: ```baseExercise/```

Request body:

```
{
  "Name": String,
  "Description": String,
  "VideoURL": String,
  "CategoryID": Int,
  "DefaultSeries": Int,
  "DefaultRepetitions": Int,
  "Muscles": []Muscle
}
```

<input type="checkbox" checked> Filtered

#### **Base Routines**

Route: ```baseRoutine/```

Request body:

```
{
  "CategoryID": Int,
  "Name": String,
  "Description": String,
  "BaseExercises": []BaseExercise
}
```

<input type="checkbox" checked> Filtered

#### **Routine Category**

Route: ```routineCategory/```

Request body:

```
{
    "Name": String
}
```

#### **Exercise Category**

Route: ```exerciseCategory/```

Request body:

```
{
    "Name": String
}
```

#### **Muscle**

Route: ```muscle/```

Request body:

```
{
    "Name": String
}
```

#### **Special calls**

These calls do not have the basic CRUD functionality, instead these are for geting specific data or performing various specific tasks.

#### **Reset password**

Route: ```resetPassword/:userId```
route method: GET

Request body: none

#### **User stats**

This element only has two calls, one for getting the user stats and other for updating them (GET & PUT).

Route: ```userStat/:userId```

Request body: none


#### **User weight history**

**POST**
This call adds a new weight history point to the specified user.

Route: ```userWeight/:userId```

Request body:
```
  {
    "Weight": Float,
    "Date": String,
    "UserID": Int
  }
```

**GET**
This call gets a paged lsit of the user's weight history.

Route: ```userWeight/:userId```

Request body: 
```
{
	"Size": Int,
	"Offset": Int
}	
```

Return body:

```
[
  {
    "ID": Int,
    "Weight": Float,
    "Date": String,
    "UserID": Int
  }
  ...
]
```
#### **Exercise history**

**POST**
This call creates a history point of the routine specific exercise indicated in the request body

Route ```history/```

Request body:

```
{
  "Measure": float,
  "MeasureUnit": String 's' | 'kg' | 'ms',
  "MeasureType": String 'weight' | 'exerciseMetric',
  "Date": String,
  "RoutineSpecificExercisesID": Int
}
```

**GET**
This call gets a paged list of the history points of a specific exercise

Route ```history/:specExerciseId```

Request body:

```
{
	"Size": Int,
	"Offset": Int
}	
```

Return body:
```
{
  "Measure": Int,
  "MeasureUnit": String,
  "MeasureType": String,
  "Date": String,
  "RoutineSpecificExercisesID": Int
}
```