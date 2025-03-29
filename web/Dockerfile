FROM node:23.10.0 AS base
WORKDIR /app
COPY package*.json ./

FROM base AS development
RUN npm install
COPY . .
CMD ["npm", "run", "dev"]

FROM base AS build
ARG VITE_API_PORT
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:1.27.4-alpine-slim AS production
COPY --from=build /app/dist /usr/share/nginx/html
COPY ./nginx /etc/nginx
CMD ["nginx", "-g", "daemon off;"]