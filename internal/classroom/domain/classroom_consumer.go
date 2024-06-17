package domain

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/sumelms/microservice-classroom/internal/shared"
)

type ClassroomConsumer struct {
	service *Service
}

type CoursesEvent struct {
	Event       string
	CourseUUID  *uuid.UUID
	MatrixUUID  *uuid.UUID
	SubjectUUID *uuid.UUID
}

func NewClassroomConsumer(coursesAMQPConnection *shared.AMQPConnection, service *Service) error {
	matricesQueue, err := coursesAMQPConnection.NewAMQPQueue("MatricesQueue")
	if err != nil {
		return err
	}

	amqpDelivery, err := matricesQueue.Consume()
	if err != nil {
		return fmt.Errorf("AMQP Consume error: %w", err)
	}

	consumer := &ClassroomConsumer{
		service: service,
	}

	forever := make(chan bool)
	go func() {
		for delivery := range amqpDelivery.Delivery {
			consumer.ConsumeEvent(delivery.Body)
		}
	}()
	<-forever

	return nil
}

var classroomConsumers = map[string]func(consumer *ClassroomConsumer, event CoursesEvent) error{
	"SubjectDeleted": ConsumeSubjectDeletedEvent,
}

func (consumer *ClassroomConsumer) ConsumeEvent(data []byte) error {
	var event CoursesEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		return fmt.Errorf("JSON Unmarshal error: %w", err)
	}

	if eventConsumerFunction, exists := classroomConsumers[event.Event]; exists {
		eventConsumerFunction(consumer, event)
	} else {
		fmt.Println("Event not found")
	}

	return nil
}

func ConsumeSubjectDeletedEvent(consumer *ClassroomConsumer, event CoursesEvent) error {
	fmt.Println(event)
	classroom := &DeletedClassroom{
		SubjectUUID: event.SubjectUUID,
	}
	deletedClassrooms := consumer.service.classrooms.DeleteClassroom(classroom)
	fmt.Println("Deleted classrooms: ", deletedClassrooms)
	return nil
}
