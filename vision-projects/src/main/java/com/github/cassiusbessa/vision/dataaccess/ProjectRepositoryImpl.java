package com.github.cassiusbessa.vision.dataaccess;

import com.github.cassiusbessa.vision.dataaccess.entities.ProjectDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.entities.TagDataBaseEntity;
import com.github.cassiusbessa.vision.dataaccess.mappers.ProjectDataBaseMapper;
import com.github.cassiusbessa.vision.dataaccess.repositories.ProjectJpaRepository;
import com.github.cassiusbessa.vision.dataaccess.repositories.TagJpaRepository;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.ports.output.repositories.ProjectRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.*;


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
    public List<Project> findAllByAccountId(UUID accountId) {
        return projectRepository.findAllByAccountId(accountId).stream().map(projectDataBaseMapper::dbEntityToProject).toList();
    }

		@Override
		public List<Project> findAllByProfileId(UUID profileId) {
				return projectRepository.findAllByProfileId(profileId).stream().map(projectDataBaseMapper::dbEntityToProject).toList();
		}

    @Override
    public Project findByTitle(String title) {
        return null;
    }

    @Override
    public void save(Project project) { projectRepository.save(projectDataBaseMapper.projectToDbEntity(project));}

    @Override
    public void update(Project project) {

        ProjectDataBaseEntity projectDataBaseEntity = projectRepository.findById(project.getId().getValue()).orElse(null);
        if (projectDataBaseEntity == null) {
            throw new ResourceNotFoundException("Project not found");
        }

        projectDataBaseMapper.updateProject(project, projectDataBaseEntity);
        projectRepository.save(projectDataBaseEntity);
    }

    @Override
    public Project delete(UUID projectId) {

        Optional<ProjectDataBaseEntity> projectDataBaseEntity = projectRepository.findById(projectId);
        if (projectDataBaseEntity.isEmpty()) {
            throw new ResourceNotFoundException("Project not found");
        }

        projectRepository.delete(projectDataBaseEntity.get());
        return projectDataBaseMapper.dbEntityToProject(projectDataBaseEntity.get());

    }
}
