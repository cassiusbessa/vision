# Usar uma imagem base oficial do Node.js
FROM node:22-alpine3.19 AS builder

# Definir diretório de trabalho no contêiner
WORKDIR /app

# Copiar arquivos de configuração do projeto para o contêiner
COPY package.json package-lock.json ./

# Instalar as dependências do projeto (incluindo as de desenvolvimento)
RUN npm install --legacy-peer-deps

# Copiar o restante do código-fonte para o contêiner
COPY . .

# Expor a porta do Next.js
EXPOSE 3000

# Iniciar a aplicação
CMD ["npm", "run", "dev"]

# Etapa de produção
# FROM node:22-alpine3.19 AS production

# # Definir diretório de trabalho no contêiner
# WORKDIR /app

# # Copiar apenas os arquivos essenciais para produção
# COPY --from=builder /app/.next ./.next
# COPY --from=builder /app/public ./public
# COPY --from=builder /app/package.json ./
# COPY --from=builder /app/node_modules ./node_modules

# # Expor a porta que o Next.js irá rodar
# EXPOSE 3000

# # Comando para iniciar a aplicação
# CMD ["npm", "start"]
