FROM node:16.14 as node
WORKDIR /app
COPY . .

RUN npm i -g npm@8.5.1
RUN npm install
RUN npm run build --prod

FROM nginx:alpine
COPY --from=node /app/dist/lab4 /usr/share/nginx/html
