package com.github.cassiusbessa.vision.domain.core.events;

public interface DomainEventPublisher <T extends DomainEvent>{
    public void publish(T event);
}
