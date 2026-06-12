# ===== Stage 1: Build frontend =====
FROM node:20-alpine AS builder

# 使用 pnpm（比 npm 快 2-3 倍）+ 国内镜像
RUN corepack enable && corepack prepare pnpm@latest --activate
RUN pnpm config set registry https://registry.npmmirror.com

WORKDIR /app
COPY frontend/package.json ./
RUN pnpm install

COPY frontend/ .
RUN pnpm run build

# ===== Stage 2: Serve with Nginx =====
FROM nginx:alpine

# 复制 nginx 配置
COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

# 复制前端构建产物
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
