package com.github.cassiusbessa.vision.dataaccess;

import com.github.cassiusbessa.vision.dataaccess.mappers.ProjectDataBaseMapper;
import com.github.cassiusbessa.vision.dataaccess.repositories.ProjectJpaRepository;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.service.ports.output.repositories.ProjectRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.UUID;


@Repository
public class ProjectRepositoryImpl implements ProjectRepository {

    private final ProjectJpaRepository projectRepository;
    private final ProjectDataBaseMapper projectDataBaseMapper;

    @Autowired
    public ProjectRepositoryImpl(ProjectJpaRepository projectRepository, ProjectDataBaseMapper projectDataBaseMapper) {
        this.projectRepository = projectRepository;
        this.projectDataBaseMapper = projectDataBaseMapper;
    }


    @Override
    public List<Project> findALlByAccountId(UUID accountId) {
        return List.of();
    }

    @Override
    public Project findByTitle(String title) {
        return null;
    }

    @Override
    public void save(Project project) { projectRepository.save(projectDataBaseMapper.projectToDbEntity(project));}

    @Override
    public void update(Project project) {

    }

    @Override
    public void delete(Project project) {

    }
}
