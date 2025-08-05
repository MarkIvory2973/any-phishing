FROM node:24-alpine AS build

WORKDIR /src/

COPY package.json yarn.lock ./
RUN yarn

COPY . .
RUN yarn format
RUN yarn build

FROM nginx:alpine

WORKDIR /usr/share/nginx/html/

COPY --from=build /src/dist/ .

CMD ["nginx", "-g", "daemon off;"]