package com.github.cassiusbessa.vision.domain.service.mappers;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectId;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectLinks;
import com.github.cassiusbessa.vision.domain.service.dtos.project.ProjectDTO;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectUpdateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.tag.TagDTO;
import org.springframework.stereotype.Component;

import java.util.*;
import java.util.stream.Collectors;

@Component
public class ProjectMapper {

    public Project projectCreateCommandToProject(ProjectCreateCommand command, Account account, List<Tag> tags) {
            if (command == null) {
                return null;
            }

            return Project.builder()
                    .withId(new ProjectId(UUID.randomUUID()))
                    .withAccount(account)
                    .withTitle(command.title())
                    .withDescription(command.description())
                    .withImage(command.imageLink())
                    .withLinks(projectCreateCommandsToProjectLinks(command))
                    .withTechnologies(new HashSet<>(tags))
                    .withCreatedAt(new Date())
                    .build();
    }

    public Project projectUpdateCommandToProject(ProjectUpdateCommand command, Account account, List<Tag> tags) {
            if (command == null) {
                return null;
            }

            return Project.builder()
                    .withId(new ProjectId(command.projectId()))
                    .withAccount(account)
                    .withTitle(command.title())
                    .withDescription(command.description())
                    .withImage(command.imageLink())
                    .withLinks(projectUpdateCommandsToProjectLinks(command))
                    .withTechnologies(new HashSet<>(tags))
                    .withCreatedAt(new Date())
                    .build();
    }

    public ProjectDTO projectToProjectDTO(Project project) {
        return new ProjectDTO(
                project.getId().getValue(),
                project.getAccount().getId().getValue(),
                project.getTitle(),
                project.getImage(),
                project.getDescription(),
                project.getLinks(),
                project.getTechnologies().stream().map(tag -> new TagDTO(tag.getId().getValue(), tag.getName())).collect(Collectors.toSet()),
                project.getCreatedAt()
        );
    }

    private ProjectLinks projectCreateCommandsToProjectLinks(ProjectCreateCommand command) {
        return new ProjectLinks(command.githubLink(), command.demoLink());
    }

    private ProjectLinks projectUpdateCommandsToProjectLinks(ProjectUpdateCommand command) {
        return new ProjectLinks(command.githubLink(), command.demoLink());
    }
}
