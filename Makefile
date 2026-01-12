
up:
	@echo "Iniciando imagen docker.."
	docker compose up -d
	@echo "Imagenes docker iniciadas!"

down:
	@echo "Deteniendo imagenes docker"
	docker compose down 
	@echo "Hecho!"