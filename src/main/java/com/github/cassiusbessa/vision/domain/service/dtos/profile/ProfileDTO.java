package com.github.cassiusbessa.vision.domain.service.dtos.profile;

import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.service.dtos.project.ProjectDTO;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;

import java.util.Set;
import java.util.UUID;

public class ProfileDTO {

    private final UUID id;
    private final String name;
    private final String image;
    private final String title;
    private final String description;
    private final Set<TagDTO> technologies;
    private final ProjectDTO starProject;
    private final String link;

    public ProfileDTO(UUID id, String name, String image, String title, String description, Set<TagDTO> technologies, ProjectDTO starProject, String link) {
        this.id = id;
        this.name = name;
        this.image = image;
        this.title = title;
        this.description = description;
        this.technologies = technologies;
        this.starProject = starProject;
        this.link = link;
    }

    public UUID getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public String getImage() {
        return image;
    }

    public String getTitle() {
        return title;
    }

    public String getDescription() {
        return description;
    }

    public Set<TagDTO> getTechnologies() {
        return technologies;
    }

    public ProjectDTO getStarProject() {
        return starProject;
    }

    public String getLink() {
        return link;
    }
}
