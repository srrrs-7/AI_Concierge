FROM node:20-alpine

WORKDIR /app

COPY . /app

CMD [ "npm", "run", "storybook" ]
