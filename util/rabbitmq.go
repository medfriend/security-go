package util

import (
	"github.com/streadway/amqp"
	"log"
	"sync"
)

// RabbitMQ es la estructura que representa nuestra conexión y canal
type RabbitMQ struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

// Instancia es la única instancia de RabbitMQ
var instancia *RabbitMQ
var once sync.Once

// GetInstance devuelve la instancia única de RabbitMQ
func GetInstance() *RabbitMQ {
	once.Do(func() {
		var err error
		instancia = &RabbitMQ{}

		instancia.conn, err = amqp.Dial(GetRabbitConn())
		if err != nil {
			log.Fatalf("Error al conectar a RabbitMQ: %s", err)
		}

		instancia.ch, err = instancia.conn.Channel()
		if err != nil {
			log.Fatalf("Error al abrir un canal: %s", err)
		}
	})

	return instancia
}

// SendMessage envía un mensaje a la cola especificada
func (r *RabbitMQ) SendMessage(queueName string, message string) {
	// Declarar la cola
	_, err := r.ch.QueueDeclare(
		queueName, // Nombre de la cola
		false,     // Duradera
		false,     // Autodelete
		false,     // Exclusiva
		false,     // Sin esperar
		nil,       // Argumentos adicionales
	)
	if err != nil {
		log.Fatalf("Error al declarar la cola: %s", err)
	}

	// Enviar un mensaje
	err = r.ch.Publish(
		"",        // Intercambio
		queueName, // Clave de enrutamiento
		false,     // Requiere confirmación de entrega
		false,     // Exclusivo
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		log.Fatalf("Error al enviar el mensaje: %s", err)
	}

	log.Println("Mensaje enviado:", message)
}

// Close cierra la conexión y el canal
func (r *RabbitMQ) Close() {
	if r.ch != nil {
		r.ch.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}
