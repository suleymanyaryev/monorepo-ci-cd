FROM node:14-alpine as base
WORKDIR /home/app
COPY package.json yarn.lock ./
RUN yarn install
COPY . .

FROM base as test

FROM base as build
ENV MODE=production
RUN yarn build

FROM nginx:1.20-alpine as prod
WORKDIR /home/app
RUN rm /etc/nginx/conf.d/default.conf && mkdir -p /etc/nginx/templates/
COPY ./nginx/app.conf.template /etc/nginx/templates/
COPY --from=build /home/app/dist/ ./
