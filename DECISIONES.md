# DECISIONES (CP1) — máximo 10 líneas

Responda: ¿qué entidades identificó en las pantallas y por qué se relacionan
así? Sea concreto: qué vio en cada pantalla que lo llevó a cada campo y a
cada clave foránea.

1. **Arreglo**: La pantalla 01 muestra el catálogo de arreglos florales con su nombre, precio y stock disponible.
2. **Cliente**: La pantalla 02 exige los datos de quien compra (Nombre, Cédula, Teléfono) separándolo de la orden.
3. **Encargo**: Representa la orden de compra. La pantalla 02 muestra que un Cliente elige un Arreglo (generando `cliente_id` y `arreglo_id` como claves foráneas), indicando una `Cantidad`, para calcular un `Total` y registrarlo en un `Estado`.
