package com.github.cassiusbessa.vision.domain.service.ports.output.messages;

import com.github.cassiusbessa.vision.domain.core.events.DomainEventPublisher;
import com.github.cassiusbessa.vision.domain.core.events.ProjectCreatedEvent;

public interface ProjectCreatedMessagePublisher extends DomainEventPublisher<ProjectCreatedEvent> {
}
