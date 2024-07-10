package com.github.cassiusbessa.vision.domain.service.dtos.profile;

import com.github.cassiusbessa.vision.domain.core.entities.Project;

import java.util.List;
import java.util.UUID;

public class ProfileUpdateCommand {

    private final UUID profileId;
    private final String name;
    private final String title;
    private final String description;
    private final List<UUID> technologies;
    private final Project starProject;
    private final UUID accountId;
    private final String link;

    public ProfileUpdateCommand() {
        this.profileId = null;
        this.name = null;
        this.title = null;
        this.description = null;
        this.technologies = null;
        this.starProject = null;
        this.accountId = null;
        this.link = null;
    }

    public ProfileUpdateCommand(UUID profileId, String name, String title, String description, List<UUID> technologies, Project starProject, UUID accountId, String link) {
        this.profileId = profileId;
        this.name = name;
        this.title = title;
        this.description = description;
        this.technologies = technologies;
        this.starProject = starProject;
        this.accountId = accountId;
        this.link = link;
    }

    public UUID getProfileId() { return profileId; }

    public String getName() {
        return name;
    }

    public String getTitle() {
        return title;
    }

    public String getDescription() {
        return description;
    }

    public List<UUID> getTechnologies() {
        return technologies;
    }

    public Project getStarProject() {
        return starProject;
    }

    public UUID getAccountId() {
        return accountId;
    }

    public String getLink() {
        return link;
    }

}
