FROM node:alpine AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

FROM base AS builder
ENV NODE_ENV development
WORKDIR /app
COPY ./package.json .
COPY ./pnpm-lock.yaml .
COPY ./web ./web
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile

FROM base as app
ENV NODE_ENV development
COPY --from=builder /app /app
WORKDIR /app/web
CMD ["pnpm", "run", "dev"]
