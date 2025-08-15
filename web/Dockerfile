FROM node:20-alpine AS build
WORKDIR /web
COPY . .
RUN npm ci || npm i && npm run build

FROM nginx:alpine
COPY --from=build /web/dist /usr/share/nginx/html
EXPOSE 80