# need to run the build for the first time if you have a new project. necessary to copy the package files in to do a yarn install and get all the dependencies in.
#docker build -f .\Dockerfile . -t react_dev_envt

# run this whenever you want to start developing
#docker run -it --rm -p 3000:3000 -v [current_path]:/app react_dev_envt
#docker run -it --rm -p 3000:3000 -v D:\Sync\NUS\CS3219\Assignment" "and" "Project\cs3219_project\frontend\app:/app react_dev_envt
FROM node:current-alpine

ENV NPM_CONFIG_LOGLEVEL warn
ENV CHOKIDAR_USEPOLLING true

RUN mkdir /node_modules

WORKDIR /app
COPY ./app ./

# don't install on volume. this is to improve performance.
RUN yarn install --modules-folder /node_modules

# will start the development server on startup

# create symlink
# ensure packages up to date
# start server
CMD ln -sf /node_modules/; yarn install --modules-folder /node_modules; yarn start