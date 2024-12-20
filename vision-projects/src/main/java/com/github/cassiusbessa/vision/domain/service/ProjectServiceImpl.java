package com.github.cassiusbessa.vision.domain.service;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.core.events.ProjectCreatedEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectDeletedEvent;
import com.github.cassiusbessa.vision.domain.core.events.ProjectUpdatedEvent;
import com.github.cassiusbessa.vision.domain.service.dtos.project.ProjectDTO;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectDeleteCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.project.commands.ProjectUpdateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.project.queries.LoadProjectsByAccountIdQuery;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.LoadedProjectsResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.ProjectCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.ProjectDeletedResponse;
import com.github.cassiusbessa.vision.domain.service.dtos.project.responses.ProjectUpdatedResponse;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.mappers.ProjectMapper;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProjectService;
import com.github.cassiusbessa.vision.domain.service.ports.output.messages.ProjectEventMessagePublisher;
import com.github.cassiusbessa.vision.domain.service.ports.output.repositories.AccountRepository;
import com.github.cassiusbessa.vision.domain.service.ports.output.repositories.ProjectRepository;
import com.github.cassiusbessa.vision.domain.service.ports.output.repositories.TagRepository;
import jakarta.transaction.Transactional;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Date;
import java.util.List;
import java.util.UUID;
import java.util.concurrent.CompletableFuture;

@Service
@Slf4j
public class ProjectServiceImpl implements ProjectService {

    private final ProjectMapper projectMapper;
    private final ProjectRepository projectRepository;
    private final TagRepository tagRepository;
    private final AccountRepository accountRepository;
    private final ProjectEventMessagePublisher projectEventPublisher;

    @Autowired
    public ProjectServiceImpl(ProjectMapper projectMapper, ProjectRepository projectRepository, TagRepository tagRepository, AccountRepository accountRepository, ProjectEventMessagePublisher projectEventPublisher) {
        this.projectMapper = projectMapper;
        this.projectRepository = projectRepository;
        this.tagRepository = tagRepository;
        this.accountRepository = accountRepository;
        this.projectEventPublisher = projectEventPublisher;
    }

    @Override
    @Transactional
    public ProjectCreatedResponse createProject(ProjectCreateCommand command) {
        log.info("Creating project {}, for account: {}", command.title(), command.accountId());

        Account account = getAccount(command.accountId());

        List<Tag> tags = getTags(command.technologies());

        Project project = projectMapper.projectCreateCommandToProject(command, account, tags);

        validateProject(project);

        projectRepository.save(project);
        log.info("Project created successfully: {}", project.getId().getValue());

        fireProjectCreatedEvent(project);

        return new ProjectCreatedResponse("Project created successfully");
    }

    @Override
    @Transactional
    public ProjectUpdatedResponse updateProject(ProjectUpdateCommand command) {
        log.info("Updating project {}, for account: {}", command.title(), command.accountId());

        Account account = getAccount(command.accountId());

        List<Tag> tags = getTags(command.technologies());

        Project project = projectMapper.projectUpdateCommandToProject(command, account, tags);

        validateProject(project);

        projectRepository.update(project);
        log.info("Project updated successfully: {}", project.getId().getValue());

        fireProjectUpdatedEvent(project);

        return new ProjectUpdatedResponse("Project updated successfully");
    }

    @Override
    public ProjectDeletedResponse deleteProject(ProjectDeleteCommand command) {
        log.info("Deleting project: {}", command.projectId());

        Project isDeleted = projectRepository.delete(command.projectId());
        if (isDeleted == null) {
            log.error("Project does not exist: {}", command.projectId());
            throw new ResourceNotFoundException("Project does not exist: " + command.projectId());
        }

        log.info("Project deleted successfully: {}", command.projectId());

//        fireProjectDeletedEvent(isDeleted);
        return new ProjectDeletedResponse("Project deleted successfully");
    }

    @Override
    public LoadedProjectsResponse loadProjectsByAccountId(LoadProjectsByAccountIdQuery query) {
        log.info("Getting projects for account: {}", query.accountId());

        List<ProjectDTO> projects = projectRepository.findAllByAccountId(query.accountId()).stream().map(projectMapper::projectToProjectDTO).toList();

        log.info("Projects found: {}", projects.size());

        if (projects.isEmpty()) {
            log.error("No projects found for account: {}", query.accountId());
            throw new ResourceNotFoundException("No projects found for account: " + query.accountId());
        }

        return new LoadedProjectsResponse(projects, "Projects loaded successfully");
    }

		@Override
		public LoadedProjectsResponse loadProjectsByProfileId(LoadProjectsByAccountIdQuery query) {
				log.info("Getting projects for profile: {}", query.accountId());

				List<ProjectDTO> projects = projectRepository.findAllByProfileId(query.accountId()).stream().map(projectMapper::projectToProjectDTO).toList();

				log.info("Projects found: {}", projects.size());

				if (projects.isEmpty()) {
						log.error("No projects found for profile: {}", query.accountId());
						throw new ResourceNotFoundException("No projects found for profile: " + query.accountId());
				}

				return new LoadedProjectsResponse(projects, "Projects loaded successfully");
		}


    private void validateProject(Project project) {
        project.validate();
        if (!project.getFailureMessages().isEmpty()) {
            log.error("Project creation failed: {}", project.getFailureMessagesAsString());
            throw new ValidationException(project.getFailureMessagesAsString());
        }
    }

    private List<Tag> getTags(List<UUID> tagIds) {
        if (tagIds == null || tagIds.isEmpty()) {
            return List.of();
        }
        List<Tag> tags = tagRepository.findAllById(tagIds);
        validateTags(tags, tagIds);
        return tags;
    }

    private void validateTags(List<Tag> tags, List<UUID> tagIds) {
        if (tags.size() != tagIds.size()) {
            log.error("Some tags do not exist: {}, only found: {}", tagIds, tags);
            throw new ResourceNotFoundException("Some tags do not exist: " + tagIds);
        }
    }

    private Account getAccount(UUID accountId) {
        if (accountId == null) {
            log.error("Account ID is required");
            throw new ValidationException("Account ID is required");
        }
        Account account = accountRepository.findById(accountId);
        if (account == null) {
            log.error("Account does not exist: {}", accountId);
            throw new ResourceNotFoundException("Account does not exist: " + accountId);
        }
        return account;
    }

    private void fireProjectCreatedEvent(Project project) {
        CompletableFuture.runAsync(() -> {
            try {
                Thread.sleep(10000);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                log.error("Thread was interrupted while waiting fire create project event", e);
            }

            new ProjectCreatedEvent(project, new Date(), projectEventPublisher).fire();
            log.info("Project created event fired: {}", project.getId().getValue());
        });
    }

    private void fireProjectUpdatedEvent(Project project) {
        CompletableFuture.runAsync(() -> {
            try {
                Thread.sleep(10000);
            } catch (InterruptedException e) {
                Thread.currentThread().interrupt();
                log.error("Thread was interrupted while waiting fire updated project event", e);
            }
            new ProjectUpdatedEvent(project, new Date(), projectEventPublisher).fire();
            log.info("Project updated event fired: {}", project.getId().getValue());
        });
    }

    private void fireProjectDeletedEvent(Project project) {
        CompletableFuture.runAsync(() -> {
            new ProjectDeletedEvent(project, new Date(), projectEventPublisher).fire();
            log.info("Project deleted event fired: {}", project.getId().getValue());
        });
    }

}
