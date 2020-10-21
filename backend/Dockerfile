FROM node:12.19.0
WORKDIR /app
RUN apt-get update && \
    npm install && \
    rm -rf /var/lib/apt/lists/*
COPY src/ src/
COPY package.json .
CMD ["node", "src/server.js"]