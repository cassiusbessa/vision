package com.github.cassiusbessa.vision.domain.service.ports.output.messages;

import com.github.cassiusbessa.vision.domain.core.events.*;

public interface ProjectEventMessagePublisher extends DomainEventPublisher<ProjectEvent> {

    void publish(ProjectCreatedEvent event);

    void publish(ProjectUpdatedEvent event);

    void publish(ProjectDeletedEvent event);
}
