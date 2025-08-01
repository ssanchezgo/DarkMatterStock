# DarkMatterStock

Part 1: Connect to the API and store the data
✅ Obtener datos de la API. Conexion exitosa, paginacion activa, datos procesando en consola. (main.go)
⏳ Guardar datos en DB. utilizando CockroachDB Cloud , crear tabla stock e insertar datos de la API (connect.go)

Part 2: Create a simple API & UI
⏳ Backend API en Go. Crear endpoints para listar stocks, buscar, ordenar, etc.
❌ Frontend con Vue 3. Mostrar los datos en una interfaz ordenada, amigable y filtrable

Part 3: Recommend the best stocks to invest today
❌ Algoritmo de recomendación. (Basado en campos target, rating, acción, u otras fuentes externas)

Part 4: Write unit tests for your code
❌ Pruebas unitarias. Agregar tests en Go para asegurar que el sistema es confiable

Part 5: Deploy your code using Terraform
❌ Deploy con Terraform. Infraestructura como código, para desplegar en AWS o LocalStack

/Dark_Matter_Stock
├── /cmd               # Para los puntos de entrada de la aplicación
│   └── myapp         # Código del ejecutable principal
│       └── main.go
├── /pkg               # Paquetes que pueden ser importados en proyectos externos
│   └── mypackage
├── /internal          # Código que solo será utilizado dentro de este proyecto
│   └── myinternal
├── /api               # Rutas y controladores de la API REST
│   ├── router.go
│   └── handlers.go
├── /web               # Código del frontend
│   ├── /public        # Archivos estáticos públicos
│   ├── /src          # Fuente del frontend
│   │   ├── main.ts    # Punto de entrada de Vue
│   │   ├── App.vue    # Componente raíz de Vue
│   │   ├── /components # Componentes de Vue reutilizables
│   │   ├── /views      # Vistas de la aplicación
│   │   ├── /store      # Estado management con Pinia
│   │   ├── /assets     # Recursos estáticos como imágenes y estilos
│   │   └── /styles      # Archivos CSS, incluyendo Tailwind
├── /migrations         # Scripts para gestionar cambios en la base de datos
├── go.mod             # Manejo de dependencias de Go
└── package.json       # Manejo de dependencias de Node.js