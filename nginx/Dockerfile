# ===== Stage 1: Build frontend =====
FROM node:22-alpine AS builder

# 使用国内镜像加速
RUN npm config set registry https://registry.npmmirror.com

WORKDIR /app
COPY frontend/package.json ./
RUN npm install

COPY frontend/ .
RUN npm run build

# ===== Stage 2: Serve with Nginx =====
FROM nginx:alpine

# 复制 nginx 配置
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

# 复制前端构建产物
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
