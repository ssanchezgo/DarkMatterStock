# DarkMatterStock

Part 1: Connect to the API and store the data
✅ Obtener datos de la API. Conexion exitosa, paginacion activa, datos procesando en consola. (connect.go , api.go)
✅ Guardar datos en DB. utilizando CockroachDB Cloud , crear tabla stock e insertar datos de la API (load_data.go)

Part 2: Create a simple API & UI
⏳ Backend API en Go. Crear endpoints para listar stocks, buscar, ordenar, etc.
❌ Frontend con Vue 3. Mostrar los datos en una interfaz ordenada, amigable y filtrable

Part 3: Recommend the best stocks to invest today
❌ Algoritmo de recomendación. (Basado en campos target, rating, acción, u otras fuentes externas)

Part 4: Write unit tests for your code
❌ Pruebas unitarias. Agregar tests en Go para asegurar que el sistema es confiable

Part 5: Deploy your code using Terraform
❌ Deploy con Terraform. Infraestructura como código, para desplegar en AWS o LocalStack


DarkMatterStock/
├── cmd/
│   └── main.go         // Punto de entrada principal para el backend
├── internal/
│   ├── api/            // Lógica para interactuar con la API externa
│   │   └── api.go
│   ├── db/             // Lógica para la conexión y operaciones con la base de datos
│   │   ├── connect.go  // Conexión a CockroachDB
│   │   ├── handlers.go // Manejadores de la API (si aplica aquí, o en cmd/main.go)
│   │   └── load_data.go // Carga inicial de datos desde la API a la DB
│   ├── model/          // Definición de modelos de datos (estructuras Go para Stocks, etc.)
│   │   └── stock.go
│   └── ui/             // Carpeta para el frontend (Vue 3)
│       ├── main.ts    # Punto de entrada de Vue
│   │   ├── App.vue    # Componente raíz de Vue
│   │   ├── /components # Componentes de Vue reutilizables
│   │   ├── /views      # Vistas de la aplicación
│   │   ├── /store      # Estado management con Pinia
│   │   ├── /assets     # Recursos estáticos como imágenes y estilos
│   │   └── /styles      # Archivos CSS, incluyendo Tailwindmain.ts    # Punto de entrada de Vue
├── go.mod
├── go.sum
└── README.md
