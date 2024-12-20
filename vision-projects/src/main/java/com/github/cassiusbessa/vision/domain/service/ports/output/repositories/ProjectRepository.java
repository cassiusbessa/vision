package com.github.cassiusbessa.vision.domain.service.ports.output.repositories;

import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.core.entities.Project;

import java.util.List;
import java.util.UUID;

public interface ProjectRepository {

    List<Project> findAllByAccountId(UUID accountId);

		List<Project> findAllByProfileId(UUID profileId);

    Project findByTitle(String title);

    void save(Project project);

    void update(Project project);

    Project delete(UUID projectId);
}
