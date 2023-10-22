FROM oven/bun:latest

WORKDIR /app

ENV PATH="/app/node_modules/.bin:$PATH"

COPY ./frontend/ /app/

# Install dependencies
RUN bun install

ENV NODE_ENV production

# Build
RUN bun run build

# Expose port 3000
EXPOSE 3000

# Set env to serve at port 3000
ENV PORT 3000

# Start serving
ENTRYPOINT [ "bun", "start" ]
