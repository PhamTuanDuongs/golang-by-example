#Use the official node:16.20.2 image as the base image
FROM  node:16.20.2

# set the working direction
WORKDIR /app

# add app
COPY . ./

# Expose port 3000 for react app
EXPOSE 3000
# start app
CMD ["yarn", "start"]