FROM node:20

WORKDIR /app

COPY . /app

RUN npm i esbuild -D
RUN npx astro add node

RUN npm i

CMD [ "npm", "run", "start" ]
