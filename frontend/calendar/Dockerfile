# ベースイメージを指定します
FROM node:latest

# 作業ディレクトリを指定します
WORKDIR /app

# 依存関係をインストールします
COPY package*.json ./
RUN npm install

# アプリケーションのソースコードをコピーします
COPY . .

# Astroをビルドします
RUN npx astro build

# ビルドされたファイルを公開するためのWebサーバーを指定します（例：nginx）
FROM nginx:alpine

# Nginxのデフォルト設定を上書きします
RUN rm -rf /usr/share/nginx/html/*
COPY --from=0 /app/dist /usr/share/nginx/html

# ポート80を公開します
EXPOSE 80

# コンテナが起動した際に実行するコマンドを指定します
CMD ["nginx", "-g", "daemon off;"]
