# Stage 1: Build the frontend
FROM node:20-alpine AS builder

WORKDIR /app
COPY package.json package-lock.json tsconfig.json vite.config.ts ./
RUN npm install

COPY src ./src
COPY index.html ./
RUN npm run build

# Stage 2: Serve the built files
FROM node:20-alpine AS runner

WORKDIR /app
RUN npm install -g serve
COPY --from=builder /app/dist /app
EXPOSE 3000
CMD ["serve", "-s", "-l", "3000", "/app"]
