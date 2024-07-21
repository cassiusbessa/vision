package com.github.cassiusbessa.vision.domain.service.ports.output.messages;

import com.github.cassiusbessa.vision.domain.core.events.DomainEventPublisher;
import com.github.cassiusbessa.vision.domain.core.events.ProjectCreatedEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectUpdatedEvent;

public interface ProjectEventMessagePublisher extends DomainEventPublisher<ProjectEvent> {

    void publish(ProjectCreatedEvent event);

    void publish(ProjectUpdatedEvent event);
}
