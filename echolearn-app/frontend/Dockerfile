# Stage 1: Build the frontend using Node.js
FROM node:20-alpine AS builder

WORKDIR /app

# Copy dependencies and install
COPY package*.json tsconfig.json vite.config.ts ./
RUN npm install --frozen-lockfile

# Copy source files and build the static files
COPY src ./src
COPY index.html ./
RUN npm run build

# Stage 2: Run the static server
FROM node:20-alpine AS runner

WORKDIR /app

# Install serve to serve static files
RUN npm install -g serve

# Copy built static files
COPY --from=builder /app/dist /app

# Expose port and run the server
EXPOSE 3000
CMD ["serve", "-s", "-l", "3000", "/app"]
