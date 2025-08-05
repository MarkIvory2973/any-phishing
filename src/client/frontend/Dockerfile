FROM nginx:alpine

WORKDIR /usr/share/nginx/html/

COPY ./nginx.conf /etc/nginx/

COPY . .

CMD ["nginx", "-g", "daemon off;"]