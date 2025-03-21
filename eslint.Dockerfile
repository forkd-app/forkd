FROM node:22-alpine AS base
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
ENV COREPACK_INTEGRITY_KEYS=0

RUN corepack enable

WORKDIR /app

FROM base AS builder
ENV NODE_ENV development
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml .eslintrc.cjs .eslintignore .prettierignore .prettierrc tsconfig.json ./
COPY ./web/ ./web/
RUN corepack prepare
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --frozen-lockfile
CMD ["pnpm", "run", "lint"]
