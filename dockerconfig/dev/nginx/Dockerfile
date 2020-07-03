FROM nginx:1.17-alpine


RUN rm /etc/nginx/conf.d/default.conf
RUN rm /etc/nginx/nginx.conf

COPY ./dockerconfig/dev/nginx/post.conf /etc/nginx/conf.d/post.conf
COPY ./dockerconfig/dev/nginx/nginx.conf /etc/nginx/nginx.conf





