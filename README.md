TODO - try using standard apis and plain SQL

# GoLang - REST Backend API

_Provide API endpoints for fetching from a GoLang server._

## Purpose of the application

The server / app receives [`HTTP requests`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Methods) and from the [`allowed origins`](https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS) that were defined. Depending on which endpoint (e.g. "golang.server.com/user/get/1") is being called, the server will handle the incoming requests differently and return a corresponding response.

This project is plain [_GoLang_](https://go.dev/), I think that it is a really easy and maintainable language for backend implementations. MySQL was a good alternative for my case of the local development.

### Features

The main feature is the fetching of data from the database ( in my case from [MySQL](https://www.mysql.com/) ). Furthermore, I added a authentication middleware, where the middleware makes sure that the required [`auth-header`](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Authorization) is in the HTTP-header of the request.

When a 'admin' logs in, the timestamp, which indicates when to the admin logged in the last time, will updated to the current time.  

## Installation

Here, you find information about how to install the application.

### GoLang Setup

To run the app you have to have GoLang installed.

1. You need to head to the download-page for the latest version of Go -> [here](https://go.dev/doc/install) and download it. Choose your operation system and the binary package that is meant for it.
2. Follow the installation instructions of the deamon.
3. To check if GoLang is installed, enter the following command in the terminal:
    ```bash
        go version
    ```
    The expected result is: 
    > go version go1.23.0 darwin/amd64 (for mac)

**Now, GoLang is installed, we are good to go!**

### Database Setup

For MySQL, I needed to specify a connection-string and pass it to [GORM](https://gorm.io/index.html): 

```golang
// from dbhandler
    var (
	DB               *gorm.DB
	connectionString = "root:MySqLt3sT25#@tcp(127.0.0.1:3306)/website?charset=utf8&parseTime=true&loc=Local"
) 

func Connect() {
	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
```

Depending on which database you use, you will have to specify a different connection string. You can find more help [here](https://gorm.io/docs/connecting_to_the_database.html)!

**_Of course, you have to install the database yourself and prepare it for the development!_**

### Pull from GitHub

So, now all technologies are installed. The next step is to pull the repository to your local machine. 

#### Git (optional)

You can check if Git is installed by entering this command in the terminal:
```bash
    git version
```

You should get something like this, showing the version:
> git version 2.46.0

If you already have git installed, then you can skip this part!

You need to install [Git](https://git-scm.com/), a version control tool, to pull the repository from GitHub. 

#### Local Repository Setup

1. Go to the repository on my profile and copy this link:
   > https://github.com/M4tt1-Coder/backend-portfolio-API.git

2. Head to your local location where you want to place the repository. Open a terminal in your chosen folder. Enter this command:
    ```bash
        git clone https://github.com/M4tt1-Coder/backend-portfolio-API.git
    ```
    It will copy the repository to your local location!

Now, you are good to go!

#### Environment Setup

You need to prepare a `.env` - file with three variables.

Open a terminal in the route folder of the project. Enter this command to create a `.env` - file: 
```bash
    touch .env
```

Now, add your variables to the file by coping this code-block and changing the value of the variables **_API_KEY_** & **_DECRYPTION_KEY_** to yours:  

```bash
    # Copy these variables into your .env file and add your variables values
    API_KEY="Your API key to access the application through a HTTP request"
    DECRYPTION_KEY="Add a secret key to decrypt the authentication header"
```

128-bit secret keys would be a good choice for good practice. You generate some [here](https://generate-random.org/encryption-key-generator?count=1&bytes=128&cipher=aes-256-cbc&string=&password=).

### Start the app

Dependencies will be pulled automatically! You can find more information [here](https://go.dev/doc/modules/managing-dependencies).

To get started, maybe take a look at the GoLang [documentation](https://go.dev/doc/tutorial/getting-started)!

Open a terminal in the project folder!
By entering the following command, you can run the application:
```bash
    go run main.go
```

## Usage

At the very last, we are ready to use the server for incoming requests. The GoLang server serves on port 8080: `http://localhost:8080`. 

All individual routes can be looked up in the `./routes/routes.go` module.
```golang
    var AllRoutes = func(router *mux.Router) {
	    //partner routes
	    router.HandleFunc("/partner/create/{adminid}", controller.CreatePartner).Methods("POST")
	    router.HandleFunc("/partner/getAll/", controller.GetAllPartners).Methods("GET")
	    router.HandleFunc("/partner/get/{id}", controller.GetPartner).Methods("GET")

        // more routes
}
```

The different controller functions sometimes need a HTTP body or url parameter. You can find more information in the documentation of the single functions!

## Credits 

Core concepts can be found in the [this](https://www.youtube.com/watch?v=QQfyTmz3a18&list=PL5dTjWUk_cPbjazI1vRuTRZi6o5QlVAAR&index=1) YouTube video of `Akhil Sharma`. 

Besides that, the main technologies are [MySQL](https://www.mysql.com/) and [GORM](https://gorm.io/index.html) for a database and ORM.

### Social Media Links

- [**LinkedIn**](https://www.linkedin.com/in/matthis-gei%C3%9Fler-4198b9258)
- [**GitHub**](https://github.com/M4tt1-Coder)
- [**PortFolio Website**](https://matthisgeissler.pages.dev)
