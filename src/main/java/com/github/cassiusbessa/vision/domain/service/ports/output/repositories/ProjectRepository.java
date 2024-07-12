package com.github.cassiusbessa.vision.domain.service.ports.output.repositories;

import com.github.cassiusbessa.vision.domain.core.entities.Profile;
import com.github.cassiusbessa.vision.domain.core.entities.Project;

import java.util.List;
import java.util.UUID;

public interface ProjectRepository {

    List<Project> findALlByAccountId(UUID accountId);

    Project findByTitle(String title);

    void save(Project project);

    void update(Project project);

    void delete(Project project);
}
