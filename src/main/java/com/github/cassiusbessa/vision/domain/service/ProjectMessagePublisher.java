package com.github.cassiusbessa.vision.domain.service;

import com.github.cassiusbessa.vision.domain.core.events.ProjectCreatedEvent;
import com.github.cassiusbessa.vision.domain.service.ports.output.messages.ProjectCreatedMessagePublisher;
import org.springframework.stereotype.Component;

@Component
public class ProjectMessagePublisher implements ProjectCreatedMessagePublisher {

    @Override
    public void publish(ProjectCreatedEvent event) {
        System.out.println("Project created>>>>>: " + event.getProject().getTitle());
    }
}
