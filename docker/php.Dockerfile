FROM node:12-buster-slim

WORKDIR /usr/src/app
COPY package.json ./
RUN npm install
COPY . .
RUN npm install vuex@next --save && npm install --save-dev vue-loader@next
COPY . .

EXPOSE 7777
CMD ["npm", "start"]