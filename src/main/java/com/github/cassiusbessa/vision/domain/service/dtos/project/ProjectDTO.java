package com.github.cassiusbessa.vision.domain.service.dtos.project;

import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectLinks;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;

import java.util.Date;
import java.util.Set;
import java.util.UUID;

public class ProjectDTO {

    private final UUID id;

    private final UUID accountId;

    private final String title;

    private final String image;

    private final String description;

    private final ProjectLinks links;

    private final Set<TagDTO> technologies;

    private final Date createdAt;

    public ProjectDTO(UUID id, UUID accountId, String title, String image, String description, ProjectLinks links, Set<TagDTO> technologies, Date createdAt) {
        this.id = id;
        this.accountId = accountId;
        this.title = title;
        this.image = image;
        this.description = description;
        this.links = links;
        this.technologies = technologies;
        this.createdAt = createdAt;
    }

    public UUID getId() {
        return id;
    }

    public UUID getAccountId() {
        return accountId;
    }

    public String getTitle() {
        return title;
    }

    public String getImage() {
        return image;
    }

    public String getDescription() {
        return description;
    }

    public ProjectLinks getLinks() {
        return links;
    }

    public Set<TagDTO> getTechnologies() {
        return technologies;
    }

    public Date getCreatedAt() {
        return createdAt;
    }
}
