FROM ghcr.io/astral-sh/uv:0.8.8-python3.9-alpine
RUN apk update && apk add git
RUN uv tool install commitizen
WORKDIR /app
CMD [ "cz" ]
