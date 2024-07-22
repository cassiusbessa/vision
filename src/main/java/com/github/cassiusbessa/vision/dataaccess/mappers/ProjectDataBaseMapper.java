package com.github.cassiusbessa.vision.dataaccess.mappers;

import com.github.cassiusbessa.vision.dataaccess.entities.ProjectDataBaseEntity;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectId;
import com.github.cassiusbessa.vision.domain.core.valueobjects.ProjectLinks;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.stream.Collectors;

@Component
public class ProjectDataBaseMapper {

    private final AccountDataBaseMapper accountMapper;
    private final TagDataBaseMapper tagMapper;

    @Autowired
    public ProjectDataBaseMapper(AccountDataBaseMapper accountMapper, TagDataBaseMapper tagMapper) {
        this.accountMapper = accountMapper;
        this.tagMapper = tagMapper;
    }

    public ProjectDataBaseEntity projectToDbEntity(Project project) {
        if (project == null) {
            return null;
        }
        return new ProjectDataBaseEntity(
                project.getId().getValue(),
                project.getTitle(),
                project.getImage(),
                project.getDescription(),
                project.getLinks().getRepository(),
                accountMapper.accountToDbEntity(project.getAccount()),
                project.getCreatedAt(),
                project.getTechnologies().stream().map(tagMapper::tagToDbEntity).collect(Collectors.toSet())
        );
    }

    public Project dbEntityToProject(ProjectDataBaseEntity dbEntity) {
        if (dbEntity == null) {
            return null;
        }
        return Project.builder()
                .withId(new ProjectId(dbEntity.getId()))
                .withTitle(dbEntity.getTitle())
                .withImage(dbEntity.getImage())
                .withDescription(dbEntity.getDescription())
                .withLinks(new ProjectLinks(dbEntity.getRepositoryLink(), dbEntity.getDemoLink()))
                .withAccount(accountMapper.dbEntityToAccount(dbEntity.getAccount()))
                .withCreatedAt(dbEntity.getCreatedAt())
                .withTechnologies(dbEntity.getTechnologies().stream().map(tagMapper::dbEntityToTag).collect(Collectors.toSet()))
                .build();
    }

    public void updateProject(Project project, ProjectDataBaseEntity projectDataBaseEntity) {
        projectDataBaseEntity.setTitle(project.getTitle());
        projectDataBaseEntity.setImage(project.getImage());
        projectDataBaseEntity.setDescription(project.getDescription());
        projectDataBaseEntity.setRepositoryLink(project.getLinks().getRepository());
        projectDataBaseEntity.setDemoLink(project.getLinks().getDemo());
        projectDataBaseEntity.setCreatedAt(project.getCreatedAt());
        projectDataBaseEntity.setTechnologies(project.getTechnologies().stream().map(tagMapper::tagToDbEntity).collect(Collectors.toSet()));
    }
}
