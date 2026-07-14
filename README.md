# 1. Construir la imagen
docker build -t go-demo-app:1.0.0 .

# 2. Correr el contenedor
docker run -d -p 8080:8080 --name go-demo-container go-demo-app:1.0.0

# 3. Verificar respuesta
curl http://localhost:8080