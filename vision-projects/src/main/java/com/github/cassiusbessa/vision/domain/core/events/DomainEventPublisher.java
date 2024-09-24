package com.github.cassiusbessa.vision.domain.core.events;

public interface DomainEventPublisher <T extends DomainEvent>{
    void publish(T event);
}
