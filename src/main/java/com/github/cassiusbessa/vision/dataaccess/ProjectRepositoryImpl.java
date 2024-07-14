package com.github.cassiusbessa.vision.dataaccess;

import com.github.cassiusbessa.vision.dataaccess.entities.ProjectDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.mappers.ProjectDataBaseMapper;
import com.github.cassiusbessa.vision.dataaccess.repositories.ProjectJpaRepository;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.service.ports.output.ProjectRepository;
import org.springframework.stereotype.Repository;

import java.util.Optional;
import java.util.UUID;

@Repository
public class ProjectRepositoryImpl implements ProjectRepository {

    private final ProjectJpaRepository projectRepository;
    private final ProjectDataBaseMapper projectDataBaseMapper;

    public ProjectRepositoryImpl(ProjectJpaRepository projectRepository, ProjectDataBaseMapper projectDataBaseMapper) {
        this.projectRepository = projectRepository;
        this.projectDataBaseMapper = projectDataBaseMapper;
    }


    @Override
    public Project findByProjectId(UUID projectId) {

        System.out.println(projectId);

        Optional<ProjectDataBaseEntity> projectDataBaseEntity = projectRepository.findById(projectId);

        return projectDataBaseEntity.map(projectDataBaseMapper::dbEntityToProject).orElse(null);


    }
}
