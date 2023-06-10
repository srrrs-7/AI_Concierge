FROM node:20

WORKDIR /app

COPY . /app

RUN npm i esbuild -D
RUN npx astro add node
RUN npm install playwright
RUN npx playwright install
RUN npx playwright install-deps
RUN npm i @astrojs/node
RUN npm i @storybook/test-runner --save-dev

RUN npm i

CMD [ "npm", "run", "storybook" ]
