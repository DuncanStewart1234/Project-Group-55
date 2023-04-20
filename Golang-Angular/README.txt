

In order to run our project:

1. Download the Golang-Angular directory and open two terminals in that directory.

2. Run 'go run src/server/main.go' in the first terminal. This is to launch the backend portion of our project.
   If you get an error message such as 'cannot find package "github.com/..." in any of: ...', then run
   'go get "github.com/..."' for that package.

3. In the second terminal, run 'ng serve' to launch the frontend portion of our project.

4. Now, with both the backend and frontend running, our project can be found if you go to http://localhost:4200

