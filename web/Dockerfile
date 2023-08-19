FROM node:16

WORKDIR /todo_web/
COPY . .

RUN yarn && yarn build

FROM nginx:alpine

RUN rm -rf /etc/nginx/conf.d/default.conf # 使用80端口需要删除此默认配置文件
COPY .docker-compose/nginx/my.conf /etc/nginx/conf.d/my.conf
COPY --from=0 /todo_web/dist /usr/share/nginx/html
