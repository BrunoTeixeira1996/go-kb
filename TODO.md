# TODO

- [x] Scan phase, that recursively finds all folders and files and make some kind of struct
- [x] Display the folder struct in html
- [x] Parsing .md to .html and display html
- [x] Press a path and open the .md file in browser
- [x] Use go template to display the .md -> html file
- [x] Clean html templates
- [x] Create a template that holds de nav bar and import that to other templates
- [x] Create a template that represents all .md files (press a .md file and navigate to that template with the custom .md but with the nav bar)
- [x] Clean code and organize to a better understanding
- [x] Strip full path to only show the folder or the file name
- [x] Use a different color if folder or file
- [x] Change theme to be more like https://github.com/tsoding/tsoding.github.io 
- [x] Create the logic for the rendering
- [x] Create the back button
- [x] Render images from markdown
- [x] Pass notes location from command line
  - [x] Working, but now it needs some clean and separation
- [x] Clean code and organize to a better understanding
- [x] Clean the hardcoded paths for `template.Must`
- [ ] Implement app in a docker container
	- Implemented, now it needs some tweaks
		- [ ] Make it possible to rebuild image
		- [ ] Create a small image
		- [x] Change CMD in Dockerfile to run only "/app/server" so I can pass the name of the folder I want when building the docker image
		- [ ] Document how to build from Docker
- [ ] Create the back button on .md files too
- [ ] Create the search functionality
- [ ] Update files in runtime and check that live in webserver
- [ ] Refactor code to make it simple and normative


# Idea

- Starts program -> reads `notes` folder recurs and make some kind of map of files and folders and store like
  - `client1` is a folder with the path `/notes/client1`
  - `misc` is a folder with the path `/notes/client1/misc`
  - `test.md` is a file with the path `/notes/client1/misc/test.md`
  - `test123.md` is a file with the path `/notes/client1/misc/test123./md`

- After this, display on index.html all paths and see if it works
- After this, when pressing a path that has `.md` transform that to `.html`
- After this convert folders to buttons 

