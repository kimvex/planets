# Weather

En una galaxia lejana, existen tres civilizaciones. Vulcanos, Ferengis y Betasoides. Cada civilización vive en paz en su respectivo planeta.
Dominan la predicción del clima mediante un complejo sistema informático.

### Premisas
● El planeta Ferengi se desplaza con una velocidad angular de 1 grados/día en sentido horario. Su distancia con respecto al sol es de 500Km.

● El planeta Betasoide se desplaza con una velocidad angular de 3 grados/día en sentido horario. Su distancia con respecto al sol es de 2000Km.

● El planeta Vulcano se desplaza con una velocidad angular de 5 grados/día en sentido anti­horario, su distancia con respecto al sol es de 1000Km.

● Todas las órbitas son circulares.

Cuando los tres planetas están alineados entre sí y a su vez alineados con respecto al sol, el sistema solar experimenta un período de sequía.

Cuando los tres planetas no están alineados, forman entre sí un triángulo. Es sabido que en el momento en el que el sol se encuentra dentro del triángulo, el sistema solar experimenta un período de lluvia, teniendo éste, un pico de intensidad cuando el perímetro del triángulo está en su máximo.

Las condiciones óptimas de presión y temperatura se dan cuando los tres planetas están alineados entre sí pero no están alineados con el sol.

Para poder utilizar el sistema como un servicio a las otras civilizaciones, los Vulcanos requieren tener una base de datos con las condiciones meteorológicas de todos los días y brindar una API REST de consulta sobre las condiciones de un día en particular.
1) Generar un modelo de datos con las condiciones de todos los días hasta 10 años en adelante utilizando un job para calcularlas.
2) Generar una API REST la cual devuelve en formato JSON la condición climática del día consultado.
3) Hostear el modelo de datos y la API REST en un cloud computing libre (como APP Engine o Cloudfoudry) y enviar la URL para consulta:

## Prueba de solución
Para poder comprobar la solución se levanto un servidor en DigitalOcean.

##### Metodo GET
La url de la API es https://planets.kimvex.com en ella se exponen dos endpoints
1) `/api/start_simulation/:years` Este nos permite generar una predicción de clima en los años que se pasan por parametro basado en la orbita de 360 grados de los planetas implicados.

1.1) `/api/start_simulation/:years?realLife=true` Este nos permite generar una predicción de clima en los años que se pasan por parametro basado en los años terrestres de 365 días.

Respuesta del endpoint:

    {
        "Days":3600,
        "DroughtDays":20,
        "OptimumDays":40,
        "RainDays":1194,
        "RainPeaksDay":[2628,2808,2952,3492],
        "RegularDays":2346,
        "token":"44e6c40d7b31"
    }

##### Metodo GET
2) `/api/weather_day/:day?token=44e6c40d7b31` Este endpoint nos permite obtener el clima de un dia específico de una silumlación realizada identificandola por el token generado al momento de la simulación. 

Respuesta del endpoint:

    {
        "day":200,
        "weather":"Regular"
    }
    

## Infraestructura
La infraestructura utilizada para este proyecto es a siguiente

● Lenguaje de Programación Go(Golang)

● Base de datos MySQL

● Cloud DigitalOcean

● Servidor NGINX

● Manejador DNS Cloudflare


#### Estructura de la base de datos
Nombre de la base de datos: `weather`

Se crearon dos tablas `solarsystem` y `weather_day` las cuales guardan las simulaciones y el clima por día.

Estructura de la tabla `solarsystem`
![solarsystem](https://res.cloudinary.com/h27hacklab/image/upload/v1613356271/shop_images/Captura_de_Pantalla_2021-02-14_a_la_s_20.27.24.png)

Estructura de la tabla `weather_day`
![weather_day](https://res.cloudinary.com/h27hacklab/image/upload/v1613356270/shop_images/Captura_de_Pantalla_2021-02-14_a_la_s_20.27.44.png)

