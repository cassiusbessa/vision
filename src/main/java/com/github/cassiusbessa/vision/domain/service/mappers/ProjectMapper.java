package com.github.cassiusbessa.vision.domain.service.mappers;

import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.service.dtos.project.ProjectDTO;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;
import org.springframework.stereotype.Component;

import java.util.HashSet;

@Component
public class ProjectMapper {

    ProjectDTO projectToProjectDTO(Project project) {

        if (project == null) {
            return null;
        }

        var projectDTOTags = new HashSet<TagDTO>();

        for (var tag : project.getTechnologies()) {
            projectDTOTags.add(new TagDTO(tag.getId().getValue(), tag.getName()));
        }

        return new ProjectDTO(
            project.getId().getValue(),
            project.getAccount().getId().getValue(),
            project.getTitle(),
            project.getImage(),
            project.getDescription(),
            project.getLinks(),
            projectDTOTags,
            project.getCreatedAt()
        );
    }
}
