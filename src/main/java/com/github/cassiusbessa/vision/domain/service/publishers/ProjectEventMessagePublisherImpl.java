package com.github.cassiusbessa.vision.domain.service.publishers;

import com.github.cassiusbessa.vision.domain.core.events.ProjectCreatedEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectUpdatedEvent;
import com.github.cassiusbessa.vision.domain.service.ports.output.messages.ProjectEventMessagePublisher;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Component;

@Slf4j
@Component
public class ProjectEventMessagePublisherImpl implements ProjectEventMessagePublisher {


    @Override
    public void publish(ProjectCreatedEvent event) {
        log.info("Project created event fired: {}", event.getProject().getTitle());
    }

    @Override
    public void publish(ProjectUpdatedEvent event) {
        log.info("Project updated event fired: {}", event.getProject().getTitle());
    }

    @Override
    public void publish(ProjectEvent event) {

        System.out.println("Project event: " + event.getProject().getTitle());

    }
}
