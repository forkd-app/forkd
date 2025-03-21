FROM node:23-alpine AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /app

FROM base AS builder
ENV NODE_ENV development
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./
COPY web ./web/
WORKDIR /app/web
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
RUN ls
WORKDIR /app

FROM base AS app
ENV NODE_ENV=development
WORKDIR /app/web
COPY --from=builder /app/web /app/web
COPY --from=builder /app/node_modules /app/node_modules
CMD ["pnpm", "run", "dev"]
