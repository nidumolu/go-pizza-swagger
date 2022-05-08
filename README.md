Prerequisites:
	Install git and then clone gp-swaggerrepo :

	git clone https://github.com/go-swagger/go-swagger

Step 1
	mac
		brew tap go-swagger/go-swagger
	brewbrew install go-swagger

	(or)
		git clone https://github.com/go-swagger/go-swagger
		It can be installed using :

	go install ./cmd/swagger

Step 2
	To verify that go-swagger has been installed, go to cmd and type “swagger” and press ENTER. That should give below output.

	srinivasni@Admins-MacBook-Pro go-pizza-swagger % swagger                     
	Please specify one command of: diff, expand, flatten, generate, init, mixin, serve, validate or version
	srinivasni@Admins-MacBook-Pro go-pizza-swagger % 

	We’ll start by creating a api that lists Pizza(s) with various combos and topping available. We’ll describe an endpoint to get menu data as ComboList Object:

	Note: Make sure you validate the yaml file using https://editor.swagger.io/

Step 3

	Create the project folder
	a. Create a dir named go-pizza-swagger  .

	b. then run below command:

	go mod init go-pizza-swagger 

Step 4

	Once done, run below command:

	swagger generate server -f api/swagger.yaml --default-scheme http
	This will create a 3 directories

	Take you time and go through the files which has been created by go-swagger

Step 5
	To install all dependencies of this project, run below command under go-pizza-swagger

	For this generation to compile you need to have some packages in your GOPATH:

		* github.com/go-openapi/runtime
		* github.com/jessevdk/go-flags

	You can get these now with: go get -u -f ./...


Step 6
	Run the app locally on port 8080
	go run ./cmd/pizza-with-swag-server/main.go --scheme http --port=8080

Step 7
	Test from your loacal postman
	GET localhost:8080/v1/menu

	You will see a response like : "operation menu.ComboList has not yet been implemented"
	As you are yet to implement it