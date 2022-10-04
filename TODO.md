# TODO

- [x] Scan phase, that recursively finds all folders and files and make some kind of struct
- [x] Display the folder struct in html
- [x] Parsing .md to .html and display html
- [x] Press a path and open the .md file in browser
- [x] Use go template to display the .md -> html file
- [] Clean html templates
- [] Create a template that holds de nav bar and import that to other templates
- [] Create a template that represents all .md files
- [] Create a template that represents the page with buttons
- [] Clean code and organize to a better understanding
- [] Render images and make css work
- [] Clean code and organize to a better understanding
- [] Create the search functionality
- [] Update files in runtime and check that live in webserver
- [] Refactor code to make it simple and normative


# Idea

- Starts program -> reads `notes` folder recurs and make some kind of map of files and folders and store like
  - `client1` is a folder with the path `/notes/client1`
  - `misc` is a folder with the path `/notes/client1/misc`
  - `test.md` is a file with the path `/notes/client1/misc/test.md`
  - `test123.md` is a file with the path `/notes/client1/misc/test123./md`

- After this, display on index.html all paths and see if it works
- After this, when pressing a path that has `.md` transform that to `.html`
- After this convert folders to buttons 

