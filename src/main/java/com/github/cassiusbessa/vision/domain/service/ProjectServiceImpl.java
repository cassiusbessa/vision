package com.github.cassiusbessa.vision.domain.service;

import com.github.cassiusbessa.vision.domain.core.entities.Account;
import com.github.cassiusbessa.vision.domain.core.entities.Project;
import com.github.cassiusbessa.vision.domain.core.entities.Tag;
import com.github.cassiusbessa.vision.domain.core.events.ProjectCreatedEvent;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreateCommand;
import com.github.cassiusbessa.vision.domain.service.dtos.ProjectCreatedResponse;
import com.github.cassiusbessa.vision.domain.service.exceptions.ResourceNotFoundException;
import com.github.cassiusbessa.vision.domain.service.exceptions.ValidationException;
import com.github.cassiusbessa.vision.domain.service.mappers.ProjectMapper;
import com.github.cassiusbessa.vision.domain.service.ports.input.ProjectService;
import com.github.cassiusbessa.vision.domain.service.ports.output.messages.ProjectCreatedMessagePublisher;
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
    private final ProjectCreatedMessagePublisher projectCreatedMessagePublisher;
    private final ProjectRepository projectRepository;
    private final TagRepository tagRepository;
    private final AccountRepository accountRepository;

    @Autowired
    public ProjectServiceImpl(ProjectMapper projectMapper, ProjectCreatedMessagePublisher projectCreatedMessagePublisher, ProjectRepository projectRepository, TagRepository tagRepository, AccountRepository accountRepository) {
        this.projectMapper = projectMapper;
        this.projectCreatedMessagePublisher = projectCreatedMessagePublisher;
        this.projectRepository = projectRepository;
        this.tagRepository = tagRepository;
        this.accountRepository = accountRepository;
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
            new ProjectCreatedEvent(project, new Date(), projectCreatedMessagePublisher).fire();
            log.info("Project created event fired: {}", project.getId().getValue());
        });
    }

}
