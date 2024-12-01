FROM node:alpine
WORKDIR /app
COPY . .
RUN npm i -g pnpm
RUN pnpm i --frozen-lockfile
CMD ["pnpm", "run", "dev"]
