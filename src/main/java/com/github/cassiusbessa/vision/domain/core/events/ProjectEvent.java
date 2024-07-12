package com.github.cassiusbessa.vision.domain.core.events;

import com.github.cassiusbessa.vision.domain.core.entities.Project;

import java.util.Date;

public abstract class ProjectEvent implements DomainEvent<Project>{

    private final Project project;
    private final Date occurredOn;

    public ProjectEvent(Project project, Date occurredOn) {
        this.project = project;
        this.occurredOn = occurredOn;
    }

    public Project getProject() {
        return project;
    }

    public Date getOccurredOn() {
        return occurredOn;
    }
}
