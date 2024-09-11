FROM node:alpine
WORKDIR /code
COPY . .
RUN npm i -g pnpm
RUN pnpm i --frozen-lockfile
CMD ["pnpm", "run", "dev"]
