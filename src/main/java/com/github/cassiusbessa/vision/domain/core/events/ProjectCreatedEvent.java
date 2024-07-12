package com.github.cassiusbessa.vision.domain.core.events;

import com.github.cassiusbessa.vision.domain.core.entities.Project;

import java.util.Date;

public class ProjectCreatedEvent extends ProjectEvent{

    private final DomainEventPublisher<ProjectCreatedEvent> publisher;

    public ProjectCreatedEvent(Project project, Date occurredOn, DomainEventPublisher<ProjectCreatedEvent> publisher) {
        super(project, occurredOn);
        this.publisher = publisher;
    }

    @Override
    public void fire() {
        publisher.publish(this);
    }
}
