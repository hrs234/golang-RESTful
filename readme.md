# Simple Golang RESTful API



## guides

- ## database requirement
  
  - the database is using MySQL engines
  
  -  the database structure can be found on database folder with file name <em>dblogistik.sql</em> to be imported
  

- ## running an server 
  
  - firstly you need install an [golang](https://golang.org/dl/ "link to download the golang compiler") compiler, if you already have, just skip it

  - after that run this command first
  
    ```
    go get github.com/gorilla/mux
    go get github.com/go-sql-driver/mysql
    go get github.com/joho/godotenv
    ```

  - create an <em>.env</em> file and write this into the file
    
    ```
    DATABASE=[DATABASE_NAME]
    DATABASE_PORT=[DATABASE_PORT]
    SERVER_PORT=[SERVER_PORT]

    ```
    OR

    you can just rename an file <em>.env.default</em> to <em>.env</em> and write up our settings

  - and type ` go run main.go ` to start the server
  
  <small>Notes: for password and username database can checked in `database/database.go` you can edit in here `[MYSQL_SERVER_DATABASE_USERNAME]:[MYSQL_SERVER_PASSWORD]@tcp("+host+":"+dbPort+")/"+dbName` </small>
  

- ## used external package
  
  ```
    - github.com/gorilla/mux
    - github.com/go-sql-driver/mysql
    - github.com/joho/godotenv
  ```


- # Routes
  
  
  ## Item routes

  
  ### GET Method for fetch an data

    - ` api/Item ` ->listing our notes data
      
      - __query list__
        - id -> for seeing specific Item table record (must in integer value)
        - search -> search data
  
  ### PUT Method for updating data
    
    - ` api/Item/[ID]`

  ### POST Method for inserting data

    - ` api/Item `

  ### DELETE Method for deleting data

    - ` api/Item/[ID] `

-------------------------------------------------------------

  ## Kategori routes

- ### GET Method for fetch data

  - __query list__
  
        - id -> for seeing specific Kategori (must in integer value)
        - search -> search data
  

- ### PUT Method for updating data
  
  - ` api/Kategori/[ID] ` 

- ### POST Method for insert data
  
   - `api/Kategori` 

- ### DELETE Method for deleting data
  
  - `api/Kategori/[ID]`