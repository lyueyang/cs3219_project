# Developing for React on Docker

## Motivation
I develop across MacOS and Windows and the lack of a package manager on Windows frustrates me.

To make development more OS agnostic, I've decided to pursue using Docker as part of the development process.

## How to Use
`Dockerfile.init` contains the info and config to generate a react project using create-react-app using the container instead of locally on your machine. 

You may skip this if you already have your project.

**Note:** The container may prompt you to install create-react-app packages on start. 

`Dockerfile` contains the information and configuration required to run the yarn development server.

This copies a project in the /app directory into the container. You may rename the directory but you must change it accordingly in the dockerfile.
You are encouraged to build this container only after the project files are ready so as to allow node_modules to be built and cached by the docker's host.
